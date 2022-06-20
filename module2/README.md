## 模块二：Kubernetes基础架构和对象

### 什么是Kubernetes?

Kubernetes是谷歌开源的容器集群管理系统，是Google多年大规模容器管理技术Borg的开源版本，主要功能包括∶
- 基于容器的应用部署、维护和滚动升级；

![](./imgs/kubelet.jpeg)

- 负载均衡和服务发现；
- 跨机器和跨地区的集群调度；
- 自动伸缩；
- 无状态服务和有状态服务；
- 插件机制保证扩展性。

### Kubernetes基础架构

Kubernetes分布式架构
![Kubernetes架构](imgs/kubernetes.jpeg)


分布式组件的功能
![Kubernetes组件](./imgs/kubernetes-detail.jpeg)

控制平面（Master Node）
- API服务器（APIServer）：这是Kubernetes控制面板中唯一带有用户可访问API以及用户可交互的组件。API服务器会暴露一个RESTful的Kubernetes API并使用JSON格式的清单文件（manifest files)
- 集群的数据存储（Cluster Data Store）：Kubernetes 使用"etcd"。这是一个强大的、稳定的、高可用的键值存储，被Kubernetes用于持久储存所有的 API对象
- 控制管理器（Controller Manager）：被称为"kube-controller manager"，它运行着所有处理集群日常任务的控制器。包括了节点控制器、副本控制器、端点（endpoint）控制器以及服务账户等
- 调度器（Scheduler）：调度器会监控新建的 pods（一组或一个容器）并将其分配给节点

数据平面（Worker Node）

- Kubelet：负责调度到对应节点的 Pod 的生命周期管理，执行任务并将 Pod 状态报告给主节点的渠道，通过容器运行时（拉取镜像、启动和停止容器等）来运行这些容器。它还会定期执行被请求的容器的健康探测程序
- Kube-proxy：它负责节点的网络，在主机上维护网络规则并执行连接转发。它还负责对正在服务的 pods进行负载平衡。

推荐的 Add-ons

- kube-dns：负责为整个集群提供 DNS服务;
- Ingress Controller：为服务提供外网入口;
- MetricsServer：提供资源监控;
- Dashboard：提供GUI;
- Fluentd-Elasticsearch：提供集群日志采集、存储与查询。

### ETCD（自行了解）

安装

```shell
ETCD_VER=v3.5.4

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GITHUB_URL}

rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

/tmp/etcd-download-test/etcd --version
/tmp/etcd-download-test/etcdctl version
/tmp/etcd-download-test/etcdutl version
```

使用场景

- 基本的key-value存储

```shell
# etcdctl member list--write-out=table
+------------------+---------+---------+-----------------------+-----------------------+------------+
|        ID        | STATUS  |  NAME   |      PEER ADDRS       |     CLIENT ADDRS      | IS LEARNER |
+------------------+---------+---------+-----------------------+-----------------------+------------+
| 8e9e05c52164694d | started | default | http://localhost:2380 | http://localhost:2379 |      false |
+------------------+---------+---------+-----------------------+-----------------------+------------+
# etcdctl put x 0
OK
# etcdctl get x 
x
0
```

- 服务注册与发现

- 基于监听机制的分布式系统

重要原理

- 基于Raft的一致性
  - http://thesecretlivesofdata.com/raft/   
  - Leader Election
  - Log Relication

- 基于Raft的安全性
  - 选举安全性：每个Term只能选举出一个Leader
  - Leader完整性：只有Term较大，Index较大的Cadidate可以当选

- 基于Raft的失效处理
  - Leader失效：恢复后会成为Follower，并被新的Leader数据覆盖
  - Follower不可用：恢复后继续作为Follower，同步Leader数据
  - 多个Candidate：随机一个Leader Election timeout（150~300ms），重新发起投票

- WAL日志

![](imgs/wal_and_mvcc.jpg)

- Watch机制



### Kubernetes的架构原则

![](./imgs/k8s_design_rules.png)

![](./imgs/k8s_layers.png)

- 核心层：Kubernetes最核心的功能。对外提供API构建高层应用；对内提供插件式应用的执行环境
- 应用层：提供部署（无状态应用、有状态应用、批处理任务、集群应用等）和路由（服务发现、DNS解析等）的能力
- 管理层：提供自动化（动态扩展）、策略管理（RBAC、Quota、PSP、NetworkPolicy）能力
- 接口层：Kubectl命令行工具、客户端SDK、集群联邦
- 生态系统：分为两个范畴
  - Kubernetes外部：日志、监控、ServiceMesh、配置管理、CICD等
  - Kubernetes内部：CRI、CNI、CSI、镜像仓库、Cloud Provider、集群自身的配置和管理

### Kubernetes的对象设计原则

- 所有API对象都是声明式的

![声明式](./imgs/declare.jpg)


- API对象是彼此互补而且可组合的


- 高层API是以操作意图为基础设计


- 底层API根据高层API的控制需要设计


- 不要有在外部API无法显式知道的内部隐藏机制

### Kubernetes核心对象