### 两种方式实现ip的限制访问

1. 用sync.RWMutex来控制全局的ip访问数据（map）

2. 通过channel来控制ip访问数据
> 性能测试结果，选用方式1的效率较高。
