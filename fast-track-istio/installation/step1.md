## Start the cluster

Your shell is currently running inside a Kubernetes cluster. To start the cluster, we are required to run a Katacoda script. This will launch a two-node Kubernetes cluster with one master and one node.

`launch.sh`{{execute}}

#### Health Check

Once started, you can get the status of the cluster with `kubectl cluster-info`{{execute}}
