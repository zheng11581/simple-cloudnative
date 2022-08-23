## Demo

```yaml vs
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews-route
  namespace: foo
spec:
  hosts:
  - reviews # interpreted as reviews.foo.svc.cluster.local
  http:
  - match:
    - uri:
        prefix: "/wpcatalog"
    - uri:
        prefix: "/consumercatalog"
    rewrite:
      uri: "/newcatalog"
    route:
    - destination:
        host: reviews # interpreted as reviews.foo.svc.cluster.local
        subset: v2
  - route:
    - destination:
        host: reviews # interpreted as reviews.foo.svc.cluster.local
        subset: v1
```

```yaml dr
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: reviews-destination
  namespace: foo
spec:
  host: reviews # interpreted as reviews.foo.svc.cluster.local
  subsets:
  - name: v1
    labels:
      version: v1
  - name: v2
    labels:
      version: v2

```

- DestinationRule一般和VirtualService配合使用来控制请求路由

## Spec details

- host
- trafficPolicy
  - loadBalancer
    - simple
    - consistentHash
    - localityLbSetting
    - warmupDurationSecs
  
  - connectionPool
    - tcp
      - maxConnections
      - connectTimeout
      - tcpKeepalive
  
    - http
      - http1MaxPendingRequests
      - http2MaxRequests
      - maxRequestsPerConnection
      - maxRetries
      - idleTimeout
      - h2UpgradePolicy
      - useClientProtocol
  
  - outlierDetection
    - splitExternalLocalOriginErrors
    - consecutiveLocalOriginFailures	
    - consecutiveGatewayErrors
    - consecutive5xxErrors	
    - interval	
    - baseEjectionTime
    - maxEjectionPercent
    - minHealthPercent

  - tls
  - portLevelSettings
  - tunnel

- subsets
  - name	
  - labels
  - trafficPolicy

- exportTo
- workloadSelector
- port
