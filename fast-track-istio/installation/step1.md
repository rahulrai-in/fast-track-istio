## Start the cluster

Your shell is currently running inside a Kubernetes cluster in the stopped state. To start the cluster, we will execute the `launch.sh` Katacoda script now, which will start the cluster.

`launch.sh`{{execute}}

## Health check

Let's now check the status of our cluster by executing the following command.

`kubectl cluster-info`{{execute}}

## Download the source code

Let's download the scripts that we will use in this scenario from GitHub.

`git clone https://github.com/rahulrai-in/fast-track-istio.git; mv fast-track-istio/scripts/scenario-1 my-workshop; rm -rf fast-track-istio`{{execute}}

In the editor window, you will now find a folder named _my-workshop_ that contains the scripts that we will use in this scenario.

Let's now install Istio on our cluster.
