launch.sh

# Install istioctl
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.6.5 sh -
export PATH="$PATH:/root/istio-1.6.5/bin"

# Install Istio operator
istioctl operator init
kubectl create ns istio-system
kubectl apply -f - <<EOF
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: istio-system
  name: example-istiocontrolplane
spec:
  profile: demo
EOF
