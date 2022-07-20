### 1. Rediness Probe

Create pod readiness-demo

```shell
# kubectl apply -f 7-readiness-probe.yaml
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

Describe the Pod

```shell
# kubectl describe pod liveness-demo
...
Events:
  Type     Reason     Age                    From               Message
  ----     ------     ----                   ----               -------
  Normal   Scheduled  8m40s                  default-scheduler  Successfully assigned default/liveness-demo to kbadmnode4
  Normal   Pulled     8m38s                  kubelet            Successfully pulled image "centos" in 773.672774ms
  Normal   Pulled     7m12s                  kubelet            Successfully pulled image "centos" in 15.509891433s
  Normal   Created    5m42s (x3 over 8m38s)  kubelet            Created container liveness
  Normal   Started    5m42s (x3 over 8m38s)  kubelet            Started container liveness
  Normal   Pulled     5m42s                  kubelet            Successfully pulled image "centos" in 15.512520281s
  Normal   Killing    4m58s (x3 over 7m58s)  kubelet            Container liveness failed liveness probe, will be restarted
  Normal   Pulling    4m28s (x4 over 8m39s)  kubelet            Pulling image "centos"
  Warning  Unhealthy  3m38s (x10 over 8m8s)  kubelet            Liveness probe failed: cat: /tmp/healthy: No such file or directory

```


### 3. Lifecycle postStart

Create pod poststart-demo

```shell
# kubectl apply -f 7-poststart.yaml

# kubectl apply -f 7-poststart.yaml
NAME             READY   STATUS    RESTARTS   AGE    IP              NODE    NOMINATED NODE   READINESS GATES
poststart-demo   1/1     Running   0          118s   10.244.104.29   node2   <none>           <none>

# curl 10.244.104.29

Hello from the postStart handler
```

### 4. Rediness gate

Create pod readiness-gate and svc readiness-gate

```shell
# kubectl apply -f 7-readiness-gates.yaml
```

Curl the Nginx pod，succeed

```shell
# kubectl get pod readiness-gate -owide
NAME             READY   STATUS    RESTARTS   AGE     IP               NODE    NOMINATED NODE   READINESS GATES
readiness-gate   1/1     Running   0          5m48s   10.244.166.185   node1   <none>           0/1

# curl 10.244.166.185
...
<title>Welcome to nginx!</title>
...

# kubectl get svc readiness-gate -owide
NAME             TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE     SELECTOR
readiness-gate   ClusterIP   10.105.92.39   <none>        80/TCP    9m36s   app=readiness-gate

# curl 10.105.92.39
curl: (7) Failed to connect to 10.105.92.39 port 80: Connection refused
```

