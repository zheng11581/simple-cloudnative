# simple-cloudnative
- 预习：安装Docker
- 模块一：Docker核心技术
- 模块二：Kubernetes架构原理
- 模块三：Kubernetes容器编排
- 模块四：基于Kubernetes的CICD

## 预习：课前准备
1. 安装Ubuntu 20.04虚拟机
- 安装VirtualBox
- 安装vagrant

```shell
# VirtualBox下载安装，非常方便
# https://www.virtualbox.org/wiki/Downloads
# Vagrant下载安装，非常方便
# https://www.vagrantup.com/downloads
```

- 启动使用Vagrantfile文件启动虚拟机

```shell
# cd module0/Vagrantfile
# mkdir data
# vagrant up
# vagrant ssh
```


2. 安装Docker

```shell
# 在 ubuntu上安装 Docker运行时，参考：https∶//docs.docker.com/engine/install/ubuntu/
# sudo apt-get update 
# sudo apt-get install \
apt-transport-https \
ca-certificates \
curl \
gnupg-agent \
software-properties-common
# curl -fsSL https://download.docker.com/linux/ubuntu/gpg|sudo apt-key add -
# sudo add-apt-repository \
"deb [arch=amd64] https://download.docker.com/linux/ubuntu \
$(lsb_release -cs) \
stable"
# sudo apt-get update
# sudo apt-get install docker-ce docker-ce-cli containerd.io

```

## 模块一：Docker核心技术
### 目标
1. 理解什么是Docker容器
2. Docker的底层技术
3. 将我们的业务容器化，过程需要注意什么

### 什么是Docker容器？
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
REPOSITORY  TAG   IMAGEID  CREATED   SIZE
registry/nginx  v1  c682fc3d4b9a  4 seconds ago  277MB
```

3. 用自己做的nginx镜像启动一个容器
```shell
# docker run -d registry/nginx:v1

# docker ps
CONTAINER ID      IMAGE        COMMAND     CREATED       STATUS        PORTS               NAMES
881b0539eae8      registry/nginx:v1   "/usr/sbin/nginx -g daemon off;"   2 seconds ago       Up 2 seconds                            loving_jackson
```

4. 测试一下容器是不是正常工作
- 进入容器中查看nginx服务是不是启动了，配置文件是否正确
- 在容器外部使用curl测试下载文件

```shell
# docker exec 881b0539eae8 ps -ef
UID          PID    PPID  C STIME TTY          TIME CMD
root           1       0  0 13:25 ?        00:00:00 nginx: master process nginx -g daemon off;
www-data       7       1  0 13:25 ?        00:00:00 nginx: worker process
www-data       8       1  0 13:25 ?        00:00:00 nginx: worker process
root          23       0  0 13:34 ?        00:00:00 ps -ef

# docker exec 881b0539eae8 ls /var/www/html
file1
file2

# docker exec 881b0539eae8 ip addr
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

100     6  100     6    0     0   1500      0 --:--:-- --:--:-- --:--:--  1500

# ls
file2 
```
通过这上面的这些操作，估计你已经初步感知到，容器的文件系统是独立的，运行的进程环境是独立的，网络的设置也是独立的.

让进程在一个资源可控的独立环境运行，这就是容器

### Docker的底层技术

虚拟机运行状态
![虚拟机运行状态](module1/imgs/vm_run.png)

容器运行状态
![容器运行状态](module1/imgs/container_run.png)


1. Namespace

```shell
# docker exec 881b0539eae8 ps -ef
UID          PID    PPID  C STIME TTY          TIME CMD
root           1       0  0 13:25 ?        00:00:00 nginx: master process nginx -g daemon off;
www-data       7       1  0 13:25 ?        00:00:00 nginx: worker process
www-data       8       1  0 13:25 ?        00:00:00 nginx: worker process
root          23       0  0 13:34 ?        00:00:00 ps -ef


# ps -ef | grep nginx

UID        PID  PPID  C STIME TTY          TIME CMD
root     20731 20684  0 18:59 ?        00:00:01 nginx: master process nginx -g daemon off;
48       20787 20731  0 18:59 ?        00:00:00 nginx: worker process
48       20788 20731  0 18:59 ?        00:00:06 nginx: worker process

```

HOST  -> CONTAINER PID不一样，有对应关系
```text
20731 -> 1
20787 -> 7
20788 -> 8
```

除了PID namespace外还有：UTS USR Mount Network IPC 

2. Cgroup



3. Union FS

### Dockerfile最佳实践

## 模块二：Kubernetes架构原理

## 模块三：Kubernetes容器编排

## 基于Kubernetes的CICD