Istio has a set of in-built configuration profiles that can be used to specify the Istio components that you wish to install on your cluster. You can also specify the list of components that you want to install using the long form of the operator specification. For example, the following specification will install Grafana on your cluster.

`
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: istio-system
  name: example-istiocontrolplane
spec:
  profile: default
  addonComponents:
    grafana:
      enabled: true
`

[Here is the list](https://istio.io/latest/docs/setup/additional-setup/config-profiles/) of components that are installed when you specify an installation profile.

In the previous step, we installed Istio using the `minimum` configuration, which installs only *istiod* on your cluster. Let's update the installation specification to install all the components using the `demo` profile.

`kubectl apply -f - <<EOF
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: istio-system
  name: example-istiocontrolplane
spec:
  profile: demo
EOF`{{execute}}

Let's check the services deployed by the operator by executing the following command.

`kubectl get svc -n istio-system`{{execute}}

Again, let's watch the health of Istio control plane resources (pods, deployment, services) deployed on our cluster.

`watch -n .5 kubectl get pods,deploy,svc -o wide -n istio-system`{{execute}}

Once all the resources are running, press "CTRL+C" to exit the watch. Excellent! You successfully installed Istio and customized the installation to install the components that you need.