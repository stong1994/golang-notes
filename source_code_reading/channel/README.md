# channel源码阅读

*go version go1.12.5 linux/amd64*

**代码位置**: runtime/chan.go  
 一共约700行
 
 
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