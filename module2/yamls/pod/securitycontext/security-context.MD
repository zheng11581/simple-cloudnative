### Create Pod with SecurityContext

```shell
# kubectl apply -f security-context.yaml
# kubectl get pod security-context-demo
```

### Enter the container

```shell
# kubectl exec -it security-context-demo -- sh
# ps -ef
PID   USER     TIME  COMMAND
    1 1000      0:00 sleep 1h # uid=1000
    8 1000      0:00 sh
   14 1000      0:00 ps -ef
# cd /data/
# ls -l
total 0
drwxrwsrwx    2 root     2000             6 Jul 20 09:13 demo # groups=2000
# id
uid=1000 gid=3000 groups=2000
# cd demo/
# echo hello > testfile
# ls -l
total 4
-rw-r--r--    1 1000     2000             6 Jul 20 09:40 testfile
```