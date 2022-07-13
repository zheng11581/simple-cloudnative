### Two containers Pod
```shell 
# cat yamls/two-container-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: two-containers
spec:
  restartPolicy: Never
  volumes:
  - name: shared-data
    hostPath:      
      path: /mydata
  containers:
  - name: nginx-container
    image: bitnami/nginx:1.16-centos-7
    volumeMounts:
    - name: shared-data
      mountPath: /usr/share/nginx/html
  - name: centos-container
    image: centos:7.9.2009
    volumeMounts:
    - name: shared-data
      mountPath: /pod-data
    command: 
    - /bin/sh
    args: 
    - -c
    - echo Hello from the centos container > /pod-data/index.html; sleep 3600
# kubeclt apply -f yamls/two-container-pod.yaml
```

### Share same Network namespace
```shell
# kubectl exec -it two-containers -c centos-container -- hostname -i
# kubectl exec -it two-containers -c nginx-container -- hostname -i
```