Traffic shifting, or canary deployment, is the practice of routing a small portion of the application traffic to a newly deployed workload to validate the increment quality. As the validations succeed, you can keep ramping up the percentage of traffic to the new workload until all traffic reaches the new workload.

The canary deployment model is a little different from the blue/green deployment model. Unlike routing all the traffic to one of the two replicas of the application (blue and green), which requires double the capacity of deployment, you need a constant amount of additional resources to propagate the changes.

You can direct the traffic to the canary service in several ways, such as by splitting the traffic by percentage or directing traffic from internal users to the canary service. For this demo, we will split the traffic in a 1:10 ratio between version 2 and version 1 of the **movies** service. The following portion of the specification in the `my-workshop/canary/movies-gw-vs-canary.yaml` file is responsible for directing Istio to split the traffic in the said proportions.

<pre>
http:
  - route:
      - destination:
          host: movies-api-service
          port:
            number: 3000
          subset: v1
        weight: 10
      - destination:
          host: movies-api-service
          port:
            number: 3000
          subset: v2
        weight: 90
</pre>

Execute the following script to deploy the two subsets of service.

`kubectl apply -f my-workshop/canary/movies-gw-vs-canary.yaml`{{execute}}

After waiting for some time, execute the following script to dispatch 10 requests to the service.

`for ((i=1;i<=10;i++)); do curl "http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/"; done`

The task to split the canary traffic by HTTP attributes is left to you as an exercise. Let’s discuss the next pattern now.
