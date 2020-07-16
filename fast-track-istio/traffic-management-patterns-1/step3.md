It is important to test individual microservices of the application for resiliency to ensure that the application degrades gracefully if one or more of its supporting services behave inconsistently. Istio allows you to configure faults for some percentage of HTTP traffic. You can inject arbitrary delays or return specific response codes (such as 500) and use the failures to write integration tests that ascertain the application behavior in the presence of failures of its dependencies.

Let’s introduce deliberate failures in the traffic routed to version 1 of the **movies** service. Apply the following configuration to update the **movies** virtual service.

`kubectl apply -f my-workshop/fault-injection/fruits-api-dr-vs-gw-test-fault.yml`{{execute}}
