# SELECT源码

对几种select进行了汇编，结果如下：
```
1. select {} => runtime.block()
2. select {
   	    case ch <- struct{}{}:   => runtime.chansend1()
   }
3. select {
        case ch <- struct{}{}:   => runtime.selectnbsend()
        default:
   }
4. select {
   	    case <-ch1:             => runtime.chanrecv1()
   }
5. select {
        case <-ch1:             =>  runtime.selectnbrecv()
        default:
    }
6. select {
        case <-ch1:
        case ch2 <- struct{}{}:     => runtime.selectgo()
    }
```

先看简单的代码: `runtime/chan.go`
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
**结论**：
1. 如果为空的select，直接阻塞
2. 如果有一个case，在编译时都会直接调用channel中的函数，进行判断case能否立即接收/发送
3. 如果有一个case + default，那么为if-else结构，即如果case能够立即接收/发送，那么执行case，否则执行default
4. 如果有多个case，调用selectgo，即下面的代码分析。

### selectgo
#### 描述case的结构体
```go
type scase struct {
	c           *hchan         // chan
	elem        unsafe.Pointer // data element
	kind        uint16
	pc          uintptr // race pc (for race detector / msan)
	releasetime int64
}
```

#### select语句的实现
```go
// runtime/select.go
// selectgo实现了select语句
//
// cas0 指向了一个scase的数组，长度为ncases
// order0 指向了一个类型为[2*ncases]uint16的数组，他们都存在goroutine的栈上(不管selectgo中的任何逃逸).
//
// selectgo 返回了被选中的scase的索引, 并且匹配并调用他们各自的 select{recv,send,default}
// 并且，如果被选中的scase是一个接收操作，它还能体现出是否接收到一个值
func selectgo(cas0 *scase, order0 *uint16, ncases int) (int, bool) {
    // (*[1 << 16]scase)为数组指针，unsafe.Pointer(cas0)为case0的指针地址，所以cas1为scase的数组指针地址？
	cas1 := (*[1 << 16]scase)(unsafe.Pointer(cas0))
	order1 := (*[1 << 17]uint16)(unsafe.Pointer(order0))

	scases := cas1[:ncases:ncases]
	pollorder := order1[:ncases:ncases]
	lockorder := order1[ncases:][:ncases:ncases]

	// 如果case的channel为nil，并且case不是default，则将case初始化为空结构体
	// 所以下边的逻辑代码中讲没有空的channel
	// 所以在一个for-select结构中，如果当命中一个case后，我们把case所在的channel置为nil，则下次for循环不再访问该case
	for i := range scases {
		cas := &scases[i]
		if cas.c == nil && cas.kind != caseDefault {
			*cas = scase{}
		}
	}

	var t0 int64
	if blockprofilerate > 0 {
		t0 = cputicks()
		for i := 0; i < ncases; i++ {
			scases[i].releasetime = -1
		}
	}

	// 编译器将静态地只有0或1个case加上默认值的选项重写为更简单的构造。
	// 在这里，我们可以使用这么小的ncase值的唯一方法是对于一个更大的select，其中大多数通道由于nil而忽略
	// The only way we can end up with such small sel.ncase
	// values here is for a larger select in which most channels
	// have been nilled out. 
	// 通用代码正确地处理这些情况，它们需要进行优化(并且需要进行测试)。

	// 生成随机排列顺序——case的随机选择
	for i := 1; i < ncases; i++ {
		// 获取随机数
		// https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
		j := fastrandn(uint32(i + 1))
		pollorder[i] = pollorder[j]
		pollorder[j] = uint16(i)
	}

	// 根据hchan地址对case进行排序，来得到锁定顺序
	// 简单的堆排序, 以保证n log n的时间和不变的堆栈占用空间。
	for i := 0; i < ncases; i++ {
		j := i
		// 以 pollorder开始， 在相同的channel上交换case
		c := scases[pollorder[i]].c
		for j > 0 && scases[lockorder[(j-1)/2]].c.sortkey() < c.sortkey() {
			k := (j - 1) / 2
			lockorder[j] = lockorder[k]
			j = k
		}
		lockorder[j] = pollorder[i]
	}
	for i := ncases - 1; i >= 0; i-- {
		o := lockorder[i]
		c := scases[o].c
		lockorder[i] = lockorder[0]
		j := 0
		for {
			k := j*2 + 1
			if k >= i {
				break
			}
			if k+1 < i && scases[lockorder[k]].c.sortkey() < scases[lockorder[k+1]].c.sortkey() {
				k++
			}
			if c.sortkey() < scases[lockorder[k]].c.sortkey() {
				lockorder[j] = lockorder[k]
				j = k
				continue
			}
			break
		}
		lockorder[j] = o
	}

	// 锁定select中涉及到的所有channel
	sellock(scases, lockorder)

	var (
		gp     *g
		sg     *sudog
		c      *hchan
		k      *scase
		sglist *sudog
		sgnext *sudog
		qp     unsafe.Pointer
		nextp  **sudog
	)

loop:
	// pass 1 - look for something already waiting
	var dfli int // default case 的索引
	var dfl *scase // default case
	var casi int // for循环中当前的case的索引
	var cas *scase // for循环中当前的case
	var recvOK bool
	for i := 0; i < ncases; i++ {
		casi = int(pollorder[i])
		cas = &scases[casi]
		c = cas.c

		switch cas.kind {
		case caseNil:
			continue
        // 当前case是接收case
        // 如果channel中的等待发送队列中有数据，那么直接调用chan文件中的recv函数（赋值并唤醒goroutine），然后返回当前的case和true
        // 如果channel中的缓存中有数据，那么从缓存中获取数据，赋值并返回当前case和true
        // 如果channel已经关闭,那么直接返回当前case和false
		case caseRecv:
			sg = c.sendq.dequeue()
			if sg != nil {
				goto recv
			}
			if c.qcount > 0 {
				goto bufrecv
			}
			if c.closed != 0 {
				goto rclose
			}
        // 当前case是发送case
        // 如果channel已经关闭,那么直接panic
        // 如果channel中的等待接收队列中有数据，那么直接调用chan文件中的send函数（赋值并唤醒goroutine），然后返回当前的case和true
        // 如果channel中的缓存中有数据，那么从缓存中获取数据，赋值并返回当前case和true
		case caseSend:
			if raceenabled {
				racereadpc(c.raceaddr(), cas.pc, chansendpc)
			}
			if c.closed != 0 {
				goto sclose
			}
			sg = c.recvq.dequeue()
			if sg != nil {
				goto send
			}
			if c.qcount < c.dataqsiz {
				goto bufsend
			}
        // 当前case是default case，赋值，并继续下个循环
		case caseDefault:
			dfli = casi
			dfl = cas
		}
	}
    // 如果存在default case并且其他case都没有命中， 那么将返回default case
	if dfl != nil {
		selunlock(scases, lockorder)
		casi = dfli
		cas = dfl
		goto retc
	}

	// pass 2 - enqueue on all chans
	// 如果不存在default case，并且其他case都没有命中，则需要阻塞当前goroutine
	// 遍历lockorder，并将所有的发送或者接收的case所在的sudog放到case所在的channel的等待队列中
	gp = getg() // 当前goroutine
	if gp.waiting != nil {
		throw("gp.waiting != nil")
	}
	nextp = &gp.waiting
	for _, casei := range lockorder {
		casi = int(casei)
		cas = &scases[casi]
		if cas.kind == caseNil {
			continue
		}
		c = cas.c
		sg := acquireSudog()
		sg.g = gp
		sg.isSelect = true
		// No stack splits between assigning elem and enqueuing
		// sg on gp.waiting where copystack can find it.
		sg.elem = cas.elem
		sg.releasetime = 0
		if t0 != 0 {
			sg.releasetime = -1
		}
		sg.c = c
		// Construct waiting list in lock order.
		*nextp = sg
		nextp = &sg.waitlink

		switch cas.kind {
		case caseRecv:
			c.recvq.enqueue(sg)

		case caseSend:
			c.sendq.enqueue(sg)
		}
	}

	// 等待被唤醒
	gp.param = nil
	gopark(selparkcommit, nil, waitReasonSelect, traceEvGoBlockSelect, 1)
    // 唤醒后的位置
	sellock(scases, lockorder)

	gp.selectDone = 0
	sg = (*sudog)(gp.param) // 当前goroutine的sudog
	gp.param = nil

	// pass 3 - 从没有成功的channel中出列,否则他们会一直停留在channel中
	// 如果有成功的case，记录下来。
	// 我们以单链表的方式将sudog以锁的顺序连起来
	casi = -1
	cas = nil
	sglist = gp.waiting // 当前goroutine正在等待的sudog
	// 在从gp.waiting断开连接之前，清除所有elem。 
	for sg1 := gp.waiting; sg1 != nil; sg1 = sg1.waitlink {
		sg1.isSelect = false
		sg1.elem = nil
		sg1.c = nil
	}
	// 清空gp.waiting 
	gp.waiting = nil

	for _, casei := range lockorder {
		k = &scases[casei]
		if k.kind == caseNil {
			continue
		}
		if sglist.releasetime > 0 {
			k.releasetime = sglist.releasetime
		}
		// 如果当前goroutine的sudog和当前goroutine等待的sudog相等，说明当前goroutine的sudog就是从被唤醒的g中出列得到的
		if sg == sglist {
			// sg has already been dequeued by the G that woke us up.
			casi = int(casei)
			cas = k
		} else {
			c = k.c
			if k.kind == caseSend {
				c.sendq.dequeueSudoG(sglist)
			} else {
				c.recvq.dequeueSudoG(sglist)
			}
		}
		sgnext = sglist.waitlink
		sglist.waitlink = nil
		// 释放sudog
		releaseSudog(sglist)
		sglist = sgnext
	}

	if cas == nil {
		// 当select中涉及到的channel被关闭时，g被唤醒，gp.param可能为nil，这种情况下 cas也为nil
		// 循环并且重复执行操作是简单的
		// 我们可以得到channel现在被关闭了
		// 也许以后我们能够更简明的标记关闭，但是我们必须区分在发送时关闭还是在接收时关闭
		// 只是重复检查而不向上述说的做区分是简单的
		// 我们知道一些东西被关闭了，并且不会被解除关闭状态，所以我们不会再次阻塞
		goto loop
	}

	c = cas.c

	if cas.kind == caseRecv {
		recvOK = true
	}

	if raceenabled {
		if cas.kind == caseRecv && cas.elem != nil {
			raceWriteObjectPC(c.elemtype, cas.elem, cas.pc, chanrecvpc)
		} else if cas.kind == caseSend {
			raceReadObjectPC(c.elemtype, cas.elem, cas.pc, chansendpc)
		}
	}
	if msanenabled {
		if cas.kind == caseRecv && cas.elem != nil {
			msanwrite(cas.elem, c.elemtype.size)
		} else if cas.kind == caseSend {
			msanread(cas.elem, c.elemtype.size)
		}
	}

	selunlock(scases, lockorder)
	goto retc

bufrecv:
	// can receive from buffer
	if raceenabled {
		if cas.elem != nil {
			raceWriteObjectPC(c.elemtype, cas.elem, cas.pc, chanrecvpc)
		}
		raceacquire(chanbuf(c, c.recvx))
		racerelease(chanbuf(c, c.recvx))
	}
	if msanenabled && cas.elem != nil {
		msanwrite(cas.elem, c.elemtype.size)
	}
	recvOK = true
	qp = chanbuf(c, c.recvx)
	if cas.elem != nil {
		typedmemmove(c.elemtype, cas.elem, qp)
	}
	typedmemclr(c.elemtype, qp)
	c.recvx++
	if c.recvx == c.dataqsiz {
		c.recvx = 0
	}
	c.qcount--
	selunlock(scases, lockorder)
	goto retc

bufsend:
	// can send to buffer
	if raceenabled {
		raceacquire(chanbuf(c, c.sendx))
		racerelease(chanbuf(c, c.sendx))
		raceReadObjectPC(c.elemtype, cas.elem, cas.pc, chansendpc)
	}
	if msanenabled {
		msanread(cas.elem, c.elemtype.size)
	}
	typedmemmove(c.elemtype, chanbuf(c, c.sendx), cas.elem)
	c.sendx++
	if c.sendx == c.dataqsiz {
		c.sendx = 0
	}
	c.qcount++
	selunlock(scases, lockorder)
	goto retc

recv:
	// can receive from sleeping sender (sg)
	recv(c, sg, cas.elem, func() { selunlock(scases, lockorder) }, 2)
	if debugSelect {
		print("syncrecv: cas0=", cas0, " c=", c, "\n")
	}
	recvOK = true
	goto retc

rclose:
	// read at end of closed channel
	selunlock(scases, lockorder)
	recvOK = false
	if cas.elem != nil {
		typedmemclr(c.elemtype, cas.elem)
	}
	if raceenabled {
		raceacquire(c.raceaddr())
	}
	goto retc

send:
	// can send to a sleeping receiver (sg)
	if raceenabled {
		raceReadObjectPC(c.elemtype, cas.elem, cas.pc, chansendpc)
	}
	if msanenabled {
		msanread(cas.elem, c.elemtype.size)
	}
	send(c, sg, cas.elem, func() { selunlock(scases, lockorder) }, 2)
	if debugSelect {
		print("syncsend: cas0=", cas0, " c=", c, "\n")
	}
	goto retc

retc:
	if cas.releasetime > 0 {
		blockevent(cas.releasetime-t0, 1)
	}
	return casi, recvOK

sclose:
	// send on closed channel
	selunlock(scases, lockorder)
	panic(plainError("send on closed channel"))
}
```

