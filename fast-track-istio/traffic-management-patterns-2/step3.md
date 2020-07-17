Traffic mirroring or dark launch is a model of deployment by which a new version of a service is deployed to production without affecting the traffic to the existing service. To enforce dark launch, Istio duplicates the requests to the primary service and sends them asynchronously to the service under test without waiting for the secondary service to respond. The asynchronous mirroring of traffic ensures that the critical path of live traffic remains unaffected.

Let's execute the traffic mirroring sample from the [istio.io documentation](https://istio.io/latest/docs/tasks/traffic-management/mirroring/). In this task, you will first force all traffic to v1 of a test service. Then, you will apply a rule to mirror a portion of traffic to v2.

Let's deploy two versions of the [httpbin](https://github.com/istio/istio/tree/release-1.6/samples/httpbin) service.

