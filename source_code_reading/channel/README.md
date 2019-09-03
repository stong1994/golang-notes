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
	sendq    waitq  // 等待法发送元素的列表
	lock mutex // 锁
}

type waitq struct {
	first *sudog // 等待列表中的第一个
	last  *sudog // 等待列表中的最后一个
}

// sudog 和 g 之间的关系是一对多,即一个g可能有多个sudog.比如一个g可能在多个等待列表中,那么每个都会产生一个sudog
// 其元素都被其所属的hchan中的lock保护
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
    // 如果channel为空,说明channel已经关闭或者只是声明,如果处于非阻塞状态,直接返回false
    //                                           如果处于阻塞状态,则调用gopark,直接panic掉,因为deadlock
	if c == nil {
		if !block {
			return false
		}
		gopark(nil, nil, waitReasonChanSendNilChan, traceEvGoStop, 2)
		throw("unreachable")
	}

	if debugChan {
		print("chansend: chan=", c, "\n")
	}

	if raceenabled {
		racereadpc(c.raceaddr(), callerpc, funcPC(chansend))
	}

	// Fast path: 不获取锁京能检查失败的非阻塞操作
	//
	// 如果channel没有关闭,尼玛
	// After observing that the channel is not closed, we observe that the channel is
	// not ready for sending. Each of these observations is a single word-sized read
	// (first c.closed and second c.recvq.first or c.qcount depending on kind of channel).
	// Because a closed channel cannot transition from 'ready for sending' to
	// 'not ready for sending', even if the channel is closed between the two observations,
	// they imply a moment between the two when the channel was both not yet closed
	// and not ready for sending. We behave as if we observed the channel at that moment,
	// and report that the send cannot proceed.
	//
	// It is okay if the reads are reordered here: if we observe that the channel is not
	// ready for sending and then observe that it is not closed, that implies that the
	// channel wasn't closed during the first observation.
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
	
    // 从等待接收的队列中拿出一个sudog,并且直接进行拷贝,而不通过buffer和锁
	if sg := c.recvq.dequeue(); sg != nil {
		// Found a waiting receiver. We pass the value we want to send
		// directly to the receiver, bypassing the channel buffer (if any).
		send(c, sg, ep, func() { unlock(&c.lock) }, 3)
		return true
	}
    
	// 如果等待队列为空,那么就需要等待其他goroutine来接收,查看ring buffer是否已满
	// 如果环形队列的大小大于当前队列中的数据,说明ring buffer还有多余的空间,则将当前发送的值填充到ring buffer
	if c.qcount < c.dataqsiz {
		// Space is available in the channel buffer. Enqueue the element to send.
		qp := chanbuf(c, c.sendx)
		if raceenabled {
			raceacquire(qp)
			racerelease(qp)
		}
		typedmemmove(c.elemtype, qp, ep)
		c.sendx++
		if c.sendx == c.dataqsiz { // 环形队列,保证FIFO
			c.sendx = 0
		}
		c.qcount++
		unlock(&c.lock)
		return true
	}

	// 如果ring buffer已满,并且为非阻塞状态,则直接返回false? 
	if !block {
		unlock(&c.lock)
		return false
	}

	
	// 阻塞住当前的channel 什么时候环形?
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
	goparkunlock(&c.lock, waitReasonChanSend, traceEvGoBlockSend, 3)
	// Ensure the value being sent is kept alive until the
	// receiver copies it out. The sudog has a pointer to the
	// stack object, but sudogs aren't considered as roots of the
	// stack tracer.
	KeepAlive(ep)

	// someone woke us up.
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
	releaseSudog(mysg)
	return true
}
```