Using Service Entry, you can add or remove a service from the service registry of Istio. Once defined, the service names become well known to Istio, and they can then be mapped to other Istio configurations. By default, Istio allows services to send requests to hostnames that are not present in the service registry. To reduce the attack surface area of the application, this setting should be inverted, and only legitimate services outside the cluster should be whitelisted by creating service entries for them.

Creating a service entry does not create a DNS record in Kubernetes, so you won’t be able to perform a DNS lookup from the application to fetch the resolved IP address. A CoreDNS Istio plugin can generate DNS records from service entry records, which enables Istio to populate DNS records outside Istio. With service entry records, you can implement traffic management patterns such as timeouts on external services as well.

A Workload Entry along with a service entry allows you to configure Clusters in Envoy. A Cluster is a group of upstream targets (or hosts) where the traffic has to be routed based on certain match conditions. After you create a workload in Istio, it is treated the same way as any other service on the mesh. Let's use service entry and workload entry to bring our externally hosted **Independent** service to the mesh.

Let's first declare a service entry specification to add DNS records for the **Independent** service. The following is a part of the specification that you will find in the file `my-workshop/independent-se-we.yaml`.

<pre>
spec:
  hosts:
    - independent.australiaeast.azurecontainer.io
  location: MESH_INTERNAL
  ports:
    - number: 8080
      name: http
      protocol: HTTP
  resolution: DNS
  workloadSelector:
    labels:
      app: independent
</pre>

> Remember to change the value of the key `hosts` in the specification with the actual DNS name of your service.

Let's now create a workload entry for the service. The workload entry specification is located in the `my-workshop/independent-se-we.yaml` as well.

https://medium.com/rapido-labs/kubernetes-istio-and-the-world-outside-rapido-75da3666db4a

https://stackoverflow.com/questions/61980095/what-is-the-purpose-of-a-virtualservice-when-defining-an-wildcard-serviceentry-i
