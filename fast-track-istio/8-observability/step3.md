Prometheus is a time-series database and visualization tool. It is an excellent tool for visualizing platform and custom metrics. In general, you will find that organizations use both Prometheus and Grafana for gaining very high visibility into the applications and the platform hosting those applications.

As we did for Jaeger, we will expose the Prometheus service over the internet through the following specifications:

`kubectl apply -f my-workshop/prometheus.yaml`{{execute}}

Let's now navigate to the following URL to explore the dashboard.

https://[[HOST_SUBDOMAIN]]-15443-[[KATACODA_HOST]].environments.katacoda.com/

Next, let's look at the companion visualization tool of Prometheus, Grafana. Before proceeding, delete the resources that you created.

`kubectl delete -f my-workshop/prometheus.yaml`{{execute}}

> We have to delete the resources solely to reuse the gateway port, which is limited in number. You can increase the ports available in the cluster by patching the istio-ingressgateway service, but it is out of the scope of this workshop. To check the ports that are available by default in your cluster, execute the command: `kubectl get svc/istio-ingressgateway -n istio-system -o jsonpath={.spec.ports[*].targetPort}`
