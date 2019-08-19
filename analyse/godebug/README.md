# GODEBUG

## **推荐阅读文章**
- [A whirlwind tour of Go’s runtime environment variables](https://dave.cheney.net/tag/godebug)
- [用 GODEBUG 看调度跟踪—煎鱼](https://mp.weixin.qq.com/s?__biz=MzU3Mzk5OTk1OQ==&mid=2247483675&idx=1&sn=012214890450d9f4196a160a34ee65af&chksm=fd385f23ca4fd635a5c4962b401fe583ab9abd9b333cc037ad5af35d09e3b011df323b1bd928&mpshare=1&scene=1&srcid=&sharer_sharetime=1566130293312&sharer_shareid=d4327b643a49d77dba82ac9630233b4f&key=e569ae84dd481d07aab30825f6127925297c6c3d30bc10f4623f376377ca179a9ac3af77f174c95d5f3b9c23cbe2e777c39b595fd225075bd731e145a9164b9d977c74b3e80b139f4380c7d5bbc68514&ascene=14&uin=MjEwMjA3MTA2NQ%3D%3D&devicetype=Windows+10&version=62060834&lang=zh_CN&pass_ticket=Zye%2B3E7Q9fZFmCh9kIJ4J7Oz0bXfbT%2FbLKuidX%2F5PBYhnMeE%2BqYJ1whbN6jm%2Flu3)
- [godebug—跨平台的debug库（是第三方库，跟前两者不一样）](https://github.com/mailgun/godebug)

*以下理论部分为摘抄**煎鱼**大佬的文章片段：*

## **实践**
```
go build -o project
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

## **注意事项**
- `windows`下的`GODEBUG`无效，`linux`下可以。
- 参数之间用逗号来分割，且不能有空格，否则后边的参数无效