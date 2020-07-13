You can define custom labels for all Kubernetes objects, which are essentially key-value pairs. Istio relies on a specific label named `istio-injection` to decide the namespace on which the Envoy proxies are applied. This label supports two values. Setting the value of the label to _enabled_ will insturct Istio to automatically deploy sidecars for all the pods of your service. On the other hand, setting this value to _disabled_ means that Istio will not inject sidecars automatically; however, it won’t affect the services within the namespace that have sidecars attached to them.

The journey to virtualization in enterprises is an incremental process. This means that in most of the cases, organizations would want to gradually migrate their existing workloads running on Kubernetes to the service mesh. In such cases, you will have to manually inject Envoy sidecars to existing services.

To simulate this scenario, let’s deploy a service to our cluster with automatic sidecar injection turned off. Explere the specification in the `book-club-without-istio.yaml` file that we will use to deploy a stateless service to our cluster.

Let's now deploy this service to our cluster with the following command.

`kubectl apply -f my-workshop/book-club-without-istio.yaml`{{execute}}

Once deployed, you can check the containers created in the pods of this service with the following command.

`kubectl get pods --selector app=bookclub -n micro-shake-factory -o jsonpath={.items[*].spec.containers[*].name}`{{execute}}
