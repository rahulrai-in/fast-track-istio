curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.6.5 sh -
cd istio-1.6.5
export PATH=$PWD/bin:$PATH
istioctl install --set profile=demo