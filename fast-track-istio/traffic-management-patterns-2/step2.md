A resilient system degrades gracefully when failures occur in downstream systems. To build resilient systems, Istio provides several turnkey features such as client-side load balancing, outlier detection, automatic retry, and request timeouts. We have already discussed the various client-side load balancing strategies; now we will see how you can combine outlier detection (also called circuit breaker), request timeouts, and retries to ensure reliable communication between services. We will configure the following policies together:

 - Circuit breaker/outlier detection to eject faulty instances from the load balancer pool.
 - Timeout to avoid waiting on a faulty service.
 - Retries to forward the request to another instance in the load balancer pool if the primary instance is not responding.

For this demo, we will use an endpoint in the juice-shop API service that has a probability of 0.8 to fail. In the real world, you rarely encounter such services that perform so poorly. Our goal is to implement multiple resiliency policies such that we receive much better quality of service (QoS) without making any changes to the API itself. To ensure that we don’t have any existing resources that can affect our demo, execute the following command to delete all resources within the namespace and the namespace itself.
