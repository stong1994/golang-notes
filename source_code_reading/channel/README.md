# channel源码阅读

*go version go1.12.5 linux/amd64*

**代码位置**: runtime/chan.go  
 一共约700行
 
 #### go代码函数 对应 编译后的函数
 ```
1. make(chan interface{}, size) ⇒ runtime.makechan(interface{}, size)

   make(chan interface{})       ⇒ runtime.makechan(interface{}, 0)

2. ch <- v                      ⇒ runtime.chansend1(ch, &v)

3. v <- ch                      ⇒ runtime.chanrecv1(ch, &v)

   v, ok <- ch                  ⇒ ok := runtime.chanrecv2(ch, &v)

4. close(ch)                    ⇒ runtime.closechan(ch)
```
 
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
// c为要发送到的channel, ep为发送的数据,block为是否需要阻塞, 返回值为是否发送成功
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
	gp := sg.g // 获取到sudog所在的goroutine
	unlockf() 
	gp.param = unsafe.Pointer(sg)
	if sg.releasetime != 0 {
		sg.releasetime = cputicks()
	}
	// 唤醒goroutine（由于其一直在等待数据被接收，因此一直处在睡眠状态，即阻塞）
	goready(gp, skip+1) // 第二个参数 追踪ip寄存器的位置
}
```

```go
// 在一个无缓冲的通道或者空缓冲的同道中人发送或者接收数据，只需要向另一个goroutine直接写入
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
#### 关闭channel
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

    // 遍历所有队列，将所有的goroutine都放在glist中，然后解锁，再唤醒这些goroutine
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

```go
// chanrecv 在c上接收到值,并将接收到的值写入ep,ep可以为空,在这种情况下接收到的值会被忽略掉.
// 如果block为flase并且没有元素可用,返回false,false, 表示执行失败但是没有被接收
// 如果c已经关闭了,把ep置空,返回true,false, 表示执行成功但是没有被接收
// 其他的话,赋值给ep,并返回true,true
// 一个非空的ep必须指向堆或者调用者的栈
// received表示是否接收到值, 即 _, ok := <-ch 中的ok
func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool) {
    // 向一个为nil的channel上发送数据,如果为非阻塞,那么直接返回false,false,如果为阻塞,那么阻塞当前goroutine,但是没有办法唤醒(第一个参数为nil),导致死锁
	if c == nil {
		if !block { 
			return
		}
		gopark(nil, nil, waitReasonChanReceiveNilChan, traceEvGoStop, 2)
		throw("unreachable")
	}

	// Fast path: 在不用锁的情况下对非阻塞的操作检查失败
	if !block && (c.dataqsiz == 0 && c.sendq.first == nil ||
		c.dataqsiz > 0 && atomic.Loaduint(&c.qcount) == 0) &&
		atomic.Load(&c.closed) == 0 {
		return
	}

	var t0 int64
	if blockprofilerate > 0 {
		t0 = cputicks()
	}

	lock(&c.lock)
    
	// 如果channel已经关闭,并且队列中没有元素,那么直接返回ture,false
	if c.closed != 0 && c.qcount == 0 {
		if raceenabled {
			raceacquire(c.raceaddr())
		}
		unlock(&c.lock)
		if ep != nil {
			typedmemclr(c.elemtype, ep)
		}
		return true, false
	}
	
	// 如果channel中的等待发送的队列中有数据,那么获取队列中的一个数据,并直接赋值给ep
	if sg := c.sendq.dequeue(); sg != nil {
		recv(c, sg, ep, func() { unlock(&c.lock) }, 3)
		return true, true
	}

	// 如果channel中的等待发送的队列中没有数据,但是等待接收的队列中有数据,那么填充buffer
	if c.qcount > 0 {
		// Receive directly from queue
		qp := chanbuf(c, c.recvx)
		if raceenabled {
			raceacquire(qp)
			racerelease(qp)
		}
		if ep != nil {
			typedmemmove(c.elemtype, ep, qp)
		}
		typedmemclr(c.elemtype, qp)
		c.recvx++
		if c.recvx == c.dataqsiz {
			c.recvx = 0
		}
		c.qcount--
		unlock(&c.lock)
		return true, true
	}
	
    // 如果队列中没有数据,那么说明将数据发送到了一个空的channel,如果为非阻塞,那么直接返回false,false
	if !block {
		unlock(&c.lock)
		return false, false
	}

	// 如果队列中没有数据,并且是阻塞,那么阻塞当前goroutine,直到有发送者接收该值
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
	gp.waiting = mysg
	mysg.g = gp
	mysg.isSelect = false
	mysg.c = c
	gp.param = nil
	c.recvq.enqueue(mysg)
	goparkunlock(&c.lock, waitReasonChanReceive, traceEvGoBlockRecv, 3) // 阻塞当前goroutine

	// 等待唤醒
	if mysg != gp.waiting {
		throw("G waiting list is corrupted")
	}
	gp.waiting = nil
	if mysg.releasetime > 0 {
		blockevent(mysg.releasetime-t0, 2)
	}
	closed := gp.param == nil
	gp.param = nil
	mysg.c = nil
	releaseSudog(mysg) // 释放掉复制的sudog
	return true, !closed // 如果channel关闭后,会接收到 _, false := <-ch
}
```

