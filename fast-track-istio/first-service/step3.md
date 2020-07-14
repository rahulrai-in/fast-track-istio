Let's deploy our service to the mesh again. In the project files, inspect the contents of the `my-workshop\book-club-istio.yaml` specification. Here you will notice that we have set the value of the label `istio-injection` to *true*. Istio will read this value and automatically inject sidecars to any pods that it creates in this namespace.


`
apiVersion: v1
kind: Namespace
metadata:
  name: fast-track-istio 
  labels:
    istio-injection: enabled
`


There are no changes to the deployment specification. Also, since we will use a gateway to access the service, the type of service is set to ClusterIP (default). A ClusterIP service receives an IP address from the Kubernetes DNS, but it is accessible only within the cluster.


`
apiVersion: v1
kind: Service
metadata:
  name: bookclub
  namespace: fast-track-istio
  labels:
    app: bookclub
spec:
  ports:
    port: 8080
    protocol: TCP
  selector:
    app: bookclub
`


Execute the following command to deploy the application now.

`kubectl apply -f my-workshop/book-club-istio.yaml`{{execute}}

After some time, let's execute the following command to verify whether Istio injected a sidecar to our pod.

`kubectl get pods --selector app=bookclub -n fast-track-istio -o jsonpath={.items[*].spec.containers[*].name}`{{execute}}

## Ingress gateway

Now that our first service is deployed on the mesh, we will make the service accessible outside the cluster. We will deploy an ingress gateway that will route traffic originating from outside our cluster to the book-club application via the Envoy proxy. The specification of the gateway is present in the `my-workshop/bookclub-gw-vs.yaml` file.

The specification that you studied will create a gateway object that enables the book-club application to accept traffic originating outside the cluster. A virtual service applies traffic routing rules such that only the traffic that matches the routing rules reach the destination.

Let's apply the specification to our cluster with the following command.

`kubectl apply -f my-workshop/bookclub-gw-vs.yaml`{{execute}}

After the deployment succeeds, you will be able to access the application at the following URL.

https://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/

Congratulations! You learned to migrate existing services to service mesh (manual sidecar injection) and deploy new services to mesh (automatic sidecar injection). We are now ready to learn and apply some key Traffic Management capabilities to our services.