# GODEBUG

## **推荐阅读文章**
- [A whirlwind tour of Go’s runtime environment variables](https://dave.cheney.net/tag/godebug)
- [用 GODEBUG 看调度跟踪—煎鱼](https://mp.weixin.qq.com/s?__biz=MzU3Mzk5OTk1OQ==&mid=2247483675&idx=1&sn=012214890450d9f4196a160a34ee65af&chksm=fd385f23ca4fd635a5c4962b401fe583ab9abd9b333cc037ad5af35d09e3b011df323b1bd928&mpshare=1&scene=1&srcid=&sharer_sharetime=1566130293312&sharer_shareid=d4327b643a49d77dba82ac9630233b4f&key=e569ae84dd481d07aab30825f6127925297c6c3d30bc10f4623f376377ca179a9ac3af77f174c95d5f3b9c23cbe2e777c39b595fd225075bd731e145a9164b9d977c74b3e80b139f4380c7d5bbc68514&ascene=14&uin=MjEwMjA3MTA2NQ%3D%3D&devicetype=Windows+10&version=62060834&lang=zh_CN&pass_ticket=Zye%2B3E7Q9fZFmCh9kIJ4J7Oz0bXfbT%2FbLKuidX%2F5PBYhnMeE%2BqYJ1whbN6jm%2Flu3)
- [gcvis—一款可视化DEBUG程序](https://dave.cheney.net/2014/07/11/visualising-the-go-garbage-collector)
- [godebug—跨平台的debug库（是第三方库，跟前者不一样）](https://github.com/mailgun/godebug)

*以下理论部分为摘抄**煎鱼**大佬的文章片段：*

## **实践**
```
go build -o project demo.go
GODEBUG=schedtrace=1000 ./project
```
输出：
```
SCHED 0ms: gomaxprocs=4 idleprocs=2 threads=5 spinningthreads=1 idlethreads=2 runqueue=0 [9 0 0 0]
SCHED 1016ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 [2 0 4 0]
SCHED 2017ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 [2 0 4 0]
SCHED 3026ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 [2 0 4 0]
SCHED 4030ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 [2 0 4 0]
SCHED 5035ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 [2 0 4 0]
SCHED 6044ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 [2 0 4 0]
SCHED 7048ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 [2 0 4 0]
SCHED 8056ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 [2 0 4 0]
...
```
解释：
- sched：每一行都代表调度器的调试信息，后面提示的毫秒数表示启动到现在的运行时间，输出的时间间隔受 schedtrace 的值影响。
- gomaxprocs：当前的 CPU 核心数（GOMAXPROCS 的当前值）。
- idleprocs：空闲的处理器数量，后面的数字表示当前的空闲数量。
- threads：OS 线程数量，后面的数字表示当前正在运行的线程数量。
- spinningthreads：自旋状态的 OS 线程数量。
- idlethreads：空闲的线程数量。
- runqueue：全局队列中中的 Goroutine 数量，而后面的 [0 0 1 1] 则分别代表这 4 个 P 的本地队列正在运行的 Goroutine 数量。

查看详细信息：`scheddetail`：  
执行命令：`GODEBUG=schedtrace=1000,scheddetail=1 ./abc`  
输出：
```
SCHED 0ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=1 idlethreads=0 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
  P0: status=1 schedtick=1 syscalltick=0 m=0 runqsize=4 gfreecnt=0
  P1: status=1 schedtick=2 syscalltick=0 m=3 runqsize=1 gfreecnt=0
  P2: status=1 schedtick=1 syscalltick=0 m=4 runqsize=2 gfreecnt=0
  P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
  M4: p=2 curg=8 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=false lockedg=-1
  M3: p=1 curg=5 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=false lockedg=-1
  M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=true blocked=true lockedg=-1
  M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 spinning=false blocked=false lockedg=-1
  M0: p=0 curg=13 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=false lockedg=-1
  G1: status=4(semacquire) m=-1 lockedm=-1
  G2: status=4(force gc (idle)) m=-1 lockedm=-1
  G3: status=4(GC sweep wait) m=-1 lockedm=-1
  G4: status=1() m=-1 lockedm=-1
  G5: status=2() m=3 lockedm=-1
  G6: status=1() m=-1 lockedm=-1
  G7: status=1() m=-1 lockedm=-1
  G8: status=2() m=4 lockedm=-1
  G9: status=1() m=-1 lockedm=-1
  G10: status=1() m=-1 lockedm=-1
  G11: status=1() m=-1 lockedm=-1
  G12: status=1() m=-1 lockedm=-1
  G13: status=2() m=0 lockedm=-1
SCHED 1005ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
  P0: status=1 schedtick=1 syscalltick=0 m=0 runqsize=4 gfreecnt=0
  P1: status=1 schedtick=2 syscalltick=0 m=3 runqsize=1 gfreecnt=0
  ...
```
解释：
#### G
- status：G 的运行状态。
- m：隶属哪一个 M。
- lockedm：是否有锁定 M。

状态|值|含义  
---|---|---
_Gidle|0|刚刚被分配，还没有进行初始化。
_Grunnable|1|已经在运行队列中，还没有执行用户代码。
_Grunning|2|不在运行队列里中，已经可以执行用户代码，此时已经分配了 M 和 P。
_Gsyscall|3|正在执行系统调用，此时分配了 M。
_Gwaiting|4|在运行时被阻止，没有执行用户代码，也不在运行队列中，此时它正在某处阻塞等待中。
Gmoribundunused|5|尚未使用，但是在 gdb 中进行了硬编码。
_Gdead|6|尚未使用，这个状态可能是刚退出或是刚被初始化，此时它并没有执行用户代码，有可能有也有可能没有分配堆栈。
Genqueueunused|7|尚未使用。
_Gcopystack|8|正在复制堆栈，并没有执行用户代码，也不在运行队列中。

`goroutine`会被暂停运行的原因要素：

code|explain
---|---
    waitReasonZero| 
    waitReasonGCAssistMarking|GC assist marking
    waitReasonIOWait|IO wait                                  
    waitReasonChanReceiveNilChan|chan receive (nil chan)                      
    waitReasonChanSendNilChan|chan send (nil chan)                         
    waitReasonDumpingHeap|dumping heap                             
    waitReasonGarbageCollection|garbage collection                       
    waitReasonGarbageCollectionScan|garbage collection scan                   
    waitReasonPanicWait|panicwait                               
    waitReasonSelect|select                                  
    waitReasonSelectNoCases|select (no cases)                           
    waitReasonGCAssistWait|GC assist wait                            
    waitReasonGCSweepWait|GC sweep wait                             
    waitReasonChanReceive|chan receive                             
    waitReasonChanSend|chan send                                
    waitReasonFinalizerWait|finalizer wait                           
    waitReasonForceGGIdle|force gc (idle)                             
    waitReasonSemacquire|semacquire                              
    waitReasonSleep|sleep                                   
    waitReasonSyncCondWait|sync.Cond.Wait                            
    waitReasonTimerGoroutineIdle|timer goroutine (idle)                      
    waitReasonTraceReaderBlocked|trace reader (blocked)                      
    waitReasonWaitForGCCycle|wait for GC cycle                          
    waitReasonGCWorkerIdle|GC worker (idle)                            

#### M
- p：隶属哪一个 P。
- curg：当前正在使用哪个 G。
- runqsize：运行队列中的 G 数量。
- gfreecnt：可用的G（状态为 Gdead）。
- mallocing：是否正在分配内存。
- throwing：是否抛出异常。
- preemptoff：不等于空字符串的话，保持 curg 在这个 m 上运行。

#### P
- status：P 的运行状态。
- schedtick：P 的调度次数。
- syscalltick：P 的系统调用次数。
- m：隶属哪一个 M。
- runqsize：运行队列中的 G 数量。
- gfreecnt：可用的G（状态为 Gdead）。

状态|值|含义
---|---|---
_Pidle	|0|	刚刚被分配，还没有进行进行初始化。
_Prunning	|1|	当 M 与 P 绑定调用 acquirep 时，P 的状态会改变为 _Prunning。
_Psyscall	|2|	正在执行系统调用。
_Pgcstop	|3|	暂停运行，此时系统正在进行 GC，直至 GC 结束后才会转变到下一个状态阶段。
_Pdead	|4|	废弃，不再使用。

### gctrace
跟踪GC  
执行命令:
```
go build -o gc test_gc.go
GODEBUG=gctrace=1 ./gc
```
输出结果
```
gc 1 @38.952s 0%: 2.6+1132+2.4 ms clock, 10+3.7/27/0+9.7 ms cpu, 10->10->6 MB, 11 MB goal, 4 P
gc 2 @50.237s 0%: 2.1+31+2.5 ms clock, 8.7+29/25/0+10 ms cpu, 8->8->7 MB, 13 MB goal, 4 P
gc 3 @73.497s 0%: 2.3+12189+2.4 ms clock, 9.3+30/29/0+9.9 ms cpu, 10->11->8 MB, 14 MB goal, 4 P
gc 4 @140.364s 0%: 2.0+65+2.4 ms clock, 8.1+2.9/58/0+9.8 ms cpu, 14->14->9 MB, 17 MB goal, 4 P
scvg0: inuse: 14, idle: 6, sys: 20, released: 0, consumed: 20 (MB)
gc 5 @211.876s 0%: 1.9+41+2.8 ms clock, 7.8+2.7/34/0+11 ms cpu, 16->16->9 MB, 18 MB goal, 4 P
scvg1: inuse: 18, idle: 50, sys: 69, released: 0, consumed: 69 (MB)
gc 6 @304.116s 0%: 1.8+70+2.4 ms clock, 7.2+3.1/62/0+9.8 ms cpu, 18->18->11 MB, 19 MB goal, 4 P
gc 7 @427.008s 0%: 1.6+7816+2.9 ms clock, 6.4+2.9/58/0+11 ms cpu, 21->22->12 MB, 22 MB goal, 4 P
GC forced
scvg2: inuse: 22, idle: 67, sys: 90, released: 0, consumed: 90 (MB)
...
```
解释：  
输出有两种数据：
1. gc: 追踪GC信息
    - `2.6+1132+2.4 ms clock`: gc各个阶段的时间(`STW`, `标记`，`清除` 这三个阶段？)
    - `10->10->6 MB`: 垃圾回收堆大小的变化
    - `(forced)`: 表示强制进行垃圾回收（代码中调用`runtime.GC()`，该代码中没有调用，因此没有该标记）
    - `GC forced`: 在第7次GC后，进行了强制GC，应该是达到了某个条件（TODO）
2. scvg：定期扫描堆，查看使用情况
    - `scvg0: inuse: 14, idle: 6, sys: 20, released: 0, consumed: 20 (MB)`
    > sys: 从系统中请求的内存总量  
      inuse: 整个堆使用的内存总量，可能包括已死亡的对象  
      idle: 表示GC当前未使用的内存，包含大量已死亡的对象，该对象在GC后处于未使用状态。  
      released: 系统回收的内存。如果idle长期保持，那么就会提醒系统进行回收一部分内存。  
      consumed: 消耗的系统内存。
      
     > 总的来讲： sys = inuse+idle = released + consumed
```
scvg47: inuse: 32, idle: 108, sys: 140, released: 61, consumed: 79 (MB)
GC forced
gc 60 @7317.072s 0%: 1.9+100+3.4 ms clock, 7.9+3.0/78/0+13 ms cpu, 31->31->21 MB, 43 MB goal, 4 P
scvg48: inuse: 25, idle: 110, sys: 136, released: 61, consumed: 74 (MB)
gc 61 @7442.477s 0%: 1.5+124+2.4 ms clock, 6.2+2.6/105/0+9.6 ms cpu, 31->31->21 MB, 42 MB goal, 4 P
scvg49: 8 MB released
scvg49: inuse: 28, idle: 98, sys: 127, released: 70, consumed: 56 (MB)
GC forced
gc 62 @7563.021s 0%: 1.6+87+2.4 ms clock, 6.5+3.0/72/0+9.9 ms cpu, 30->30->20 MB, 42 MB goal, 4 P
scvg50: 9 MB released
scvg50: inuse: 28, idle: 106, sys: 134, released: 80, consumed: 54 (MB)
GC forced
gc 63 @7685.633s 0%: 2.4+111+2.5 ms clock, 9.8+2.7/104/0+10 ms cpu, 31->31->21 MB, 41 MB goal, 4 P
scvg51: inuse: 31, idle: 76, sys: 107, released: 76, consumed: 31 (MB)
```
上述是某一段时间的执行结果，尝试解读一下：
> `scvg49` 释放了8M内存，而它之前刚刚结束了的`gc61`释放了10M内存。下边也有类似的数据，那么是否是说，GC释放的内存，一部分还给了系统，另一部分用在了其他地方，可能是`inuse`,可能是`idle`  
`scvg51`中的`released`要比`scvg50`中的`released`小，说明该字段并不是累加的，它是一个状态值?  
好吧，这一块还不是很清楚，以后了解更深入后再来完善。
      
*开启gctrace对生产环境的影响几乎为0，因为他们一直在被统计，只是没有被表现出来。*

推荐一款可视化工具 [gcvis](https://github.com/davecheney/gcvis)

## **注意事项**
- `windows`下的`GODEBUG`无效，`linux`下可以。
- 参数之间用逗号来分割，且不能有空格，否则后边的参数无效