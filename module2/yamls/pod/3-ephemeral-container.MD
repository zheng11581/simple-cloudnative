### First, create a pod for the example
```shell
# kubectl run ephemeral-demo --image=ibmcom/pause:3.1 --restart=Never
# kubectl get pod -owide
NAME             READY   STATUS    RESTARTS   AGE   IP               NODE    NOMINATED NODE   READINESS GATES
ephemeral-demo   1/1     Running   0          92s   10.244.166.140   node1   <none>           <none>

# kubectl exec -it ephemeral-demo -- sh
OCI runtime exec failed: exec failed: unable to start container process: exec: "sh": executable file not found in $PATH: unknown
command terminated with exit code 126

# kubectl debug -it ephemeral-demo --image=busybox:1.28 --target=ephemeral-demo
```