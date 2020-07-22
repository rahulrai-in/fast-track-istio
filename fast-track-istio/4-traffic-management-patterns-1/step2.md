Versioning is a critical aspect of microservices to maintain compatibility with existing clients. A microservice will need to change over time, and when that happens, you will need to incrementally deploy the new version of your service to ensure that clients can gradually migrate to the updated service. We will use subsets to differentiate between the different versions of the same service.

In this exercise, we will deploy two versions of the **movies** service and route traffic to appropriate endpoints using gateway rules. In the editor window, explore the specification in the file `my-workshop/versioning/movies-versioned-svc-deploy.yaml` which defines two deployments, one for each version of the **movies** service, and a service object that will randomly route requests to the pods created for the service.

We also require a virtual service and gateway resource to route requests to our services. The specification for the objects is present in the file `my-workshop/versioning/movies-gw-dr-vs.yaml`. While exploring this file, pay special attention to the `subset` declaration in `DestinationRule` specification.

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

Also, note how the `virtualservice` object directs requests (virtually) to the subsets with the following specification.

<pre>
http:
  match:
    - headers:
        version:
            exact: "2"
    route:
    - destination:
        host: movies-api-service
        port:
            number: 3000
        subset: v2
  route:
    - destination:
        host: movies-api-service
        port:
            number: 3000
        subset: v1
</pre>

The previous policy routes the requests with a header key named `version` and the value set to 2 to version 2 of the **movies** API, and other requests to version 1 of the service. Let's apply the specifications to the cluster with the following command.

`kubectl kustomize my-workshop/versioning | kubectl apply -f -`{{execute}}

Ensure that the service is up and running, and send a request to version 1 of the service - v1.

`curl http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/`{{execute}}

The following command will send a request to version 2 of the service - v2.

`curl http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/ -H "version: 2"`{{execute}}

In this example, we used request header to filter and route requests to one of the two versions of the service. However, you can also use one or more filters based on the various Level 4–7 attributes such as URL, scheme, method, authority, and port to configure traffic routing. You can use this feature to deploy canary releases, perform A\B tests, etc.
