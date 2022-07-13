### Create new Image Registry Secret

```shell
# kubectl create secret docker-registry regcred --docker-server=192.168.110.72 --docker-username=xxx --docker-password=xxx --docker-email=xxx@gih.cn
```

### Create a Secret based on existing credentials

```shell
# kubectl create secret generic regcred --from-file=.dockerconfigjson=/root/.docker/config.json --type=kubernetes.io/dockerconfigjson

# cat regcred.yaml
apiVersion: v1
kind: Secret
metadata:
  name: regcred
  namespace: awesomeapps # control the namespace
data:
  .dockerconfigjson:UmVhbGx5IHJl... # base64 encode ~/.docker/config.json 
type: kubernetes.io/dockerconfigjson
```

