### Persistent Volumes

生命周期：Provisioning -> Binding -> Using 

根据Provisioning方式又分为：Static和Dynamic

- Static：Administrator提前分配好PersistentVolume（PV）对象，Developer通过PersistentVolumeClaim（PVC）绑定满足条件的PV，Pod声明使用PVC作为存储
- Dynamic：Administrator创建好StorageClass对象，Developer通过PersistentVolumeClaim（PVC）声明需要的PV条件并指定StorageClass，StorageClass会自动为PVC创建PV，Pod声明使用PVC作为存储

### Ephemeral Volumes

### Projected Volumes