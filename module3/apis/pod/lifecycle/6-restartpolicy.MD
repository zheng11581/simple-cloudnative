### Apply the Pod 

```shell
# kubectl apply -f 6-pod-restartpolicy.yaml
```

### Watch the Pod never-restart，no restart event

```shell
# kubectl get pod never-restart -w
```

### Watch the Pod always-restart，watch restart event

```shell
# kubectl get pod always-restart -w
```
