### Deploy nginx with 5 replicas

```shell
# kubectl apply -f 1-replicas-selector.yaml
```


### kubectl filter pod labels

```shell
# kubectl get pod -l app=nginx --show-labels
```

### Expose nginx service to the cluster

```shell
# kubectl apply -f 2-endpoint-selector.yaml
```