Traffic mirroring or dark launch is a deployment model by which a new version of a service is deployed to production without affecting the traffic to the existing service. To enforce the dark launch, Istio duplicates the requests to the primary service and sends them asynchronously to the service under test without waiting for the secondary service to respond. The asynchronous mirroring of traffic ensures that the critical path of live traffic remains unaffected.

Let's begin our exploration with the specification in the `my-workshop/mirroring/movies-versioned-svc-deploy.yaml` file. If this specification looks familiar, it is because you deployed it to realize the versioning pattern in the previous exercise.

Let's explore the specifications in the `my-workshop/mirroring/movies-mirror-gw-dr-vs.yaml` now. The following specification instructs Istio to direct all the traffic to v1 of the **movies** service and mirror the traffic to v2 of the **movies** service.

<pre>
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: movies-api-vservice
  namespace: fast-track-istio
spec:
  hosts:
    - "*"
  gateways:
    - movies-api-gateway
  http:
     - route:
        - destination:
            host: movies-api-service
            port:
              number: 3000
            subset: v1
          weight: 100
       mirror:
            host: movies-api-service
            port:
              number: 3000
            subset: v2
       mirror_percent: 100
</pre>

The specification for the destination rule is straightforward; we need to declare our subsets here.

<pre>
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: movies-api-destination
  namespace: fast-track-istio
spec:
  host: movies-api-service
  subsets:
    - name: v1
      labels:
        version: "1"
    - name: v2
      labels:
        version: "2"
</pre>

It is time to apply the specifications now. Execute the following command to apply both the specifications.

`kubectl apply -f my-workshop/mirroring/movies-versioned-svc-deploy.yaml -f my-workshop/mirroring/movies-mirror-gw-dr-vs.yaml`{{execute}}

Let's launch another terminal window by clicking on the **+** icon and selecting **Open New Terminal** from the options. Execute the following command in the new terminal window to follow the **movies** service's log trail.

`kubectl logs -l app=movies-api -n fast-track-istio -c movies-api -f`{{execute}}

Execute the following script to send 50 requests to the service in the original terminal window.

`for ((i=1;i<=50;i++)); do echo -n "Request #{$i}: "; curl -sS "http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/"; echo; done`{{execute}}

Switch to the terminal window where logs are getting populated and witness mirroring in action.
