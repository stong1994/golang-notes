## 记录工作中用到的一些封装方法

- rand_num: 生成不同的正整数随机值
- parallel_access: 同一个数据有多个不稳定的数据源时，并行请求，获取到一个数据后取消掉其它goroutine，节省资源
- sort: 封装sort.Sort()方法，避免在需要对多个属性排序时，写多个实现。
- visit_rate_limit: 对ip访问进行限制
- float_operate: 工厂模式实现浮点运算
- join_string: 比较集中字符串拼接的效率
- command_with_api: 用api来访问linux服务器并进行操作