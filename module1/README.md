## 模块一：Docker核心技术
## 目标
1. 理解什么是Docker容器
2. Docker的底层技术
3. 将我们的业务容器化，过程需要注意什么

## 什么是Docker容器？
1. 启动一个nginx服务容器看看

```shell
# docker run -d ubuntu/nginx:latest
# ubuntu是镜像仓库的地址
# nginx这个镜像的名字
# latest镜像版本版本

```
2. 自己来做一个 nginx 的镜像

```shell
# cat Dockerfile
FROM ubuntu:impish
RUN apt update && apt-get install -y nginx
COPY file1 /var/www/html/
ADD  file2.tar.gz /var/www/html/
EXPOSE 80
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
# 它提供了容器中程序执行需要的所有文件

# docker build -t registry/nginx:v1 -f ./Dockerfile . 

# docker images
REPOSITORY       TAG       IMAGE ID       CREATED          SIZE
registry/nginx   v1        e667596cdbd0   28 seconds ago   166MB
ubuntu/nginx     latest    f85d50b56ddd   4 months ago     139MB
ubuntu           impish    2dc51e04d744   6 months ago     77.4MB
```

3. 用自己做的nginx镜像启动一个容器
```shell
# docker run -d registry/nginx:v1

# docker ps
CONTAINER ID   IMAGE               COMMAND                  CREATED          STATUS          PORTS     NAMES
16983ef602cf   registry/nginx:v1   "/usr/sbin/nginx -g …"   15 seconds ago   Up 14 seconds   80/tcp    elastic_wescoff
```

4. 测试一下容器是不是正常工作
- 进入容器中查看nginx服务是不是启动了，配置文件是否正确
- 在容器外部使用curl测试下载文件

```shell
# docker exec 16983ef602cf ps -ef
UID          PID    PPID  C STIME TTY          TIME CMD
root           1       0  0 13:25 ?        00:00:00 nginx: master process nginx -g daemon off;
www-data       7       1  0 13:25 ?        00:00:00 nginx: worker process
www-data       8       1  0 13:25 ?        00:00:00 nginx: worker process
root          23       0  0 13:34 ?        00:00:00 ps -ef

# docker exec 16983ef602cf ls /var/www/html
file1
file2

# docker exec 16983ef602cf ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000

    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00

    inet 127.0.0.1/8 scope host lo

       valid_lft forever preferred_lft forever

168: eth0@if169: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default

    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0

    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0

       valid_lft forever preferred_lft forever

# curl -L -O http://172.17.0.2/file2
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0

# ls
file2 
```
通过这上面的这些操作，估计你已经初步感知到，容器的文件系统是独立的，运行的进程环境是独立的，网络的设置也是独立的.

让进程在一个资源可控的独立环境运行，这就是容器

5. 容器的优点

虚拟机运行状态
![虚拟机运行状态](imgs/vm_run.png)

容器运行状态
![容器运行状态](imgs/container_run.png)

![容器主要特征](imgs/docker.jpeg)


### Docker的底层技术

1. Namespace

Linux系统1号进程时systemd，进程是树状的，根进程是systemd，进程都是由父进程fork、clone出来的

```shell
# ps -ef
UID          PID    PPID  C STIME TTY          TIME CMD
root           1       0  0 May13 ?        00:01:37 /sbin/init
root           2       0  0 May13 ?        00:00:00 [kthreadd]
...
root      253893  253863  0 08:06 ?        00:00:00 nginx: master process /usr/sbin/nginx -g daemon off;
www-data  253928  253893  0 08:06 ?        00:00:00 nginx: worker process
www-data  253929  253893  0 08:06 ?        00:00:00 nginx: worker process
```

在进程被clone出来的时候，数据结构中会指定进程的Namespace

![](imgs/task_structs.jpeg)


