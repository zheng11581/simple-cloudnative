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