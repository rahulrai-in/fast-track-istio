Traffic mirroring or dark launch is a model of deployment by which a new version of a service is deployed to production without affecting the traffic to the existing service. To enforce dark launch, Istio duplicates the requests to the primary service and sends them asynchronously to the service under test without waiting for the secondary service to respond. The asynchronous mirroring of traffic ensures that the critical path of live traffic remains unaffected.

Let's begin our exploration with the specification in the `my-workshop\mirroring\movies-versioned-svc-deploy.yaml` file. If this specifiction looks familiar, it is because you deployed it to realize the versioning pattern in the previous exercise.

Let's explore the specifications in the `my-workshop\mirroring\movies-mirror-gw-dr-vs.yaml` now. The following specification instructs Istio to direct all the traffic to v1 of the **movies** service and mirror the traffic to v2 of the **movies** service.

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

The specification for the destination rule is straightforward, we just need to declare our subsets here.

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

It is time to apply the specifications now. Execute the followign command to apply both the specifications.

`kuebctl apply -f my-workshop\mirroring\movies-mirror-gw-dr-vs.yaml -f my-workshop\mirroring\movies-versioned-svc-deploy.yaml`{{execute}}

Let's launch two other terminal windows by clicking on the **+** icon and selecting **Open New Terminal** from the options. Execute the following command in the first of the new terminal windows.

``