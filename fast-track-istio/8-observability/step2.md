A request may travel through multiple microservices before returning a response to the user in a distributed system. Therefore, microservices must implement distributed tracing with a correlation identifier so that the request flow graph of any request can be traced. Istio implements the [OpenTracing](https://opentracing.io/) standard. Istio also supports integration with popular visualization tools such as [Prometheus](https://prometheus.io/), [Grafana](https://grafana.com/), and [Kiali](https://www.kiali.io/), which can help developers and operations visualize the state of the mesh.

First, let's bring up all the services our application on our cluster with the following command.

`kubectl kustomize my-workshop/app | kubectl apply -f -`{{execute}}

> The [**Independent** service](https://github.com/rahulrai-in/fast-track-istio/tree/master/demo-app/independent) must be hosted outside the cluster for this workshop. After deploying the service, update the environment variable `independentServiceUri` of the book-club application specification with the URL of the service before applying the specification.

You can explore the user interface of the application at the following URL.

https://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/

Several observability tools are bundled in the Istio release package, and they are installed on your cluster with other Istio components. Some of the tools packaged with Istio are the following.

- Grafana
- Jaeger
- Kiali
- Prometheus
- Zipkin

The command `istioctl dashboard <dashboard-name>` will launch a browser to bring up one of the available dashboards. Let's explore the dashboards one at a time.

## Tracing with Jaeger

Jaeger is the default tracing system installed with Istio. You can customize the installer to use Zipkin instead. Istio manages communication with the OpenTracing engine and generates request tracing headers if they don't exist. If it receives a request with the headers populated, it doesn't generate them again and treats the trace as an in-progress trace. Istio relies on the following headers for distributed tracing:

- x-request-id
- x-b3-traceid
- x-b3-spanid
- x-b3-parentspanid
- x-b3-sampled
- x-b3-flags
- x-ot-span-context

If the service invokes another service, it must propagate the trace headers with the request so that Istio can correlate the upstream request with the incoming request to the service. Our demo application is instrumented to propagate the headers across the services, and therefore, we will be able to trace a request end to end.

The `istioctl dashboard` command creates a proxy that port-forwards the dashboard services to localhost and so they don't work on mesh deployed to a remote cluster. Therefore, we will create a gateway, a virtual service, and destination rule objects to make the service accessible on internet. The specification for exposing the tracing service is available in the `my-workshop/tracing.yaml` file.

`kubectl apply -f my-workshop/tracing.yaml`{{execute}}

Let's now navigate to the following URL to explore the dashboard.

https://[[HOST_SUBDOMAIN]]-15443-[[KATACODA_HOST]].environments.katacoda.com/

To generate some telemetry data that will light up the Jaeger dashboard, navigate around the application a bit. Do not forget to check the dependency graphs. 😎

Let's delete the resources before we move on to explore another dashboard.

`kubectl delete -f my-workshop/tracing.yaml`{{execute}}
