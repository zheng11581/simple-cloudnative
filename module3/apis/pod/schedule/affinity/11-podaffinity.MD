### How podAffinity and podAntiffinity works

1. In 11-podaffinity-target.yaml

在Deployment redis-cache中，副本上设置了标签 app=store。 podAntiAffinity 规则告诉调度器避免将多个带有 app=store 标签的副本部署到同一节点上。 因此，每个独立节点上会创建一个缓存实例

2. In 11-podaffinity.yaml

web-server Deployment 用来提供 Web 服务器服务，会创建带有标签 app=web-store 的副本。 Pod 亲和性规则告诉调度器将副本放到运行有标签包含 app=store Pod 的节点上。 Pod 反亲和性规则告诉调度器不要在同一节点上放置多个 app=web-store 的服务器。