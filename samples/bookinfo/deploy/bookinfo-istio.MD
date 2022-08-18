## Download Istio

### Manual download 

```shell
https://github.com/istio/istio/releases/tag/1.14.1
```

### Auto download

```shell
# latest
# curl -L https://istio.io/downloadIstio | sh -
# special version
# curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.14.1 TARGET_ARCH=x86_64 sh -

```

### Move to the Istio package directory

```shell
# cd istio-1.14.1

```

### Add the istioctl client to your path

```shell
# export PATH=$PWD/bin:$PATH

```

## Install Istio

```shell
# istioctl install --set profile=demo -y

✔ Istio core installed                          
✔ Istiod installed                                    
✔ Ingress gateways installed                                               
✔ Egress gateways installed                    
✔ Installation complete
```

##  Label namespace

```shell
# kubectl label namespace bookinfo istio-injection=enabled

```

## Delete bookinfo application

```shell
# kubectl delete -f samples/bookinfo/platform/kube/bookinfo.yaml -n bookinfo
```

## Deploy bookinfo application

[Bookinfo](../bookinfo-app.MD)

## Create `gateway` and `virtualservice` for Bookinfo

```shell
# kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml -n bookinfo
# istioctl analyze -n bookinfo
```

## Determining the ingress IP and ports

```shell
# minikube tunnel
# export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
# export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
# export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].port}')

# export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT
```

## Verify external access

```shell
# curl -s http://${GATEWAY_URL}/productpage | grep -o "<title>.*</title>"

```

## Apply default DestinationRule

```shell
# kubectl apply -f samples/bookinfo/networking/destination-rule-all.yaml -n bookinfo

```
