### 为什么需要 ReplicaSet

如果把Pod类比为部署服务端应用的虚拟机，那么我们还需要考虑服务端的高可用性，那最直接的提高服务端高可用性的方法就是冗余部署，即部署多个副本，这也就是ReplicaSet的语义：通过定义期待的Pod的副本数和Pod模板，达到多副本部署的目的

### ReplicaSet API对象的分解
- TypeMeta
  - Group：apps
  - Kind：ReplicaSet
  - Version：v1

- ObjectMeta
  - Name
  - Namespace
  - Labels
  - Annotations

- ReplicaSetSpec：各种规格属性，定义各个对象的主要区别在这里

- ReplicaSetStatus：对象的运行状态（不需要自己管理）

ReplicaSetSpec Details

```golang
type ReplicaSetSpec struct {
	// 期望创建pod的数量
	Replicas int32

	// Pod被创建好，容器保持多少秒健康状态后，才本认定为可用，防止容器启动后过一会又Crash
	MinReadySeconds int32

	// 定义受ReplicaSet控制的Pod的标签选择器，需要和Pod模板里定义的标签一致
	Selector *metav1.LabelSelector

	// Pod 模板
	Template api.PodTemplateSpec
}
```

### ReplicaSet 的实现