# channel源码阅读

*go version go1.12.5 linux/amd64*

**代码位置**: runtime/chan.go  
 一共约700行
 
 四种操作:创建,发送,接收和关闭
 
 关闭: sendq和receq为nil,buf可能不为空,如果为空则清零reader的读取位置,如果不为空则继续读buf
 
 
 #### 结构体
 ```go
 type hchan struct {
	qcount   uint           // 当前队列中数据的数量
	dataqsiz uint           // 环形队列的大小
	buf      unsafe.Pointer // 指向环形队列的数组的指针
	elemsize uint16 // 元素大小
	closed   uint32 // 是否关闭
	elemtype *_type // 元素类型
	sendx    uint   // 发送的元素的索引
	recvx    uint   // 接收的元素的索引
	recvq    waitq  // 等待接收元素的列表
	sendq    waitq  // 等待发送元素的列表
	lock mutex // 锁，保护hchan中的元素及子元素
}

// 由于不可能同时存在发送和接收的缓存，因此只用一个buf来存储缓存数据即可

type waitq struct {
	first *sudog // 等待列表中的第一个
	last  *sudog // 等待列表中的最后一个
}

// sudog 和 g 之间的关系是一对多,即一个g可能有多个sudog.比如一个g可能在多个等待列表中,那么每个都会产生一个sudog
type sudog struct {
	g *g // GMP中的G,即goroutine

	isSelect bool // 表示g参与了select
	next     *sudog
	prev     *sudog
	elem     unsafe.Pointer

	acquiretime int64
	releasetime int64
	ticket      uint32
	parent      *sudog // semaRoot binary tree
	waitlink    *sudog // g.waiting list or semaRoot
	waittail    *sudog // semaRoot
	
	c           *hchan // channel
}
```

#### 通用的send
```go
func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {
    // 向一个nil的channel发送数据，如果是非阻塞的，直接返回false；如果是阻塞的，调用gopark
    // gopark会将当前的goroutine休眠，并通过调用第一个参数来唤醒
    // 此处第一个参数为nil，因此不会被唤醒，接收和发送的goroutine都会休眠，造成死锁
	if c == nil {
		if !block {
			return false
		}
		gopark(nil, nil, waitReasonChanSendNilChan, traceEvGoStop, 2)
		throw("unreachable")
	}
	/* ...
	   ...
	 */
	// Fast path: 不获取锁就能检查失败的非阻塞操作
	if !block && c.closed == 0 && ((c.dataqsiz == 0 && c.recvq.first == nil) ||
		(c.dataqsiz > 0 && c.qcount == c.dataqsiz)) {
		return false
	}

	var t0 int64
	if blockprofilerate > 0 {
		t0 = cputicks()
	}

	lock(&c.lock)

	// 向一个已经关闭的channel发送数据会panic
	if c.closed != 0 {
		unlock(&c.lock)
		panic(plainError("send on closed channel"))
	}
	
    // 如果当前的等待接收队列中存在数据，那么不用通过buf，直接把值发送过去。
	if sg := c.recvq.dequeue(); sg != nil {
		send(c, sg, ep, func() { unlock(&c.lock) }, 3)
		return true
	}
    
	// 如果等待队列为空,那么就需要等待其他goroutine来接收,查看ring buffer是否已满
	// 如果等待发送的队列中的数据数量 小于 环形队列（ring buffer）的大小，直接将数据放到环形队列中
	if c.qcount < c.dataqsiz {
		qp := chanbuf(c, c.sendx)
		if raceenabled {
			raceacquire(qp)
			racerelease(qp)
		}
		typedmemmove(c.elemtype, qp, ep)
		c.sendx++ // 标记发送索引
		if c.sendx == c.dataqsiz {
			c.sendx = 0
		}
		c.qcount++
		unlock(&c.lock)
		return true
	}

	// 如果ring buffer已满,并且为非阻塞,则直接返回false 
	if !block {
		unlock(&c.lock)
		return false
	}

	
	// ring buffer已满,并且为非阻塞，那么阻塞当前g，直到数据被接收
	gp := getg()
	mysg := acquireSudog()
	mysg.releasetime = 0
	if t0 != 0 {
		mysg.releasetime = -1
	}
	// No stack splits between assigning elem and enqueuing mysg
	// on gp.waiting where copystack can find it.
	mysg.elem = ep
	mysg.waitlink = nil
	mysg.g = gp
	mysg.isSelect = false
	mysg.c = c
	gp.waiting = mysg
	gp.param = nil
	c.sendq.enqueue(mysg)
	// 将当前的g从队列中移出
	goparkunlock(&c.lock, waitReasonChanSend, traceEvGoBlockSend, 3)
	// 因为调度器在停止当前 g 的时候会记录运行现场，当恢复阻塞的发送操作时候，会从此处继续开始执行
	
	// 保持数据存活
	KeepAlive(ep)

	// 被唤醒
	if mysg != gp.waiting {
		throw("G waiting list is corrupted")
	}
	gp.waiting = nil
	if gp.param == nil {
		if c.closed == 0 {
			throw("chansend: spurious wakeup")
		}
		panic(plainError("send on closed channel"))
	}
	gp.param = nil
	if mysg.releasetime > 0 {
		blockevent(mysg.releasetime-t0, 2)
	}
	mysg.c = nil 
	releaseSudog(mysg) // 释放掉sudog
	return true
}
```

