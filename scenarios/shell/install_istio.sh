curl -L https://istio.io/downloadIstio | sh -
cd istio-1.6.5
export PATH=$PWD/bin:$PATH
istioctl install --set profile=demo
kubectl label namespace default istio-injection=enabled