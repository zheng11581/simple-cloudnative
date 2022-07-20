
### 为什么需要Pod：成组调度

二进制部署: 进程 -> 虚拟机 -> CentOS/Ubuntu -> 宿主机
容器化部署：容器 -> Pod    ->  Kubernetes   -> 宿主机

```shell
# pstree -g

systemd(1)-+-accounts-daemon(1984)-+-{gdbus}(1984)
           | `-{gmain}(1984)
           |-acpid(2044)
          ...      
           |-lxcfs(1936)-+-{lxcfs}(1936)
           | `-{lxcfs}(1936)
           |-mdadm(2135)
           |-ntpd(2358)
           |-polkitd(2128)-+-{gdbus}(2128)
           | `-{gmain}(2128)
           |-rsyslogd(1632)-+-{in:imklog}(1632)
           |  |-{in:imuxsock) S 1(1632)
           | `-{rs:main Q:Reg}(1632)
           |-snapd(1942)-+-{snapd}(1942)
           |  |-{snapd}(1942)
           |  |-{snapd}(1942)
           |  |-{snapd}(1942)
           |  |-{snapd}(1942)
```
Linux 系统里的进程不是孤苦伶仃的，是以进程组的方式进行组合，例如：rsyslogd它负责的是 Linux 操作系统里的日志处理。可以看到，rsyslogd 的主程序 main，和它要用到的内核日志模块 imklog 等，同属于 1632 进程组。这些进程相互协作，共同完成 rsyslogd 程序的职责。

Kubernetes就是把这种进程组的概念映射到容器技术中形成Pod的概念（容器组）统一调度，用来解决成组调度的问题

### Pod的实现
Pod是一组共享了某些资源的容器，Pod里的所有容器，共享的是同一个 Network Namespace，并且可以声明共享同一个 Volume，它只是一个逻辑概念

![](./imgs/infra-container.png)

Infra 容器一定要占用极少的资源，所以它使用的是一个非常特殊的镜像，叫作：k8s.gcr.io/pause。这个镜像是一个用汇编语言编写的、永远处于“暂停”状态的容器，不会退出，解压后的大小也只有 100~200 KB 左右

### 为什么需要Pod：容器设计模式

[Two Containers](./container/two-containers.MD)

A B两个容器进程实际上是有“超亲密关系”的，他们需要通过文件系统进行通信，还有以下“超亲密关系”：

- 它们可以直接使用 localhost 进行通信；
- 它们看到的网络设备跟 Infra 容器看到的完全一样；
- 一个 Pod 只有一个 IP 地址，也就是这个 Pod 的 Network Namespace 对应的 IP 地址；
- 当然，其他的所有网络资源，都是一个 Pod 一份，并且被该 Pod 中的所有容器共享；
- Pod 的生命周期只跟 Infra 容器一致，而与容器 A 和 B 无关。

### Pod API对象的分解

首先先看一下所有API对象共性的东西：pkg/apis/core/types.go

- TypeMeta
  - Group
  - Kind
  - Version

- ObjectMeta
  - Name
  - [Namespace：隔离API对象、做资源隔离](../namespace/namespace-quota.MD)
  - [Labels：给对象打标签，可以做filter/selector](../selector/labels.MD)
  - Annotations
  - ...

- Spec：各种规格属性，定义各个对象的主要区别在这里

- Status：对象的运行状态（不需要自己管理）

PodSpec详解

- 容器运行时相关

```golang
// PodSpec is a description of a pod.
type PodSpec struct {
  InitContainers []Container `json:"initContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,20,rep,name=initContainers"`
  Containers []Container `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"`
  EphemeralContainers []EphemeralContainer `json:"ephemeralContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,34,rep,name=ephemeralContainers"`
  Volumes []Volume `json:"volumes,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name" protobuf:"bytes,1,rep,name=volumes"`
  ImagePullSecrets []LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"`
  RuntimeClassName *string `json:"runtimeClassName,omitempty" protobuf:"bytes,29,opt,name=runtimeClassName"`
  ...
}
```

[InitContainers](./container/2-initcontainer.MD)

[EphemeralContainers](./container/3-ephemeral-container.MD)

[Volumes](./persistentvolume/volumes.MD)

[ImagePullSecrets](./imageregistry/5-image-pull.MD)


- 生命周期（健康检查）相关

```golang
// PodSpec is a description of a pod.
type PodSpec struct {
  RestartPolicy RestartPolicy `json:"restartPolicy,omitempty" protobuf:"bytes,3,opt,name=restartPolicy,casttype=RestartPolicy"`
  ReadinessGates []PodReadinessGate `json:"readinessGates,omitempty" protobuf:"bytes,28,opt,name=readinessGates"`
  TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty" protobuf:"varint,4,opt,name=terminationGracePeriodSeconds"`
  ...
}

// A single application container that you want to run within a pod.
type Container struct {
	LivenessProbe *Probe `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`
	ReadinessProbe *Probe `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`
	StartupProbe *Probe `json:"startupProbe,omitempty" protobuf:"bytes,22,opt,name=startupProbe"`
  Lifecycle *Lifecycle `json:"lifecycle,omitempty" protobuf:"bytes,12,opt,name=lifecycle"`
}