```shell
# docker exec 16983ef602cf ps -ef
UID          PID    PPID  C STIME TTY          TIME CMD
root           1       0  0 08:06 ?        00:00:00 nginx: master process /usr/sbin/nginx -g daemon off;
www-data       7       1  0 08:06 ?        00:00:00 nginx: worker process
www-data       8       1  0 08:06 ?        00:00:00 nginx: worker process
root          21       0  0 08:09 ?        00:00:00 ps -ef


# ps -ef | grep nginx
UID        PID  PPID  C STIME TTY          TIME CMD
root      253893  253863  0 08:06 ?        00:00:00 nginx: master process /usr/sbin/nginx -g daemon off;
www-data  253928  253893  0 08:06 ?        00:00:00 nginx: worker process
www-data  253929  253893  0 08:06 ?        00:00:00 nginx: worker process
root      254370  247230  0 08:09 pts/1    00:00:00 grep --color=auto nginx

```

HOST PID Namespace  -> CONTAINER PID Namespace不一样，有对应关系
```text
253893 -> 1
253928 -> 7
253929 -> 8
```

除了PID namespace外还有：UTS USR Mount Network IPC 

![](imgs/namespace_types.jpeg)

![](imgs/namespaces.jpeg)

对Namespace的操作

### 查看当前操作系统的Namspace

```shell
# lsns -t net
        NS TYPE NPROCS    PID USER    NETNSID NSFS                           COMMAND
4026531992 net     111      1 root unassigned                                /sbin/init
4026532191 net       1    534 root unassigned                                /usr/sbin/haveged --Foreground --verbose=1 -w 1024
4026532257 net       3 253893 root          0 /run/docker/netns/fbcce4656316 nginx: master process /usr/sbin/nginx -g daemon off;
```

### 查看某进程的Namespace

```shell
# ls -la /proc/253893/ns/
total 0
dr-x--x--x 2 root root 0 May 14 08:06 .
dr-xr-xr-x 9 root root 0 May 14 08:06 ..
lrwxrwxrwx 1 root root 0 May 14 09:17 cgroup -> 'cgroup:[4026531835]'
lrwxrwxrwx 1 root root 0 May 14 08:07 ipc -> 'ipc:[4026532254]'
lrwxrwxrwx 1 root root 0 May 14 08:07 mnt -> 'mnt:[4026532252]'
lrwxrwxrwx 1 root root 0 May 14 08:06 net -> 'net:[4026532257]'
lrwxrwxrwx 1 root root 0 May 14 08:07 pid -> 'pid:[4026532255]'
lrwxrwxrwx 1 root root 0 May 14 09:27 pid_for_children -> 'pid:[4026532255]'
lrwxrwxrwx 1 root root 0 May 14 09:17 user -> 'user:[4026531837]'
lrwxrwxrwx 1 root root 0 May 14 08:07 uts -> 'uts:[4026532253]'
```

### 进入某Namespace执行命令

```shell
# nsenter -t 253893 -n ip add
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
23: eth0@if24: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```

### Namespace练习

![](imgs/namespace_exe.jpeg)

2. Cgroup

Namaspace解决进程之间的隔离，Cgroup用于解决进程资源的限制

![](imgs/cgroup_subsystems.jpeg)

![](imgs/cgroup_work.jpeg)

内核会把Cgroup子系统挂载到/sys/fs/cgroup下


### CPU subsystem
```text
cpu.shares：在满载的情况下，配置Cgroup下进程CPU使用时间的相对值
cpu.cfs_period_us：配置Cgroup下进程CPU时间周期长度，单位us
cpu.cfs_quota_us：配置Cgroup下进程CPU最大使用时间，单位us
例如：
1. CGroup A：cpu.cfs_quota_us=50000，cfs_period_us=100000，那么Cgroup A最大可以使用50000/100000=0.5个CPU，即最大CPU使用率是50%
2. Cgroup A：cpu.shares=2048，Cgroup B：cpu.shares=1024，那么在主机满载情况下，Cgroup A最大可以使用2/3的CPU，Cgroup B最大可以使用1/3的CPU
```

Cgroup CPU子系统练习：

![](imgs/cgroup_exe.jpeg)


### Memory subsystem
```text
memory.usage_in_bytes：配置Cgroup下进程使用内存的值
memory.max_usage_in_bytes：配置Cgroup下进程使用内存的最大值
memory.limit_in_bytes：配置Cgroup下进程最多使用内存的值，-1不限制
memory.oom_control：配置在Cgroup下是否使用OOM Killer，当Cgroup下进程内存使用超过最大值时，会触发OOM Kill
```

Cgroup Memory子系统练习：

![](imgs/cgroup_memory.jpeg)


