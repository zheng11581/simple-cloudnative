
## Create ConfigMap and Secret
```shell
# kubectl apply -f cm.yml
# kubectl apply -f secret.yml
```

## Inject ConfigMap with ENV
```shell
# kubectl explain pod.spec.containers.env.valueFrom

# kubectl apply -f env-pod.yml
# kubectl get pod
# kubectl exec -it env-pod -- sh

# in Pod:
# echo $COUNT
# echo $GREETING
# echo $USERNAME
# echo $PASSWORD
```

## Inject ConfigMap with VOLUME
```shell
# kubectl explain pod.spec.containers.env.valueFrom

# kubectl explain pod.spec.volumes
# kubectl explain pod.spec.containers.volumeMounts

# kubectl apply -f vol-pod.yml
# kubectl get pod
# kubectl exec -it vol-pod -- sh

# in Pod:
# cat /tmp/cm-items/greeting
# cat /tmp/sec-items/db
```