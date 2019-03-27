*这篇只是对其它文章的个人总结（摘抄），强烈推荐阅读原文，原文链接在下边。*

### 一、Goroutine Schedule
一个GO程序对操作系统来讲只是一个**用户层程序**，在它的眼中只有thread。操作系统不会知道goroutine是什么，而goroutine 的调度全靠go自己完成。

在操作系统层面，Thread竞争的是真实的物理CPU资源，但在Go层面，goroutine竞争的资源是操作系统线程。于是，将goroutine按照一定算法放到线程上执行就是go schedule的任务。这种语言层面自带调度器的，称之为**原生支持并发**

### 二、演化过程
#### 1.G-M模型
2012年3月28日，Go1.0正式发布。在这个版本的调度器中，groutine对应于runtime中的一个抽象结构：G，而os thread作为”物理CPU“的存在而被抽象为一个结构：M。

前Intel blackbelt工程师、现Google工程师Dmitry Vyukov在其《Scalable Go Scheduler Design》一文中指出了G-M模型的一个重要不足： 限制了Go并发程序的**伸缩性**，尤其是对那些有高吞吐或并行计算需求的服务程序。主要体现在如下几个方面：
> - 单一全局互斥锁(Sched.Lock)和集中状态存储的存在导致所有goroutine相关操作，比如：创建、重新调度等都要上锁；
> - goroutine传递问题：M经常在M之间传递”可运行”的goroutine，这导致调度延迟增大以及额外的性能损耗；
> - 每个M做内存缓存，导致内存占用过高，数据局部性较差；
> - 由于syscall调用而形成的剧烈的worker thread阻塞和解除阻塞，导致额外的性能损耗。

