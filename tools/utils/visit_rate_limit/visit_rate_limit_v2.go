package main

import (
	"time"
)

type VisitLimitV2 interface {
	Sum(ip string) int64    // 返回ip在限制时间内访问的次数
	CheckIp(ip string) bool // 检查ip访问是否符合限制, 会调用sum函数来判断是否符合限制
	Update(ip string)       // 更新此ip访问
}

const (
	limitCount, limitDuration, nodeCount int64 = 100, 10 * 1000, 10
)

type IpOperation func(ipList map[string]*NodeList, blankList map[string]int64)

type VisitOperation struct {
	opCh   chan IpOperation
	stopCh chan struct{}
}

func NewVisitOperation() *VisitOperation {
	return &VisitOperation{
		opCh:   make(chan IpOperation),
		stopCh: make(chan struct{}),
	}
}

func (v *VisitOperation) Start() {
	go v.loop()
}

func (v *VisitOperation) Stop() {
	v.stopCh <- struct{}{}
}

func (v *VisitOperation) loop() {
	ipList := make(map[string]*NodeList, 500)
	blankList := make(map[string]int64, 500)
	for {
		select {
		case op := <-v.opCh:
			op(ipList, blankList)
		case <-v.stopCh:
			v.Stop() // todo 平滑的关闭？
			return
		}
	}
}

func (vo *VisitOperation) CheckIp(ip string) bool {
	validCh := make(chan bool)
	op := func(ipList map[string]*NodeList, blankList map[string]int64) {
		if _, ok := blankList[ip]; ok {
			validCh <- false
			return
		}
		if sum(ipList, ip) > limitCount {
			blankList[ip] = time.Now().Unix()
			validCh <- false
			return
		}
		validCh <- true
	}
	go func() {
		vo.opCh <- op
	}()
	return <-validCh
}

// for-select 中需要传递map，因此如果不加锁的话，需要同步，所以CheckIp()中不能调用Sum()，这样会出现死锁，所以这里独立出sum()
func sum(ipList map[string]*NodeList, ip string) int64 {
	if _, ok := ipList[ip]; !ok {
		return 0
	}
	head := ipList[ip].headNode
	sum := head.num
	for node := head.next; node != head; node = node.next {
		sum += node.num
	}
	return int64(sum)
}

func (vo *VisitOperation) Sum(ip string) int64 {
	sumCh := make(chan int64)
	op := func(ipList map[string]*NodeList, blankList map[string]int64) {
		sumCh <- sum(ipList, ip)
		return
	}
	go func() {
		vo.opCh <- op
	}()
	return <-sumCh
}

func (vo *VisitOperation) Update(ip string) {
	op := func(ipList map[string]*NodeList, blankList map[string]int64) {
		now := time.Now().UnixNano() / 1e6
		// 如果是第一次访问，则初始化环形链表，并给head节点的访问量+1
		if _, ok := ipList[ip]; !ok {
			ipList[ip] = NewNodeList(now)
			ipList[ip].init()
			return
		}
		durPerNode := limitDuration / nodeCount
		// 如果不是第一次访问，判断当前时间与head节点的时间差，如果不超过限制时间则直接在对应的节点的访问量+1
		durHead := now - ipList[ip].headTime
		headNode := ipList[ip].headNode
		if durHead < limitDuration {
			nodeNum := (limitDuration - durHead) / durPerNode
			for i := 0; i < int(nodeNum); i++ {
				headNode = headNode.next
			}
			headNode.AddOneVisit()
			return
		}
		// 如果当前时间与head节点的时间差超过了限制时间的两倍，也就是说前一段时间内没有访问，则重新计算——直接清空整个链表，重置首节点时间，并将首节点访问量记为1
		if durHead >= 2*limitDuration {
			for i := 0; i < int(nodeNum); i++ {
				headNode.num = 0
				headNode = headNode.next
			}
			ipList[ip].init()
			return
		}
		// 如果当前时间与head节点的时间差在限制时间的一到两倍之间，则有些节点“过期”，需要选择新的首节点和首节点时间，并将过期的节点覆盖掉
		expireCount := (durHead-limitDuration)/durPerNode + 1 // 距离head节点刚好过期时，(durHead - limitDuration) / durPerNode 为0，需要加1
		for i := 0; i < int(expireCount)-1; i++ {
			headNode.num = 0
			headNode = headNode.next
		}
		headNode.num = 1
		ipList[ip].headNode = headNode.next
		ipList[ip].headTime = ipList[ip].headTime + expireCount*durPerNode
	}
	vo.opCh <- op
}
