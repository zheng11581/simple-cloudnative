## 模块二：Kubernetes架构原则和对象设计原则

### 什么是Kubernetes?

Kubernetes是谷歌开源的容器集群管理系统，是Google多年大规模容器管理技术Borg的开源版本，主要功能包括∶
- 基于容器的应用部署、维护和滚动升级；

![](./imgs/kubelet.jpeg)

- 负载均衡和服务发现；
- 跨机器和跨地区的集群调度；
- 自动伸缩；
- 无状态服务和有状态服务；
- 插件机制保证扩展性。

### Kubernetes基本组件

Kubernetes分布式架构

![Kubernetes架构](imgs/kubernetes.jpeg)

- API服务器（APIServer）：这是Kubernetes控制面板中唯一带有用户可访问API以及用户可交互的组件。API服务器会暴露一个RESTful的Kubernetes API并使用JSON格式的清单文件（manifest files)
- 集群的数据存储（Cluster Data Store）：Kubernetes 使用"etcd"。这是一个强大的、稳定的、高可用的键值存储，被Kubernetes用于持久储存所有的 API对象
- 控制管理器（Controller Manager）：被称为"kube-controller manager"，它运行着所有处理集群日常任务的控制器。包括了节点控制器、副本控制器、端点（endpoint）控制器以及服务账户等
- 调度器（Scheduler）：调度器会监控新建的 pods（一组或一个容器）并将其分配给节点

分布式组件的功能
![Kubernetes组件](./imgs/kubernetes-detail.jpeg)



### Kubernetes的架构原则

### Kubernetes的对象设计原则

### Kubernetes核心对象