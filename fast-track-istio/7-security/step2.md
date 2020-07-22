In the world of microservices, security can quickly become unwieldy with poor implementation and management. Traditionally, security policies have revolved around the network identity of a service, its IP address. However, IP addresses of workloads are ephemeral in Kubernetes and any other container orchestrator, so Istio solves this problem by decoupling the identity of a workload from the host.

## Applying transport authentication

Let's begin with deploying the **Books** service to the service mesh with the following command.

`kubectl apply -f my-workshop/books-api.yaml && kubectl apply -f my-workshop/books-api-vs-gw.yaml`{{execute}}

Currently, this service uses unsecured HTTP transport. Let's add a busybox pod inside and outside the mesh by applying the following specification.

`kubectl apply -f my-workshop/busybox.yaml`{{execute}}

Let's verify whether our services can communicate with each other by sending two requests to the same endpoint from the busybox pods.

- From outside the mesh

`kubectl exec -ti busybox -- curl http://books-api-service.fast-track-istio.svc.cluster.local:4000/books/1 -v`{{execute}}

- From within the mesh

`kubectl exec -ti -n fast-track-istio busybox -- curl http://books-api-service.fast-track-istio.svc.cluster.local:4000/books/1 -v`{{execute}}

Let's alter the behavior and apply a blanket mTLS-only policy over the namespace using the peer authentication policy.

<pre>
apiVersion: security.istio.io/v1beta1
kind: PeerAuthentication
metadata:
  name: default
  namespace: fast-track-istio
spec:
  mtls:
    mode: STRICT
</pre>

To apply the policy, execute the following command.

`kubectl apply -f my-workshop/default-mtls-policy.yaml`{{execute}}

To configure mTLS on the client, we will create a destination rule to enforce TLS on the client to the service channel.

<pre>
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: default-destination-rule
  namespace: istio-system
spec:
  host: "*.local"
  trafficPolicy:
    tls:
      mode: ISTIO_MUTUAL
</pre>

The wildcard match _\*.local_ makes the policy in previous listing applicable to all services in the mesh. Let's apply this policy to the mesh now.

`kubectl apply -f my-workshop/default-dr.yaml`{{execute}}

After applying the policy, only the services within the mesh can communicate with each other over a secure mTLS channel.

Finally, we require the following policy to restrict access to the **Books** service from pods that are associated with a service account.

<pre>
spec:
  selector:
    matchLabels:
      app: books
  rules:
    - from:
        - source:
            principals: ["cluster.local/ns/fast-track-istio/sa/books-viewer-sa"]
      to:
        - operation:
            methods: ["GET"]
</pre>

Execute the following command to apply the policy to the cluster.

`kubectl apply -f my-cluster/books-api-authz-policy.yaml`{{execute}}

Let's create another busybox pod inside the cluster which has a service account associated with it. We will use this pod to execute one of the tests later. Let's take a brief look at the specification first.

<pre>
spec:
  serviceAccountName: books-viewer-sa
  containers:
    - name: busybox-sa
      image: radial/busyboxplus:curl
      command:
        - sleep
        - "3600"
      imagePullPolicy: IfNotPresent
      resources:
        limits:
          cpu: "200m"
          memory: 200Mi
        requests:
          cpu: "100m"
          memory: 100Mi
  restartPolicy: Always
</pre>

The following is the specification of the service account that we configured above.

<pre>
apiVersion: v1
kind: ServiceAccount
metadata:
  name: books-viewer-sa
  namespace: fast-track-istio
  labels:
    app: busybox-sa
</pre>

To apply this specification, execute the following command.

`kubectl apply -f my-workshop/busybox-sa.yaml`{{execute}}

We are now ready to execute the busybox instructions again to witness the policy in action.

- From outside the mesh

`kubectl exec -ti busybox -- curl http://books-api-service.fast-track-istio.svc.cluster.local:4000/books/1 -v`{{execute}}

- From within the mesh without service account

`kubectl exec -ti -n fast-track-istio busybox -- curl http://books-api-service.fast-track-istio.svc.cluster.local:4000/books/1 -v`{{execute}}

- From within the mesh with service account

`kubectl exec -ti -n fast-track-istio busybox-sa -- curl http://books-api-service.fast-track-istio.svc.cluster.local:4000/books/1 -v`{{execute}}

Let's configure authorization on ingress gateway next. Before we proceed, let's delete our namespace to start afresh.

`kubectl delete namespace fast-track-istio`{{execute}}
