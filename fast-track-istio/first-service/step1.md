Let's prepare our environment for working through this workshop. In the last scenario, you installed Istio on your cluster. We will download and execute a script that will start the Katacoda Kubernetes cluster and install Istio on it. Execute the following command to download the files that we will use in this scenario.

`git clone https://github.com/rahulrai-in/fast-track-istio.git; mv fast-track-istio/scripts/scenario-2 my-workshop; rm -rf fast-track-istio`{{execute}}

Let's start our cluster and install Istio on it with the following command.


Kubernetes objects support having user-defined labels, which are essentially key-value pairs, attached to them.

Istio relies on a specific label named istio-injection to decide the namespace on which the Envoy proxies are applied. This label supports two values. Setting the value to _enabled_ means that Istio automatically deploys sidecars for the pods of your service. On the other hand, setting this value to _disabled_ means that Istio will not inject sidecars automatically; however, it won’t affect the services within the namespace that have sidecars attached to them.

The journey to virtualization in enterprises is an incremental process. This means that in some cases, you might find that organizations want to bring their existing workloads running on Kubernetes to the mesh. In such cases, you will have to manually inject Envoy sidecars into the existing services.

To simulate this scenario, let’s deploy a service to our cluster with automatic sidecar injection turned off.

`launch.sh`{{execute}}

## Health check

Let's now check the status of our cluster by executing the following command.

`kubectl cluster-info`{{execute}}

Let's now install Istio on our cluster.
