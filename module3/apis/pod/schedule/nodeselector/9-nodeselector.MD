### Add a label to a node

1. List the nodes in your cluster, along with their labels

```shell
# kubectl get nodes --show-labels
master1   Ready    control-plane,master   16d   v1.21.14   kubernetes.io/hostname=master1
node1     Ready    <none>                 16d   v1.21.14   kubernetes.io/hostname=node1
node2     Ready    <none>                 16d   v1.21.14   kubernetes.io/hostname=node2
```

2. Choose one of your nodes, and add a label to it

```shell
# kubectl label nodes node2 disktype=ssd
# kubectl label nodes node1 disktype=hdd
```

3. Verify that your chosen node has a disktype=ssd label

```shell
# kubectl get nodes node2 --show-labels
NAME    STATUS   ROLES    AGE   VERSION    LABELS
node2   Ready    <none>   16d   v1.21.14   disktype=ssd,kubernetes.io/hostname=node2
```

### Create a pod that gets scheduled to your chosen node

1. Use the configuration file to create a pod that will get scheduled on your chosen node
```shell
# kubectl apply -f 9-nodeselector.yaml
```

2. Verify that the pod is running on your chosen node
```shell
# kubectl get pods -owide
```

