# Sharing Memory By Communicating

阅读文章，并实现代码，文章[地址链接](https://mp.weixin.qq.com/s?__biz=Mzg3MTA0NDQ1OQ==&mid=2247483913&idx=1&sn=6dfb3bdea18be318da1de7df4251ed6a&chksm=ce85c60df9f24f1b7dcb493a7459b60ab1f4a028d042eb3309be3072e6fdac1d33c8c3f103c7&mpshare=1&scene=1&srcid=#rd)

**思考问题**
1. 通过channel来传递数据而不用mutex，那么就不能同时读，效率上会不会有较大的影响
2. 通过channel来传递数据，代码更多，理解不如传统的加锁这样直观。需要加深印象并深入理解。