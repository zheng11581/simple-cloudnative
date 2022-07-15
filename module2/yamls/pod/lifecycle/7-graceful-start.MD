### 1. Rediness Probe

Create pod readiness-demo

```shell
# kubectl apply -f 7-rediness-probe.yaml
# kubectl get pod readiness-demo
NAME             READY   STATUS    RESTARTS   AGE
readiness-demo   0/1     Running   0          43s

```

Describe the Pod

```shell
# kubectl describe pod readiness-demo
...
Events:
  Type     Reason     Age                   From               Message
  ----     ------     ----                  ----               -------
  Normal   Scheduled  6m34s                 default-scheduler  Successfully assigned default/readiness-demo to node1
  Normal   Pulling    6m33s                 kubelet            Pulling image "nginx"
  Normal   Pulled     6m32s                 kubelet            Successfully pulled image "nginx" in 880.028674ms
  Normal   Created    6m32s                 kubelet            Created container nginx-demo
  Normal   Started    6m32s                 kubelet            Started container nginx-demo
  Warning  Unhealthy  89s (x55 over 5m59s)  kubelet            Readiness probe failed: HTTP probe failed with statuscode: 404

```

### 2. Liveness Probe

Create pod liveness-demo

```shell
# kubectl apply -f 7-liveness-probe.yaml
# kubectl get pod liveness-demo
```
