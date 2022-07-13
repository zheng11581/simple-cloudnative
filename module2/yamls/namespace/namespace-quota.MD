### Create Foo namespace
```shell
# kubectl apply -f yamls/namespace/nginx-foo.yaml
# kubectl get pod -n foo -owide
NAME         READY   STATUS    RESTARTS   AGE   IP               NODE         NOMINATED NODE   READINESS GATES
nginx-demo   2/2     Running   0          10s   192.168.176.54   cn-master1   <none>           <none>
# curl 192.168.176.54
foo
```

### Create Bar namespace
```shell
# kubectl apply -f yamls/namespace/nginx-bar.yaml
# kubectl get pod -n bar -owide
NAME         READY   STATUS    RESTARTS   AGE   IP               NODE         NOMINATED NODE   READINESS GATES
nginx-demo   2/2     Running   0          12s   192.168.176.55   cn-master1   <none>           <none>

# curl 192.168.176.55
bar
```

### Resources Quota: Limit 1 Pod in Foo & Bar namespace
```shell
# kubectl apply -f  yamls/namespace/quota.yaml
# kubectl get resourcequotas

NAME            AGE   REQUEST                      LIMIT
object-counts   55s   configmaps: 1/1, pods: 1/1  

# kubectl apply -f yamls/namespace/execeed-pod.yaml
Error from server (Forbidden): error when creating "yamls/namespace/execeed-pod.yaml": pods "pod-execeed" is forbidden: exceeded quota: object-counts, requested: pods=1, used: pods=1, limited: pods=1
Error from server (Forbidden): error when creating "yamls/namespace/execeed-pod.yaml": pods "pod-execeed" is forbidden: exceeded quota: object-counts, requested: pods=1, used: pods=1, limited: pods=1

```