### Blkio subsystem
```text
blkio.throttle.read_iops_device：Cgroup下进程磁盘读取IOPS限制
blkio.throttle.read_bps_device：Cgroup下进程磁盘读取吞吐量限制
blkio.throttle.write_iops_device：Cgroup下进程磁盘写入IOPS限制
blkio.throttle.write_bps_device：Cgroup下进程磁盘写入吞吐量限制
```
Blkio Cgroup参数配置:
```shell
echo "252:16 10485760" > $CGROUP_CONTAINER_PATH/blkio.throttle.write_bps_device
# 256是主设备号，16是次设备号
```

额外说明：

Direct I/O 和 Buffered I/O

![](imgs/page_cache.jpeg)

在Linux里，由于考虑到性能问题，绝大多数的应用都会使用 Buffered I/O 模式

3. Union FS

想象一下，假设如果我们100个镜像，每个java镜像500M，如何存储和传输这些镜像最好？

![](imgs/ideafs.jpeg)

这就是Union FS，那Union FS是如何工作的呢？

![](imgs/unionfs.jpeg)

在容器内执行df-h可以看到是通过OverlayFS实现：

![](imgs/overlayfs.jpeg)

OverlayFS练习

```shell

mkdir upper lower merged work
echo "I'm from lower!" > lower/in_lower.txt
echo "I'm from upper!" > upper/in_upper.txt
# `in_both` is in both directories
echo "I'm from lower!" > lower/in_both.txt
echo "I'm from upper!" > upper/in_both.txt

sudo mount -t overlay overlay \
 -o lowerdir=./lower,upperdir=./upper,workdir=./work \
 ./merged


```

执行完目录结构为：

```shell
# tree -h
.
├── [4.0K]  lower
│   ├── [  16]  in_both.txt
│   └── [  16]  in_lower.txt
├── [4.0K]  merged
│   ├── [  16]  in_both.txt
│   ├── [  16]  in_lower.txt
│   └── [  16]  in_upper.txt
├── [4.0K]  upper
│   ├── [  16]  in_both.txt
│   └── [  16]  in_upper.txt
└── [4.0K]  work
    └── [4.0K]  work

# cat merged/in_both.txt
I'm from upper!
```
- lower目录是底层的目录，只读的
- upper目录是上层的目录，可读写的
- merged目录是挂在后展示给用户的目录
- work目录是挂载中间文件产生的目录

模拟用户操作：

第一，新建文件
```shell
# touch merged/newadd.txt
# tree -h
.
├── [4.0K]  lower
│   ├── [  16]  in_both.txt
│   └── [  16]  in_lower.txt
├── [4.0K]  merged
│   ├── [  16]  in_both.txt
│   ├── [  16]  in_lower.txt
│   ├── [  16]  in_upper.txt
│   └── [   0]  newadd.txt
├── [4.0K]  upper
│   ├── [  16]  in_both.txt
│   ├── [  16]  in_upper.txt
│   └── [   0]  newadd.txt
└── [4.0K]  work
    └── [4.0K]  work

5 directories, 9 files
```
这个文件会出现在 upper/ 目录中

第二，删除文件：

删除 in_upper.txt 文件

```shell
# rm -rf merged/in_upper.txt
# ll upper/
total 12
drwxr-xr-x 2 root root 4096 May 21 18:32 ./
drwxr-xr-x 6 root root 4096 May 21 18:14 ../
-rw-r--r-- 1 root root   16 May 21 18:13 in_both.txt
-rw-r--r-- 1 root root    0 May 21 18:26 newadd.txt
```
这个文件会在 upper/ 目录中消失

删除 in_lower.txt 文件

```shell
# rm -rf merged/in_lower.txt
# ll ../lower/
total 16
drwxr-xr-x 2 root root 4096 May 21 18:13 ./
drwxr-xr-x 6 root root 4096 May 21 18:14 ../
-rw-r--r-- 1 root root   16 May 21 18:13 in_both.txt
-rw-r--r-- 1 root root   16 May 21 18:13 in_lower.txt
```
在 lower/ 目录里的"in_lower.txt"文件不会有变化

第三，修改文件

修改 in_upper.txt 文件

```shell
# vim merged/in_upper.txt
# cat upper/in_upper.txt
I'm from upper! And edited!
```