```

[RestartPolicy](./lifecycle/6-restartpolicy.MD)

[StartupProbe]()

[LivenessProbe](./lifecycle/7-graceful-start.MD#2-liveness-probe)

[RedinessProbe](./lifecycle/7-graceful-start.MD#1-rediness-probe)

[Lifecycle-postStart](./lifecycle/7-graceful-start.MD#3-lifecycle-poststart)

[ReadinessGates](./lifecycle/7-graceful-start.MD#4-rediness-gate)

[Lifecycle-preStop](./lifecycle/graceful-stop/8-prestop.yaml)

[TerminationGracePeriodSeconds](./lifecycle/graceful-stop/8-termination-graceful.yaml)



- 调度相关

```golang
// PodSpec is a description of a pod.
type PodSpec struct {
  NodeSelector map[string]string `json:"nodeSelector,omitempty" protobuf:"bytes,7,rep,name=nodeSelector"`
  NodeName string `json:"nodeName,omitempty" protobuf:"bytes,10,opt,name=nodeName"`
  Subdomain string `json:"subdomain,omitempty" protobuf:"bytes,17,opt,name=subdomain"`
  Affinity *Affinity `json:"affinity,omitempty" protobuf:"bytes,18,opt,name=affinity"`
  SchedulerName string `json:"schedulerName,omitempty" protobuf:"bytes,19,opt,name=schedulerName"`
  Tolerations []Toleration `json:"tolerations,omitempty" protobuf:"bytes,22,opt,name=tolerations"`
  PriorityClassName string `json:"priorityClassName,omitempty" protobuf:"bytes,24,opt,name=priorityClassName"`
  Priority *int32 `json:"priority,omitempty" protobuf:"bytes,25,opt,name=priority"`
  TopologySpreadConstraints []TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty" patchStrategy:"merge" patchMergeKey:"topologyKey" protobuf:"bytes,33,opt,name=topologySpreadConstraints"`
  PreemptionPolicy *PreemptionPolicy `json:"preemptionPolicy,omitempty" protobuf:"bytes,31,opt,name=preemptionPolicy"`
  Overhead ResourceList `json:"overhead,omitempty" protobuf:"bytes,32,opt,name=overhead"`

  ...
}

// A single application container that you want to run within a pod.
type Container struct {
	Resources ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"`
}

// ResourceRequirements describes the compute resource requirements.
type ResourceRequirements struct {
	Limits ResourceList `json:"limits,omitempty" protobuf:"bytes,1,rep,name=limits,casttype=ResourceList,castkey=ResourceName"`
	Requests ResourceList `json:"requests,omitempty" protobuf:"bytes,2,rep,name=requests,casttype=ResourceList,castkey=ResourceName"`
}

```

[NodeSelector](./schedule/nodeselector/9-nodeselector.MD)

[Affinity-NodeAffinity & NodeAntiAffinity](./schedule/affinity/10-nodeaffinity.MD)

[Affinity-PodAffinity & PodAntiAffinity](./schedule/affinity/11-podaffinity.MD)

[Tolerations](./yaml/pod/../../schedule/tolerations/12-tolerations.MD)

[QOS]

- 网络与安全相关

```golang
// PodSpec is a description of a pod.
type PodSpec struct {
  DNSPolicy DNSPolicy `json:"dnsPolicy,omitempty" protobuf:"bytes,6,opt,name=dnsPolicy,casttype=DNSPolicy"`
  DeprecatedServiceAccount string `json:"serviceAccount,omitempty" protobuf:"bytes,9,opt,name=serviceAccount"`
	AutomountServiceAccountToken *bool `json:"automountServiceAccountToken,omitempty" protobuf:"varint,21,opt,name=automountServiceAccountToken"`
  HostNetwork bool `json:"hostNetwork,omitempty" protobuf:"varint,11,opt,name=hostNetwork"`
  HostPID bool `json:"hostPID,omitempty" protobuf:"varint,12,opt,name=hostPID"`
  HostIPC bool `json:"hostIPC,omitempty" protobuf:"varint,13,opt,name=hostIPC"`
  ShareProcessNamespace *bool `json:"shareProcessNamespace,omitempty" protobuf:"varint,27,opt,name=shareProcessNamespace"`
  SecurityContext *PodSecurityContext `json:"securityContext,omitempty" protobuf:"bytes,14,opt,name=securityContext"`
  Hostname string `json:"hostname,omitempty" protobuf:"bytes,16,opt,name=hostname"`
  HostAliases []HostAlias `json:"hostAliases,omitempty" patchStrategy:"merge" patchMergeKey:"ip" protobuf:"bytes,23,rep,name=hostAliases"`
  DNSConfig *PodDNSConfig `json:"dnsConfig,omitempty" protobuf:"bytes,26,opt,name=dnsConfig"`
  RuntimeClassName *string `json:"runtimeClassName,omitempty" protobuf:"bytes,29,opt,name=runtimeClassName"`
  EnableServiceLinks *bool `json:"enableServiceLinks,omitempty" protobuf:"varint,30,opt,name=enableServiceLinks"`
  SetHostnameAsFQDN *bool `json:"setHostnameAsFQDN,omitempty" protobuf:"varint,35,opt,name=setHostnameAsFQDN"`
  OS *PodOS `json:"os,omitempty" protobuf:"bytes,36,opt,name=os"`
  
  ...
}
```
[SecurityContext](./securitycontext/security-context.MD)