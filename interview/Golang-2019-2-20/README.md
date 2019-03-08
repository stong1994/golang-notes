1. 项目初始化时，不同包内的init函数的执行顺序
2. 切片和数组，作为参数传入方法后，改变某个元素，原切片或数组是否改变
3. make和new的区别
* new(T) 返回的是 T 的指针
* make(T, args) 返回的是 T 的 引用
* make 只能用于 slice,map,channel (注意: make不能初始化数组)
* 对于第三点的扩展：这三种类型作为参数传入函数后，在函数内部修改将影响函数外部的值，而其他类型必须传递该类型的指针才能有这种效果
* maps/slices/channels 在底层隐含了指针, 所以使用中并没有需要使用指针的语法. 但是引用内存是基于指针实现的, 本质上是有一个构造的过程的.

4. 按照规定设计接口，并实现该接口（接收者用指针类型和值类型的区别：接收者可以理解为函数的第一个参数，传递值类型，不会改变函数外的值，而指针类型相反）
5. rand & 逃逸？
> 首先直接调用rand.Intn()，每次返回值可能是一样的，因为在默认情况下，种子seed是一样的。
> 发生了逃逸,参考：https://mp.weixin.qq.com/s?__biz=MjM5OTcxMzE0MQ==&mid=2653372198&idx=1&sn=7f52768337e4ee77e7276b39afcdc516&chksm=bce4df3c8b93562ad0f288ef54bbc15fb75bb2ac6b6ba38a42d5230b15d4ad0f582aa9ba1872&mpshare=1&scene=1&srcid=&key=e3738c51d3aaafb50afb3eaad4c8ea733b7f44b398bf0453c5497cc0f34b412a0124c68b917f359ff808f1183a9d3fd745870d78d8d87715c70cf4fcb98ae3327f32af5c51ff3dd0af6d43e4f7c4c76c&ascene=1&uin=MjEwMjA3MTA2NQ%3D%3D&devicetype=Windows+10&version=62060728&lang=zh_CN&pass_ticket=b7LZnV49n4%2FGYAMBmnhW1xjJfbSpfbVAwbNv06eB9T8Oi%2F37FB6AsR3giUor%2B5RD
6. 设计两个goroutine，一个打印a,c,e...，另一个打印b,d,f...输出结构为a,b,c,d,e...z
7. 性能监控(可以从pprof或runtime来分析)
8. 用context来取消子goroutine
9. 信号的使用
> 说实话，没用过，另外，信号对于通道的有点待研究。
10. GC
> 推荐链接：http://legendtkl.com/2017/04/28/golang-gc/