#### select的流程
1. 将非default的并且所在channel为nil的case设为case的空结构体，便于下面操作不会出现为nil的channel
2. 对case进行随机排序，等到一个case的数组pollorder
3. 根据case所在channel的地址进行排序，得到一个锁定的case数组lockorder
4. 对select涉及到的channel进行锁定，开始循环
    1. 找到已经准备好的case并返回case  
        1. 当前case是接收case  
            1. 如果channel中的等待发送队列中有数据，那么直接调用chan文件中的recv函数（赋值并唤醒goroutine），然后返回当前的case和true
            2. 如果channel中的缓存中有数据，那么从缓存中获取数据，赋值并返回当前case和true
            3. 如果channel已经关闭,那么直接返回当前case和false
        2. 当前case是发送case
            1. 如果channel已经关闭,那么直接panic
            2. 如果channel中的等待接收队列中有数据，那么直接调用chan文件中的send函数（赋值并唤醒goroutine），然后返回当前的case和true
            3. 如果channel中的缓存中有数据，那么从缓存中获取数据，赋值并返回当前case和true
        3. 当前case是default case，进入下个循环，如果遍历完仍没有case能够命中，则使用default case
    2. 如果不存在default case，并且其他case都没有命中，则需要阻塞当前goroutine
        1. 遍历lockorder，并将所有的发送或者接收的case所在的sudog放到case所在的channel的等待队列中
    3. 被唤醒后
        1. 遍历case，将没有成功接收/发送的case从他们的channel中取出，并释放
        2. 如果没有命中的case（如果channel被关闭，会被唤醒,如果当前goroutine没有任何sudog，就不会命中case），即没有一个case成功接收/发送，那么就进入下个循环
        3. 如果有命中的case，那么返回该case
        
### 知识点
0. 简单的select(单个case)在编译时会直接访问channel中的函数，来判断能否立即接收/发送
1. 如果select中有多个case符合命中条件，那么命中的case是随机的，这是因为select中的case在初始化的时候进行了随机排序。（涉及到随机数和堆排序，还没搞很清楚）
2. 在for-select结构中，如果在case接收到数据后，将case所在的channel值为nil（不是关闭），那么下次for循环就不会再访问该case
    ```go
    if cas.c == nil && cas.kind != caseDefault {
        *cas = scase{}
    }
    // ...
    // ...
    // for 循环case
    if k.kind == caseNil { // caseNil为0，因为上边代码将case设置为了空的结构体，因此caseNil为0
        continue
    }
    ```
3. select也要进行循环。如果不存在符合条件的case并且没有default，则阻塞当前g。有一种情况是被唤醒后也没有命中case（case所在的channel被关闭且gp.param==nil）,这种情况下需要再次进入循环