修改 in_lower.txt 文件

```shell
# vim merged/in_lower.txt
# cat lower/in_lower.txt
I'm from lower!
# cat upper/in_lower.txt
I'm from lower! And edited!
```
会在 upper/ 目录中新建一个"in_lower.txt"文件，包含更新的内容，而在 lower/ 中的原来的实际文件"in_lower.txt"不会改变

## Dockerfile最佳实践

1. 什么样的应用【最】适合构建成镜像

- 应用程序以一个或多个进程持续在运行
- 应用进程需是无状态的
  - 可横向扩展的，多个实例角色是等价的，配置是一样的
  - 可宰杀的，多个实例中生病的实例可以被kill
- 应用程序运行过程中产生的需要持久化数据要存储在后端服务内，例如数据库
- Session中数据要存储在后端缓存服务内，例如：Redis、Memcached

思考一下这是什么样的应用?

### ---

2. 创建Docker镜像

第一步，首先创建dockerfile

```shell
FROM ubuntu:impish
RUN apt update && apt-get install -y nginx
COPY file1 /var/www/html/
ADD  file2.tar.gz /var/www/html/
EXPOSE 80
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
```
- FROM：指定容器运行时的基础环境，如：centos:7.9.2009
- ENV：指定容器内的环境变量，如：JAVA_HOME=xxxx
- EXPOSE: 指定容器运行时暴露的端口
- COPY && ADD：向容器内添加文件，注意区别
- RUN：在应用进程启动前需要执行的命令，例如安装工具
- ENTRYPOINT & CMD：启动进程的命令
  - 需要进程在前台运行
  - ENTRYPOINT是要执行的应用程序，如：/usr/bin/nginx
  - CMD是应用程序的参数，如：-g daemon off
  - 也可以都写在一起，不推荐

第二步， 构建镜像：docker build

### Build Context

运行docker build时，当前目录被称为构建上下文

- 构建上下文中的文件会被传输给docker deamon
- 构建上下文中没用的文件会造成传输时间长、构建需要资源多

```shell
# ll /root/go/src/github.com/zheng11581/simple-cloudnative/module1/nginx/Dockerfile
# cd /usr
# ls -l
total 100
drwxr-xr-x   2 root root 36864 May 21 18:16 bin
drwxr-xr-x   2 root root  4096 Dec 29 22:20 config
drwxr-xr-x   2 root root  4096 Apr 15  2020 games
drwxr-xr-x   7 root root  4096 Jan 27 12:10 include
drwxr-xr-x  91 root root  4096 Jan 27 17:05 lib
drwxr-xr-x   2 root root  4096 Jul 31  2020 lib32
drwxr-xr-x   2 root root  4096 Dec 29 22:15 lib64
drwxr-xr-x   4 root root  4096 May 14 02:58 libexec
drwxr-xr-x   2 root root  4096 Jul 31  2020 libx32
drwxr-xr-x  10 root root  4096 Dec 29 22:07 local
drwxr-xr-x   2 root root 20480 May 14 02:58 sbin
drwxr-xr-x 122 root root  4096 May 14 02:44 share
drwxr-xr-x   6 root root  4096 Dec 29 22:20 src

# docker build -f /root/go/src/github.com/zheng11581/simple-cloudnative/module1/nginx/Dockerfile -t nginx:bigctx .
Sending build context to Docker daemon 1.258GB
^Control C

# docker build -t nginx:smallctx /root/go/src/github.com/zheng11581/simple-cloudnative/module1/nginx/
Sending build context to Docker daemon  3.584kB
Step 1/6 : FROM ubuntu:impish
 ---> 2dc51e04d744
Step 2/6 : RUN apt update && apt-get install -y nginx
 ---> Using cache
 ---> 3013adfbf428
Step 3/6 : COPY file1 /var/www/html/
 ---> Using cache
 ---> de4df850b665
Step 4/6 : ADD file2.tar.gz /var/www/html/
 ---> Using cache
 ---> bddcc5fd147e
Step 5/6 : EXPOSE 80
 ---> Using cache
 ---> 5e23d73fc62a
Step 6/6 : CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
 ---> Using cache
 ---> e667596cdbd0
Successfully built e667596cdbd0
Successfully tagged nginx:smallctx
```

