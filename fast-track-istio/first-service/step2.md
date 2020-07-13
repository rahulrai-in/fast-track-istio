You can find platform-specific installation instructions for Istio [here](https://istio.io/latest/docs/setup/install).

The first component that we need to install is the Istio CLI, known as `istioctl`. To install `istioctl`, execute the following command.

`curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.6.5 sh -;cd istio-1.6.5;export PATH=$PWD/bin:$PATH`{{execute}}

Next, install the Istio operator with the following command.

`istioctl operator init`{{execute}}

Let's now create a namespace for Istio resources with the following command.

`kubectl create ns istio-system`{{execute}}

The following command will install only the Istio daemon-*istiod*  on your cluster.

`kubectl apply -f - <<EOF
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: istio-system
  name: example-istiocontrolplane
spec:
  profile: minimum
EOF`{{execute}}

You can check the services deployed by the operator by executing the following command.

`kubectl get svc -n istio-system`{{execute}}

Let's check the health of Istio control plane resources (pods, deployment, services) deployed on our cluster.

`watch -n .5 kubectl get pods,deploy,svc -o wide -n istio-system`{{execute}}

Once all the resources are running, press "CTRL+C" to exit the watch.

Let's explore how we can customize the installation next.