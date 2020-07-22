Prometheus is a time-series database and visualization tool. For an overall view of the mesh, Istio supports the Grafana visualization add-on. Grafana is a popular open-source metrics visualization tool that can be used to query, analyze, and alert on metrics. The Istio deployment of Grafana consists of some common dashboards. Grafana is dependent on Prometheus for metrics.

You know the drill. We will expose the Grafana service over the internet through the following specifications:

`kubectl apply -f my-workshop/grafana.yaml`{{execute}}

Let's now navigate to the following URL to explore the dashboard.

https://[[HOST_SUBDOMAIN]]-15443-[[KATACODA_HOST]].environments.katacoda.com/

Next, let's look at another visualization tool powered by Prometheus, Kiali. Before proceeding, delete the resources that you created.

`kubectl delete -f my-workshop/grafana.yaml`{{execute}}
