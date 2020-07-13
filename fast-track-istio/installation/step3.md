## Install options

Istio has a set of in-built configuration profiles that can be used to specify the Istio components that you wish to install on your cluster. You can also specify the list of components that you want to install using the long form of the operator specification. For example, the following specification will install grafana on your cluster.

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

In the previous step, we installed Istio using the `demo` configuration, which installs all the components. Let's update the installation to install the least number of components through the `minimum` profile.

`kubectl apply -f - <<EOF
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: istio-system
  name: example-istiocontrolplane
spec:
  profile: minimum
EOF`{{execute}}

Let's check the services deployed by the operator by executing the following command.

`kubectl get svc -n istio-system`{{execute}}

Again, let's watch the health of Istio control plane resources (pods, deployment, services) deployed on our cluster.

watch -n .5 kubectl get pods,deploy,svc -o wide -n istio-system{{execute}}

Once all the resources are running, press "CTRL+C" to exit the watch. Excellent! You successfully installed Istio and customized the installation.