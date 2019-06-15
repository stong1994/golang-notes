### 两种方式实现ip的限制访问

1. 用sync.RWMutex来控制全局的ip访问数据（map）

2. 通过channel来控制ip访问数据
> 这两种方式的性能需测试，总的来说， sync.RWMutex 适用于多读少写的操作，这里显然不适合。