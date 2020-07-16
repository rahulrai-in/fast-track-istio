Let's prepare our environment for working through this scenario. In the last scenario, you installed Istio on your cluster. We will download and execute a script that will start the Katacoda Kubernetes cluster and install Istio. Execute the following command to download the exercise files that we will use in this scenario.

`git clone https://github.com/rahulrai-in/fast-track-istio.git; mv fast-track-istio/scripts/scenario-5 my-workshop; rm -rf fast-track-istio`{{execute}}

Let's start our cluster and install Istio on it with the following command.

`. my-workshop/prepare-cluster.sh`{{execute}}

Let's move on to the next step while the script executes.
