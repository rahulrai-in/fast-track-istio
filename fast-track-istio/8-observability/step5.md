Visualizing how your mesh is performing at runtime can help give you absolute control over your mesh. Kiali (a Greek word that means "spyglass") is an open-source and versatile visualization tool that pulls data from Prometheus and the host Kubernetes to generate a communication graph of the mesh that shows service-to-service interactions. With Kiali, since the entire communication stack (Istio and Kubernetes) is available to you, you get much better visibility of the system with Kiali than with Grafana.

We will expose the Kiali service over the internet through the following specifications:

`kubectl apply -f my-workshop/kiali.yaml`{{execute}}

Let's now navigate to the following URL to explore the dashboard.

https://[[HOST_SUBDOMAIN]]-15443-[[KATACODA_HOST]].environments.katacoda.com/

You will be asked to enter a username and password, which would be **admin** for both the fields. On the Overview dashboard, you will see all the applications that are executing in your mesh. You can click on any application to view the health of the services in that application. For each application, the dashboard also shows incoming and outgoing traffic metrics. You can bring up this dashboard and send some traffic to the service in the background to light it up.

Our favorite section in Kiali is the **Istio Config** tab that can show you any misconfigurations in your mesh. This feature can surface issues such as virtual services listening to a gateway that does not exist, routes that do not exist, multiple virtual services for the same host, and non-existent service subsets.

## Conclusion

This exercise concludes our journey of learning Istio. We hope that you enjoyed learning Istio with us and that we were able to ignite the desire to explore Istio in you. Thank you for being with us on this learning journey.
