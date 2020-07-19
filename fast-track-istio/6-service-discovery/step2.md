## Service Entry

Using Service Entry, you can add or remove a service from the service registry of Istio. Once defined, the service names become well known to Istio, and they can then be mapped to other Istio configurations. By default, Istio allows services to send requests to hostnames not present in the service registry.

Creating a service entry does not create a DNS record in Kubernetes, so you won't be able to perform a DNS lookup from the application to fetch the resolved IP address. A CoreDNS Istio plugin can generate DNS records from service entry records, enabling Istio to populate DNS records outside Istio. With service entry records, you can implement traffic management patterns such as timeouts on external services.

Let's first declare a service entry specification to add DNS records for the **Independent** service. The following is a part of the specification that you will find in the file `my-workshop/independent-se-dr.yaml`.

<pre>
spec:
  hosts:
    - independent.australiaeast.azurecontainer.io
  location: MESH_EXTERNAL
  ports:
    - number: 8080
      name: http
      protocol: HTTP
  resolution: DNS
</pre>

> Remember to change the value of the key `hosts` in the specification with the actual DNS name of your service.

The previous specification will create a record in the Istio service registry. Let's create a destination rule to ensure that other services communicate with our service over TLS.

<pre>
spec:
  host: independent.australiaeast.azurecontainer.io
  trafficPolicy:
     loadBalancer:
       consistentHash:
         httpCookie:
           name: user
           ttl: 0s
</pre>

With the previous configuration, the destination rule will use a cookie to route consecutive requests to the same instance of the **Independent** service. If the cookie does not already exist, it will create one and return it with the response. The client should pass the same cookie in all consecutive requests to reach the same destination service.

Apply the policy to your cluster with the following command.

`kubectl apply -f my-workshop/independent-se-dr.yaml`{{execute}}

We are now ready to test the new routing capabilities now that the **Independent** service is now native to our service mesh.

## Testing with Busybox

You can inject pods in your namespace for debugging services within the mesh. The [Radial Busyboxplus](https://hub.docker.com/r/radial/busyboxplus) image consists of a few useful DNS utilities that you can use to debug your services in any environment. You can inspect the specification for the pod in the file `my-workshop/busybox.yaml`. Let's apply the specification to our cluster.

`kubectl apply -f my-workshop/busybox.yaml`{{execute}}

Let's now execute the following command to send an HTTP GET request to the **Independent** service, which should get routed via the destination rule that we previously created. The Independent service does not return any cookies in response, so the only cookie that we receive must be from the configuration that we applied.

`kubectl exec -ti -n fast-track-istio busybox -- curl http://independent.australiaeast.azurecontainer.io:8080/ -v`{{execute}}

We applied a load balancer policy to the service here. You may enforce other policies such as connection pooling, TLS management, and port-specific settings to the target service as well.

## Workload Entry

A Workload Entry along with a service entry allows you to configure Clusters in Envoy. A Cluster is a group of upstream targets (or hosts) where the traffic has to be routed based on certain match conditions. After you create a workload in Istio, it is treated the same way as any other service on the mesh. With workload entries, you can balance network traffic between VMs and pods similarly.
