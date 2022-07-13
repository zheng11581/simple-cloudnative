### Persistent Volumes

生命周期：Provisioning -> Binding -> Using 

根据Provisioning方式又分为：Static和Dynamic

- Static：Administrator提前分配好PersistentVolume（PV）对象，Developer通过PersistentVolumeClaim（PVC）绑定满足条件的PV，Pod声明使用PVC作为存储
- Dynamic：Administrator创建好StorageClass对象，Developer通过PersistentVolumeClaim（PVC）声明需要的PV条件并指定StorageClass，StorageClass会自动为PVC创建PV，Pod声明使用PVC作为存储

Dynamic Provisioning Demo

声明PVC并给Pod绑定PV
```shell
# kubectl apply -f 4.dynamic-provision.yaml
```

向Pod挂载的存储上写入数据
```shell
# kubectl exec -it ceph-pv-pod -- sh
# cd /mnt/ceph
# echo hello world > hello.log
# exit
```

删除、重建Pod，写入数据依然保留
```shell
# kubectl delete pod ceph-pv-pod
# kubectl apply -f 4.dynamic-provision.yaml
# kubectl exec -it ceph-pv-pod -- cat /mnt/ceph/hello.log
```

### Ephemeral Volumes

### Projected Volumes