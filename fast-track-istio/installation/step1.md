## Start the cluster

Your shell is currently running inside a Kubernetes 1.18 cluster. However, your cluster is not running currently. To start the cluster, we will execute the `launch.sh` Katacoda script. This will launch a two-node Kubernetes cluster with one master and one node.

`launch.sh`{{execute}}

## Health check

Let's now check the status of our cluster by executing this command: `kubectl cluster-info`{{execute}}

## Download the source code

Let's download the scripts that we will use in this workshop from GitHub.

`ssh root@host01 "git clone https://github.com/rahulrai-in/fast-track-istio.git; mv fast-track-istio/scripts myworkshop; rm -rf fast-track-istio"`{{execute}}

## Install Istio