### Build cache

docker build 会依次执行Dockerfile里的命令

FROM|RUN：每条指令会单独一个镜像层，简单的比较命令行字符串，字符串一致则Using cache
COPY|ADD：每条指令会单独一个镜像层，比较镜像层内文件checksum，一致则Using cache（checksum比较文件内容）
其他命令：不生成镜像层

通过Overlay FS，将【最新命令】的镜像层叠加为UpperDir，之前的命令设置为LowerDir，底层的镜像层变动会导致后面的镜像层缓存失效

```shell
# docker inspect nginx:smallctx
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/330ce617714bf507e2e00b3921604e914b28c7a530cba223e3725b682bcae273/diff:/var/lib/docker/overlay2/ad73d1fb7c90b470e967201562470b93f5f35c34efe3ff48677b0810450df958/diff:/var/lib/docker/overlay2/e25c9ddb12f1d87cd663d7f5166357c229d61111871c4f8a041dd772159a95fa/diff",
                "MergedDir": "/var/lib/docker/overlay2/ce72661d6ab0d918b1cc0983ff705cf2b54b914b78bdbc8355ad8f49e7a1dd81/merged",
                "UpperDir": "/var/lib/docker/overlay2/ce72661d6ab0d918b1cc0983ff705cf2b54b914b78bdbc8355ad8f49e7a1dd81/diff",
                "WorkDir": "/var/lib/docker/overlay2/ce72661d6ab0d918b1cc0983ff705cf2b54b914b78bdbc8355ad8f49e7a1dd81/work"
            },
```

看到LowerDir有3个镜像层，有4个FROM|RUN|ADD|COPY命令，这是因为file1和file2都是空文件，所以它们内容相同，只需要一个镜像层


### Multi-Stage Build

- docker多阶段构建：在同一个Dockerfile中先定义打包二进制文件，再将打包好的二进制文件添加到镜像中

比如Maven项目：
Stage1: copy code to working dir --> mvn package --> got target/app.jar
Stage2: copy jar file from stage1 --> entrypoint is java -jar /app.jar

比如NPM项目：
Stage1: copy code to working dir --> npm build --> got dist dir
Stage2: copy dist dir from stage1 --> entrypoint is nginx -g daemon off;

比如C#项目：
Stage1: copy code to working dir --> dotnet publish --> got app.dll
Stage2: copy dll from stage1 --> entrypoint is dotnet app.dll

```shell

# syntax = docker/dockerfile:experimental
FROM maven:3.8.5-openjdk-8-slim as maven
WORKDIR /discovery-service
COPY . .
RUN --mount=type=cache,target=/root/.m2,rw mvn -B package

FROM azul/zulu-openjdk-alpine
COPY --from=maven /discovery-service/target/discovery-service.jar discovery-service.jar
ENTRYPOINT ["java", "-jar", "/discovery-service.jar"]
EXPOSE 8761

```

- cicd多步骤构建

后面详细介绍


## 容器镜像仓库（Harbor）

1. 将CA证书ca.crt导入到本地虚拟机的Docker中，重启docker

```shell
# mkdir -p /etc/docker/certs.d/goharbor.com
# cp /vagrant_data/ca.crt /etc/docker/certs.d/goharbor.com
# systemcrl restart docker
```

2. 配置DNS解析

由于我们是本地虚拟机，而且虚拟机DNS已经指向自己电脑，只需要配置自己电脑的hosts

```shell
# vim /etc/hosts

192.168.110.72 goharbor.com

# ping goharbor.com
PING goharbor.com (192.168.110.72): 56 data bytes
64 bytes from 192.168.110.72: icmp_seq=0 ttl=62 time=10.752 ms
64 bytes from 192.168.110.72: icmp_seq=1 ttl=62 time=19.315 ms
64 bytes from 192.168.110.72: icmp_seq=2 ttl=62 time=10.964 ms
```

在虚拟机中ping goharbor.com

```shell
# ping goharbor.com
PING goharbor.com (192.168.110.72) 56(84) bytes of data.
64 bytes from goharbor.com (192.168.110.72): icmp_seq=1 ttl=63 time=12.4 ms
64 bytes from goharbor.com (192.168.110.72): icmp_seq=2 ttl=63 time=15.4 ms
```