#### 入列
```go
func (q *waitq) enqueue(sgp *sudog) {
	sgp.next = nil
	x := q.last
	if x == nil { 
		sgp.prev = nil
		q.first = sgp
		q.last = sgp
		return
	}
	sgp.prev = x
	x.next = sgp
	q.last = sgp
}
```
#### 出列
```go
func (q *waitq) dequeue() *sudog {
	for {
		sgp := q.first // 获取第一个元素,如果为空,那么说明整个队列都为空,返回nil
		if sgp == nil {
			return nil
		}
		y := sgp.next // 找到第二个元素,如果为空,那么说明在这次出列操作结束后,队列为空
		if y == nil {
			q.first = nil
			q.last = nil
		} else { // 第一个元素出列后,将第二个元素设置为首元素
			y.prev = nil
			q.first = y
			sgp.next = nil // mark as removed (see dequeueSudog)
		}

		// 如果一个goroutine是因为select放进到的队列中,在goroutine之间有一个小的窗口在一些特殊情况下被唤醒,并且它能拿到channel的锁.
		// 一旦它拿到了锁,在G的结构体中有一个标志,用来告诉我们在什么时候有其他人得到了这个竞争条件,并通知该goroutine,除非该goroutine还没有将它自己移出队列
		if sgp.isSelect {
			if !atomic.Cas(&sgp.g.selectDone, 0, 1) {
				continue
			}
		}

		return sgp
	}
}
```

该文件中还有一些优化select的代码
```go
// compiler implements
//
//	select {
//	case c <- v:
//		... foo
//	default:
//		... bar
//	}
//
// as
//
//	if selectnbsend(c, v) {
//		... foo
//	} else {
//		... bar
//	}
//
func selectnbsend(c *hchan, elem unsafe.Pointer) (selected bool) {
	return chansend(c, elem, false, getcallerpc())
}

// compiler implements
//
//	select {
//	case v = <-c:
//		... foo
//	default:
//		... bar
//	}
//
// as
//
//	if selectnbrecv(&v, c) {
//		... foo
//	} else {
//		... bar
//	}
//
func selectnbrecv(elem unsafe.Pointer, c *hchan) (selected bool) {
	selected, _ = chanrecv(c, elem, false)
	return
}

// compiler implements
//
//	select {
//	case v, ok = <-c:
//		... foo
//	default:
//		... bar
//	}
//
// as
//
//	if c != nil && selectnbrecv2(&v, &ok, c) {
//		... foo
//	} else {
//		... bar
//	}
//
func selectnbrecv2(elem unsafe.Pointer, received *bool, c *hchan) (selected bool) {
	// TODO(khr): just return 2 values from this function, now that it is in Go.
	selected, *received = chanrecv(c, elem, false)
	return
}
```
由上可知,select和channel息息相关,并且编译器会将select优化为if-else  
显然上述情况并没有涵盖所有select的清空,所以select会单独来写分析.

#### 总结
1. channel中的缓存为ring buffer
2. channel中有两个字段用来存放等待接收的sudog和等待发送的sudog
3. 缓冲中存放的是数据，而两个等待队列中存放的是sudog
4. 访问缓存中的数据为FIFO
5. 向一个为nil的channel中接收数据，如果是在select中，会得到零值（zero,false := <- ch）,但是在非select中则会死锁。这是由参数block决定的。
6. channel为nil的情况：当一个channel只是被声明。关闭状态的channel不是nil，用closed字段标识；用make构造的channel也不是nil
7. 关闭channel
    - 上锁
    - 遍历两个等待队列，释放sudog中的元素，然后将sudog所在的g放到一个数组中
    - 解锁
    - 唤醒数组中的g
8. 关闭channel，会导致等待发送的队列中的goroutine panic；关闭channel，并没有清空缓存中的数据，
9. 关闭channel，在等待接收的队列中，如果channel还有缓冲，返回缓存值和true，如果没有缓冲，返回零值和false。因此一个向关闭的channel中获取值，不一定为零值和false

#### 资料
1. [夜读分享者的PPT](https://docs.google.com/presentation/d/18_9LcMc8u93aITZ6DqeUfRvOcHQYj2gwxhskf0XPX2U/edit#slide=id.gc6f919934_0_0)
2. [夜读分享者的视频](https://www.bilibili.com/video/av64926593?t=4703)
3. [understanding channel-英文ppt](https://speakerdeck.com/kavya719/understanding-channels)
4. [深度解密Go语言之channel——饶大](https://mp.weixin.qq.com/s?__biz=MjM5MDUwNTQwMQ==&mid=2257483870&idx=1&sn=1c61a1f530b3e52d801a7916065f3eec&chksm=a5391888924e919e39ddb8f017b572fd6f199184d6ad85c4ae4dacf2749312b70717d7beee44&mpshare=1&scene=1&srcid=&key=034516426b2066d0f70d525bf66160c508cad2ca843ffddfa4ee6bd2edc545645b98d4bb5ce66bab8cfa81546462f00c8fe47c5e797a7ccbedf92eeb800ec864b921273b277153f1b88730e5d884ce91&ascene=1&uin=MjEwMjA3MTA2NQ%3D%3D&devicetype=Windows+10&version=62060834&lang=zh_CN&pass_ticket=efPNLRctgYmpKMszLe6OEG5z5f8en%2BzyyWAgphiEkVcPy2arsWBGjQPgMH5xDDSU)
5. [大彬的源码读后感](http://lessisbetter.site/2019/03/03/golang-channel-design-and-source/)