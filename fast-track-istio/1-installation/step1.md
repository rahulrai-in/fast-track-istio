Your shell is currently running inside a Kubernetes cluster in the stopped state. To start the cluster, we will execute the `launch.sh` Katacoda script now, which will start the cluster.

`launch.sh`{{execute}}

## Health check

Let's now check the status of our cluster by executing the following command.

`kubectl cluster-info`{{execute}}

Let's now install Istio on our cluster.
