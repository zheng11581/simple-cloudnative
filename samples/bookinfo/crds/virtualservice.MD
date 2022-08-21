## Demo

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews-route
spec:
  hosts:
  - reviews.bookinfo.svc.cluster.local
  gateways:
  - mesh
  http:
  - name: "reviews-v2-routes"
    match:
    - uri:
        prefix: "/wpcatalog"
    - uri:
        prefix: "/consumercatalog"
    rewrite:
      uri: "/newcatalog"
    route:
    - destination:
        host: reviews.bookinfo.svc.cluster.local
        subset: v2
  - name: "reviews-v1-route"
    route:
    - destination:
        host: reviews.bookinfo.svc.cluster.local
        subset: v1

```

- 下游客户端通过HTTP协议调用上游服务reviews，定义两条路由规则，规则会下发到网格内部
- 当匹配到前缀是/wpcatalog或者/consumercatalog的URI时，前缀将被重写到/newcatalog，并将流量路由到reviews的v1服务子集
- 路由匹配的顺序是从上到下，最后一个是默认的路由规则
- 默认路由规则是将流量路由到reviews服务的v1子集


## Spec details

- hosts

- gateways

- http

    - name
    - match
    - route
    - redirect
    - delegate
    - rewrite
    - timeout
    - retries
    - fault
    - mirror
    - mirrorPercentage
    - corsPolicy
    - headers


- tls

- tcp

- exportTo
