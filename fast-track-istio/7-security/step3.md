In this exercise, we will enforce access control on an Istio ingress gateway using an authorization policy. There are several types of authorization policies that you can apply on the ingress gateway, such as IP based allow list and deny list. To demonstrate the flexibility of the authorization policy, we will use this feature to block access to the **Books** service from the `curl` user agent.

To begin, let's reapply the specifications for the **Books** service and deployment.

`kubectl apply -f my-workshop/books-api.yaml`{{execute}}

Let's now create a gateway and expose our service to the internet with specifications in the `my-workshop/books-api-vs-gw.yaml` file. You will find that the policy `AuthorizationPolicy` in the specification is commented out. We will see the effect of applying this policy shortly.

`kubectl apply -f my-workshop/books-api-vs-gw.yaml`{{execute}}

Let's invoke our API through our browser by vising this URL.

http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/books/1

Also, let's invoke the same endpoint through the cURL command.
`curl http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/books/1 -v`{{execute}}

Let's now uncomment the specification for `AuthorizationPolicy` in the `my-workshop/books-api-vs-gw.yaml` file. Let's explore the specification in a little detail.

<pre>
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: ingress-policy
  namespace: istio-system
spec:
  selector:
    matchLabels:
      app: istio-ingressgateway
  action: DENY
  rules:
  - when:
     - key: request.headers[User-Agent]
       values: ["curl/*"]
</pre>

The authorization policy is made up of rules which are evaluated on every request. When the rules evaluate to true, then the action is triggered. As you can see in the specification, if the User-Agent that is making the request is any version of curl, the request is denied.

Apply the specification with the following command.

`kubectl apply -f my-workshop/books-api-vs-gw.yaml`{{execute}}

Again, let's invoke our API through our browser by vising this URL.

http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/books/1

Also, let's invoke the same endpoint through the cURL command.
`curl http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/books/1 -v`{{execute}}

You will find that the request made through the browser succeeds, but the one made through the curl utility fails.

## Conclusion

Congratulations! You successfully implemented a couple of security policies on the workloads in the mesh.
