A resilient system degrades gracefully when failures occur in downstream systems. To build resilient systems, Istio provides several turnkey features such as client-side load balancing, outlier detection, automatic retry, and request timeouts. We have already discussed the various client-side load balancing strategies; now, we will see how you can combine outlier detection (also called circuit breaker), request timeouts, and retries to ensure reliable communication between services. We will configure the following policies together:

- Circuit breaker/outlier detection to eject faulty instances from the load balancer pool.
- Timeout to avoid waiting on a faulty service.
- Retries to forward the request to another instance in the load balancer pool if the primary instance is not responding.

For this demo, we will use an endpoint in the **books** service that has a probability of 0.8 to fail. In the real world, you rarely encounter such services that perform so poorly. Our goal is to implement multiple resiliency policies such that we receive a much better quality of service (QoS) without making any changes to the application code itself. Let's first deploy the **books** service and its associated virtual service and gateway resources with the following command.

`kubectl apply -f my-workshop/resilience/books.yaml -f my-workshop/resilience/books-gw-vs.yaml`{{execute}}

The specifications that we applied do not implement any resiliency strategies yet. Let's execute a test that makes 50 requests to the FeelingLucky (/books/feeling-lucky/) endpoint with a high rate of producing errors. Execute the following script to launch the test.

`for ((i=1;i<=50;i++)); do echo -n "Request #{$i}: "; curl -sS "http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/books/feeling-lucky"; echo; done`{{execute}}

Let's now implement the resiliency strategies that we discussed previously. The following specification of virtual service adds a timeout period of 30 seconds to each request received from the client. Within this period, Envoy makes ten attempts to fetch results from the **books** service with a waiting period of 3 seconds between each attempt. We have also specified the failure conditions as the value of the retryOn key on which the retry policy kicks in. See `my-workshop/resilience/books-gw-dr-vs-resilence.yaml`.

<pre>
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: books-api-vservice
  namespace: fast-track-istio
spec:
  hosts:
    - "*"
  gateways:
    - books-api-gateway
  http:
    - route:
        - destination:
            host: books-api-service
            port:
              number: 4000
      timeout: 30s
      retries:
        attempts: 10
        perTryTimeout: 3s
        retryOn: "5xx,connect-failure,refused-stream"
</pre>

Next, we add a destination rule that detects and ejects the outliers from the load balancer pool. The following specification limits the number of parallel requests to the **books** service, and on receiving a single error (should not be 1 in real-world scenarios), ejects the endpoint from the load balancer pool for a minimum of three minutes. Envoy decides on the endpoints to eject and the ones to bring back in the pool every second, and it can eject up to 100 percent of the endpoints from the pool.

<pre>
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: books-api-destination-rule
  namespace: fast-track-istio
spec:
  host: books-api-service
  trafficPolicy:
    connectionPool:
      http:
        http1MaxPendingRequests: 1
        maxRequestsPerConnection: 1
    outlierDetection:
      consecutiveErrors: 1
      interval: 1s
      baseEjectionTime: 3m
      maxEjectionPercent: 100
</pre>

Let's apply this configuration to our cluster with the following command.

`kubectl apply -f my-workshop/resilience/books-gw-dr-vs-resilence.yaml`{{execute}}

Let's execute the same test that we previously executed to verify the effectiveness of this policy.

`for ((i=1;i<=50;i++)); do echo -n "Request #{$i}: "; curl -sS "http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/books/feeling-lucky"; echo; done`{{execute}}

When you execute the test this time, you will notice that the tests execute much slower due to the request timeout policy and retries happening in the mesh. That is a significant improvement without any code alterations. Note that these policies are not universal and are scoped to each client of the **books** service, as the service may fault for a single client while still functioning for others.

Let's delete the namespace and its resource before proceeding to the next pattern by executing the following command.

`kubectl delete namespace fast-track-istio`{{execute}}

It will take some time to clean up the resources within the namespace; let's proceed to the next step in the meanwhile.
