### Create Pod with two InitContainers
```shell
# kubectl apply -f yamls/pod/container/2-init-container.yaml
# kubectl get -f yamls/pod/container/2-init-container.yaml
NAME         READY   STATUS     RESTARTS   AGE   IP             NODE    NOMINATED NODE   READINESS GATES
initc-demo   0/1     Init:0/2   0          38s   10.244.104.4   node2   <none>           <none>
```

### Complete the InitCOntainers
```shell
# create file1 in node2
# touch /initor/file1

# kubectl get -f yamls/pod/container/2-init-container.yaml
NAME         READY   STATUS     RESTARTS   AGE   IP             NODE    NOMINATED NODE   READINESS GATES
initc-demo   0/1     Init:1/2   0          85s   10.244.104.4   node2   <none>           <none>

# create file1 in node2
# touch /initor/file2

# kubectl get -f yamls/pod/container/2-init-container.yaml
NAME         READY   STATUS    RESTARTS   AGE    IP             NODE    NOMINATED NODE   READINESS GATES
initc-demo   1/1     Running   0          108s   10.244.104.4   node2   <none>           <none>
```