3. Docker login 

```shell
# docker login goharbor.com
Username: admin
Password:
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded

```
[509: certificate signed by unknown authority](https://gitee.com/zheng11581/cloudnative/blob/main/kubernetes/demo/Harbor/installation/self-signed-ca.MD)

4. 推送镜像到镜像仓库

```shell
# docker tag discovery-service:latest goharbor.com/demo/discovery-service:latest

# docker push goharbor.com/demo/discovery-service:latest
```

## Dockerfile最佳实践

### 基础镜像
基础镜像按照需要选择比较小的，但不是越小越好，需要经过测试看是否满足要求

### 镜像
变动少的镜像层放到底层，也就是Dockerfile的上面，因为底层的镜像层变动会导致后面镜像层的失效

### 镜像的启动命令要选择好，可以进程可以处理signal

看一个例子，一个应用分别使用3种init进程方法：shell、java、tini启动，看看是否可以正确处理信号量

[shell as init](./entrypoint/Dockerfile-shell)

[java as init](./entrypoint/Dockerfile-shell)

[tiny as init](./entrypoint/Dockerfile-tiny)

1. init进程：shell
```shell
# docker run --name app-shell -d goharbor.com/demo/discovery-service:shell 
# docker exec -it app-shell sh
/ # ps -ef
PID   USER     TIME  COMMAND
    1 root      0:00 sh /run.sh
    7 root      0:34 java -jar /discovery-service.jar
   48 root      0:00 sh
   55 root      0:00 ps -ef

/ # kill -9 1
/ # ps -ef
PID   USER     TIME  COMMAND
    1 root      0:00 sh /run.sh
    7 root      0:35 java -jar /discovery-service.jar
   48 root      0:00 sh
   56 root      0:00 ps -ef

/ # kill 1
/ # ps -ef
PID   USER     TIME  COMMAND
    1 root      0:00 sh /run.sh
    7 root      0:35 java -jar /discovery-service.jar
   48 root      0:00 sh
   56 root      0:00 ps -ef
```
kill -9无法信号量无法被正确处理，kill信号无法被正确被处理



2. init进程：java

```shell
# docker run --name app-java -d goharbor.com/demo/discovery-service:java
# docker exec -it app-java sh
/ # ps -ef
PID   USER     TIME  COMMAND
    1 root      0:44 java -jar /discovery-service.jar
   47 root      0:00 sh
   54 root      0:00 ps -ef

/ # kill -9 1
/ # ps -ef
PID   USER     TIME  COMMAND
    1 root      0:44 java -jar /discovery-service.jar
   47 root      0:00 sh
   54 root      0:00 ps -ef

/ # kill 1
/ # %  

```
kill -9无法信号量无法被正确处理，kill信号可以正确被处理

3. init进程：tini

```shell
# docker run --name app-tini -d goharbor.com/demo/discovery-service:tini
# docker exec -it app-tini sh
/ # ps -ef
PID   USER     TIME  COMMAND
    1 root      0:00 /tini -- java -jar /discovery-service.jar
    7 root      0:23 java -jar /discovery-service.jar
   23 root      0:00 sh
   29 root      0:00 ps -ef

/ # kill -9 1
/ # ps -ef
PID   USER     TIME  COMMAND
    1 root      0:00 /tini -- java -jar /discovery-service.jar
    7 root      0:23 java -jar /discovery-service.jar
   23 root      0:00 sh
   29 root      0:00 ps -ef

/ # kill 1
/ # % 
```
kill -9无法信号量无法被正确处理，kill信号可以正确被处理

### 结论，如果可以确认业务容器的init进程可以处理SIGTERM信号（15），那么可以使用业务进程作为init进程；如果无法确认使用tini作为init进程，就像上面例子的一样

## 作业

1. 查一下docker command： 
   1. 已经推出但为被关闭的容器
   2. 进入正在运行的容器
   3. 查看容进程在操作系统的Pid

2. 将自己负责的业务应用容器化
   1. 考虑镜像大小
   2. 使用多阶段构建
   3. 将Dockerfile和业务应用代码存放在一起
   4. 业务容器init进程需要能够处理SIGTERM信号（kill）

3. 将自己负责人业务容器镜像推送到，Hatbor镜像仓库（必做）


