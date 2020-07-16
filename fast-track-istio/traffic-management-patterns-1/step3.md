It is important to test individual microservices of the application for resiliency to ensure that the application degrades gracefully if one or more of its supporting services behave inconsistently. Istio allows you to configure faults for some percentage of HTTP traffic. You can inject arbitrary delays or return specific response codes (such as 500) and use the failures to write integration tests that ascertain the application behavior in the presence of failures of its dependencies.

Let’s introduce deliberate failures in the traffic routed to version 1 of the **movies** service. Open the specification in the `my-workshop/fault-injection/movies-gw-vs-dr-fault.yaml` file in your editor. In the virtual service specification, notice the following excerpt.

<pre>
http:
  - match:
      - headers:
          version:
            exact: "2"
    route:
      - destination:
          host: movies-api-service
          port:
            number: 3000
          subset: v1
  - route:
      - destination:
          host: movies-api-service
          port:
            number: 3000
          subset: v1
    fault:
      delay:
        fixedDelay: 5s
        percentage:
          value: 40
      abort:
        httpStatus: 500
        percentage:
          value: 60
</pre>

The above configuration causes 40 percent of the requests to the v1 service to process with a delay of 5 seconds, and 60 percent of the requests to fail with HTTP status code 500.

To test this configuration, let's apply the specification so that it updates the **movies** service.

`kubectl apply -f my-workshop/fault-injection/movies-gw-vs-dr-fault.yaml`{{execute}}

After waiting a couple of seconds, we will execute a simple shell script that sends 10 requests to the service and outputs the response received.

`for ((i=1;i<=10;i++)); do curl --write-out "Request #$i:%{http_code}\n" --silent --output /dev/null "http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/"; done`{{execute}}

You will notice that because of the applied configuration, the script takes some time before returning the 200/OK response, while it returns immediately for errors. Note that you won’t always get the desired split of errors and successes, since the faults are averaged over time.
