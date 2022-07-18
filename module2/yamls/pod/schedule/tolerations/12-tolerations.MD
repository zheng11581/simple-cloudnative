### Taint a node

```shell
# kubectl taint nodes kbadmnode1 team1=:NoSchedule
# kubectl taint nodes kbadmnode2 team1=:NoSchedule
# kubectl taint nodes kbdmnode03 team1=:NoSchedule
# kubectl taint nodes kbadmnode4 team1=:NoSchedule
# kubectl taint nodes kbadmnode5 team1=:NoSchedule

# kubectl get node kbadmnode1 -oyaml

  taints:
  - effect: NoSchedule
    key: team1
```

### Create a Pod without tolerations

```shell
# kubectl apply -f 12-without-tolerations.yaml

# kubectl get pod without-tolerations-demo -owide
NAME                       READY   STATUS    RESTARTS   AGE   IP       NODE     NOMINATED NODE   READINESS GATES
without-tolerations-demo   0/1     Pending   0          7s    <none>   <none>   <none>           <none>

# kubectl describe pod without-tolerations-demo 
Events:
  Type     Reason            Age   From               Message
  ----     ------            ----  ----               -------
  Warning  FailedScheduling  29s   default-scheduler  0/7 nodes are available: 2 node(s) had taint {node-role.kubernetes.io/master: }, that the pod didn't tolerate, 5 node(s) had taint {team1: }, that the pod didn't tolerate.
  Warning  FailedScheduling  28s   default-scheduler  0/7 nodes are available: 2 node(s) had taint {node-role.kubernetes.io/master: }, that the pod didn't tolerate, 5 node(s) had taint {team1: }, that the pod didn't tolerate.
```

### Create a Pod with tolerations

```shell
# kubectl apply -f 12-with-tolerations.yaml

# kubectl get pod with-tolerations-demo -owide
NAME                    READY   STATUS    RESTARTS   AGE   IP               NODE         NOMINATED NODE   READINESS GATES
with-tolerations-demo   1/1     Running   0          6s    10.244.179.190   kbadmnode4   <none>           <none>

```


### Clear the taint

```shell
# kubectl taint nodes kbadmnode1 team1=:NoSchedule-
# kubectl taint nodes kbadmnode2 team1=:NoSchedule-
# kubectl taint nodes kbdmnode03 team1=:NoSchedule-
# kubectl taint nodes kbadmnode4 team1=:NoSchedule-
# kubectl taint nodes kbadmnode5 team1=:NoSchedule-

# kubectl get pod without-tolerations-demo -owide
```