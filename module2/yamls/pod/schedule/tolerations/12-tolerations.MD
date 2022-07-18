### Taint a node

```shell
# kubectl taint nodes kbadmnode1 team1=:NoSchedule
# kubectl taint nodes kbadmnode2 team1=:NoSchedule
# kubectl taint nodes kbadmnode3 team1=:NoSchedule
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
```

### Create a Pod with tolerations

```shell
# kubectl apply -f 12-with-tolerations.yaml

# kubectl get pod with-tolerations-demo -owide
```


### Clear the taint

```shell
# kubectl taint nodes kbadmnode1 team1=:NoSchedule-
# kubectl taint nodes kbadmnode2 team1=:NoSchedule-
# kubectl taint nodes kbadmnode3 team1=:NoSchedule-
# kubectl taint nodes kbadmnode4 team1=:NoSchedule-
# kubectl taint nodes kbadmnode5 team1=:NoSchedule-

# kubectl get pod without-tolerations-demo -owide
```