#### 2.G-M-P模型
Dmitry Vyukov亲自操刀改进Go scheduler，在Go 1.1中实现了G-P-M调度模型和work stealing算法，这个模型一直沿用至今：
![image](http://tonybai.com/wp-content/uploads/goroutine-scheduler-model.png)

## 核心概念
### G (goroutine)
- goroutine可以解释为受管理的轻量线程
- main本身就是一个goroutine
- goroutine执行异步操作时会进入休眠状态, 待操作完成后再恢复, 无需占用系统线程,
- goroutine新建或恢复时会添加到运行队列, 等待M取出并运行.

### M (machine)
- machine等同于系统线程
- M运行两种代码
> - go代码, 即goroutine, M运行go代码需要一个P
> - 原生代码, 例如阻塞的syscall, M运行原生代码不需要P
- M会从运行队列中取出G, 然后运行G, 如果G运行完毕或者进入休眠状态, 则从运行队列中取出下一个G运行, 周而复始
- 有时候G需要调用一些无法避免阻塞的原生代码, 这时M会释放持有的P并进入阻塞状态, 其他M会取得这个P并继续运行队列中的G.
- go需要保证有足够的M可以运行G, 不让CPU闲着, 也需要保证M的数量不能过多.

### P (process)
- P是process的头文字, 代表M运行G所需要的资源.
- P的数量默认等于cpu的数量，但可以通过环境变量`GOMAXPROC`修改
- P也可以理解为控制go代码的并行度的机制
- 执行原生代码的线程数量不受P控制
- 因为同一时间只有一个线程(M)可以拥有P, P中的数据都是锁自由(lock free)的, 读写这些数据的效率会非常的高.

## 数据结构 
### G的状态
- 空闲中(_Gidle): 表示G刚刚新建, 仍未初始化
- 待运行(_Grunnable): 表示G在运行队列中, 等待M取出并运行
- 运行中(_Grunning): 表示M正在运行这个G, 这时候M会拥有一个P
- 系统调用中(_Gsyscall): 表示M正在运行这个G发起的系统调用, 这时候M并不拥有P
- 等待中(_Gwaiting): 表示G在等待某些条件完成, 这时候G不在运行也不在运行队列中(可能在channel的等待队列中)
- 已中止(_Gdead): 表示G未被使用, 可能已执行完毕(并在freelist中等待下次复用)
- 栈复制中(_Gcopystack): 表示G正在获取一个新的栈空间并把原来的内容复制过去(用于防止GC扫描)

### M的状态
- 自旋中(spinning): M正在从运行队列获取G, 这时候M会拥有一个P
- 执行go代码中: M正在执行go代码, 这时候M会拥有一个P
- 执行原生代码中: M正在执行原生代码或者阻塞的syscall, 这时M并不拥有P
- 休眠中: M发现无待运行的G时会进入休眠, 并添加到空闲M链表中, 这时M并不拥有P

### P的状态
- 空闲中(_Pidle): 当M发现无待运行的G时会进入休眠, - 这时M拥有的P会变为空闲并加到空闲P链表中
- 运行中(_Prunning): 当M拥有了一个P后, 这个P的状态就会变为运行中, M运行G会使用这个P中的资源
- 系统调用中(_Psyscall): 当go调用原生代码, 原生代码又反过来调用go代码时, 使用的P会变为此状态
- GC停止中(_Pgcstop): 当gc停止了整个世界(STW)时, P会变为此状态
- 已中止(_Pdead): 当P的数量在运行时改变, 且数量减少时多余的P会变为此状态

### 本地运行队列
- 在go中有多个运行队列可以保存待运行(_Grunnable)的G, 它们分别是各个P中的本地运行队列和全局运行队列.
- 入队待运行的G时会优先加到当前P的本地运行队列, M获取待运行的G时也会优先从拥有的P的本地运行队列获取
- 本地运行队列入队和出队不需要使用线程锁.
- 本地运行队列有数量限制, 当数量达到256个时会入队到全局运行队列.
- 本地运行队列的数据结构是[环形队列](https://en.wikipedia.org/wiki/Circular_buffer), 由一个256长度的数组和两个序号(head, tail)组成.
- 当M从P的本地运行队列获取G时, 如果发现本地队列为空会尝试从其他P盗取一半的G过来,这个机制叫做[Work Stealing](http://supertech.csail.mit.edu/papers/steal.pdf)。[原文](http://www.cnblogs.com/zkweb/p/7815600.html)中有代码分析

### 全局运行队列
- 全局运行队列保存在全局变量`sched`中, 全局运行队列入队和出队需要使用线程锁.
- 全局运行队列的数据结构是链表, 由两个指针(head, tail)组成.

### 空闲M链表
- 当M发现无待运行的G时会进入休眠, 并添加到空闲M链表中, 空闲M链表保存在全局变量`sched`.
- 进入休眠的M会等待一个信号量(`m.park`), 唤醒休眠的M会使用这个信号量.
- go需要保证有足够的M可以运行G, 是通过这样的机制实现的:
> - 入队待运行的G后, 如果当前无自旋的M但是有空闲的P, 就唤醒或者新建一个M
> - 当M离开自旋状态并准备运行出队的G时, 如果当前无自旋的M但是有空闲的P, 就唤醒或者新建一个M
> - 当M离开自旋状态并准备休眠时, 会在离开自旋状态后再次检查所有运行队列, 如果有待运行的G则重新进入自旋状态
- 因为"入队待运行的G"和"M离开自旋状态"会同时进行, go会使用这样的检查顺序:
> - 入队待运行的G => 内存屏障 => 检查当前自旋的M数量 => 唤醒或者新建一个M
> - 减少当前自旋的M数量 => 内存屏障 => 检查所有运行队列是否有待运行的G => 休眠
>> 这样可以保证不会出现待运行的G入队了, 也有空闲的资源P, 但无M去执行的情况.

### 空闲P链表
当P的本地运行队列中的所有G都运行完毕, 又不能从其他地方拿到G时,
拥有P的M会释放P并进入休眠状态, 释放的P会变为空闲状态并加到空闲P链表中, 空闲P链表保存在全局变量`sched`
下次待运行的G入队时如果发现有空闲的P, 但是又没有自旋中的M时会唤醒或者新建一个M, M会拥有这个P, P会重新变为运行中的状态.

## 实操
### G被抢占调度
- 除非极端的无限循环或死循环，否则只要G调用函数，Go runtime就有抢占G的机会
- Go程序启动时，runtime会去启动一个名为`sysmon`的m(一般称为监控线程)，该m无需绑定p即可运行，该m在整个Go程序的运行过程中至关重要
> - sysmon每20us~10ms启动一次
>> - 释放闲置超过5分钟的span物理内存；
>> - 如果超过2分钟没有垃圾回收，强制执行；
>> - 将长时间未处理的netpoll结果添加到任务队列；
>> - 向长时间运行的G任务发出抢占调度；
>> - 收回因syscall长时间阻塞的P；
> - 如果一个G任务运行10ms，sysmon就会认为其运行时间太久而发出抢占式调度的请求。一旦G的抢占标志位被设为true，那么待这个G下一次调用函数或方法时，runtime便可以将G抢占，并移出运行状态，放入P的local runq中，等待下一次被调度。

### channel阻塞或network I/O情况下的调度
如果G被阻塞在某个channel操作或network I/O操作上时，G会被放置到某个wait队列中，而M会尝试运行下一个runnable的G；如果此时没有runnable的G供m运行，那么m将解绑P，并进入sleep状态。当I/O available或channel操作完成，在wait队列中的G会被唤醒，标记为runnable，放入到某P的队列中，绑定一个M继续执行。
### system call阻塞情况下的调度
如果G被阻塞在某个system call操作上，那么不光G会阻塞，执行该G的M也会解绑P(实质是被sysmon抢走了)，与G一起进入sleep状态。如果此时有idle的M，则P与其绑定继续执行其他G；如果没有idle M，但仍然有其他G要去执行，那么就会创建一个新M。  

当阻塞在syscall上的G完成syscall调用后，G会去尝试获取一个可用的P，如果没有可用的P，那么G会被标记为runnable，之前的那个sleep的M将再次进入sleep。

### 参考文章
- https://tonybai.com/2017/06/23/an-intro-about-goroutine-scheduler/
- http://www.cnblogs.com/zkweb/p/7815600.html
- [微信公众号文章](https://mp.weixin.qq.com/s?__biz=Mzg3MTA0NDQ1OQ==&mid=2247483888&idx=1&sn=665fe9a84c2fb5e40624a1038f7b4fba&chksm=ce85c5f4f9f24ce289fcdf3b7dc308b84adc23d51c91e19fd77f4914649769558d3efed34358&mpshare=1&scene=2&srcid=&from=timeline&key=87a1c8d6f03747dccf54978634dfcede5c375343cc9bd743557e2dd8382bad6195543d606f9172f666274e3772918da568c8265e36e736cc1bec7b2c4d2a15315474cab8dafe59b6ce96a6a2a1db4245&ascene=14&uin=MjEwMjA3MTA2NQ%3D%3D&devicetype=Windows+10&version=62060739&lang=zh_CN&pass_ticket=Qi5GZ1CK7GE6Z044V9J3UuOu0is1A8QN4yh4%2B6SfTjxBG4yyOQTxzm%2FG4hToQO%2Fo)
