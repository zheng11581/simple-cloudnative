## 为什么需要 ConfigMap Secret

## Create ConfigMap and Secret
```shell
# kubectl apply -f cm.yaml
# kubectl apply -f secret.yaml
```

## Inject ConfigMap and Secret with ENV
```shell
# kubectl explain pod.spec.containers.env.valueFrom

# kubectl apply -f env-pod.yaml
# kubectl get pod
# kubectl exec -it env-pod -- sh

# in Pod:
# echo $COUNT
# echo $GREETING
# echo $USERNAME
# echo $PASSWORD
```

## Inject ConfigMap and Secret with VOLUME
```shell
# kubectl explain pod.spec.containers.env.valueFrom

# kubectl explain pod.spec.volumes
# kubectl explain pod.spec.containers.volumeMounts

# kubectl apply -f vol-pod.yaml
# kubectl get pod
# kubectl exec -it vol-pod -- sh

# in Pod:
# cat /tmp/cm-items/greeting
# cat /tmp/sec-items/db
```