```go
// send 代表了对一个空的channel的发送操作（在channl的等待接收队列不为空时，channel接收到了值，那么直接调用send进行值的拷贝）
// ep被发送方拷贝给接受者的sg
// 接受者将被唤醒
// c 必须是空的，并且处于锁住状态.  调用unlockf来解锁.
// sg 是从 c 中出列 的sudog
// ep 必须是非空的，并且执行堆或者调用者的栈
func send(c *hchan, sg *sudog, ep unsafe.Pointer, unlockf func(), skip int) {
	/**
	...
	 */
	if sg.elem != nil {
		sendDirect(c.elemtype, sg, ep)
		sg.elem = nil
	}
	gp := sg.g // 获取到sudog所在的groutine
	unlockf() 
	gp.param = unsafe.Pointer(sg)
	if sg.releasetime != 0 {
		sg.releasetime = cputicks()
	}
	// 唤醒groutine（由于其一直在等待数据被接收，因此一直处在睡眠状态，即阻塞）
	goready(gp, skip+1) // 第二个参数 追踪ip寄存器的位置
}
```

```go
// 在一个无缓冲的通道或者空缓冲的同道中人发送或者接收数据，只需要向另一个groutine直接写入
func sendDirect(t *_type, sg *sudog, src unsafe.Pointer) {
	// 我们必须在一个函数调用中完成 sg.elem 指针的读取，否则当发生栈伸缩时，指针可能失效（被移动了）。
	dst := sg.elem
	// 为了确保发送的数据能够被立刻观察到，需要写屏障支持，执行写屏障，保证代码正确性
	typeBitsBulkBarrier(t, uintptr(dst), uintptr(src), t.size)
	memmove(dst, src, t.size)
}

```

```go
// 同sendDirect相反
func recvDirect(t *_type, sg *sudog, dst unsafe.Pointer) {
	src := sg.elem
	typeBitsBulkBarrier(t, uintptr(dst), uintptr(src), t.size)
	memmove(dst, src, t.size)
}
```

```go
func closechan(c *hchan) {
	// 向nil的channel发送数据会panic
	if c == nil {
		panic(plainError("close of nil channel"))
	}

	lock(&c.lock)
	// 关闭一个已经关闭的channel会panic
	if c.closed != 0 {
		unlock(&c.lock)
		panic(plainError("close of closed channel"))
	}

	c.closed = 1

    // 遍历所有队列，将所有的groutine都放在glist中，然后解锁，再唤醒这些groutine
	var glist gList 

	// 释放所有的接收者
	for {
		sg := c.recvq.dequeue()
		if sg == nil {
			break
		}
		if sg.elem != nil {
			typedmemclr(c.elemtype, sg.elem)
			sg.elem = nil
		}
		if sg.releasetime != 0 {
			sg.releasetime = cputicks()
		}
		gp := sg.g
		gp.param = nil
		if raceenabled {
			raceacquireg(gp, c.raceaddr())
		}
		glist.push(gp)
	}

	// 释放所有的发送者（会导致他们panic）
	for {
		sg := c.sendq.dequeue()
		if sg == nil {
			break
		}
		sg.elem = nil
		if sg.releasetime != 0 {
			sg.releasetime = cputicks()
		}
		gp := sg.g
		gp.param = nil
		if raceenabled {
			raceacquireg(gp, c.raceaddr())
		}
		glist.push(gp)
	}
	unlock(&c.lock)

	// 唤醒所有的goroutine
	for !glist.empty() {
		gp := glist.pop()
		gp.schedlink = 0
		goready(gp, 3)
	}
}

```