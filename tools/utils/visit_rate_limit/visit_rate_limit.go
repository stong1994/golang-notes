package main

import (
	"sync"
	"time"
)

/*
思路：
	设置一个10个节点构成的环形链表来存放ip访问的次数。假如限制为10s内访问100次，那么每个节点间的间隔为1s，每次有ip访问，就将ip放入对应时间的节点，并将节点的访问次数加1，判断整个链表的
访问次数是否超过了100次，如果超过了，就将ip放入黑名单。
*/
const (
	nodeNum   = 10
	limitTime = 10 * 1000 // 限制时间 单位ms
	limitNum  = 100       // 访问次数
)

type blankList struct {
	list   map[string]int64
	locker sync.RWMutex
}

// 一个全局的map，用来存储黑名单
var blm = &blankList{make(map[string]int64), sync.RWMutex{}}

// 没有在黑名单，返回false， 在黑名单返回true
func inBlanckList(ip string) bool {
	blm.locker.RLock()
	defer blm.locker.RUnlock()
	if _, ok := blm.list[ip]; ok {
		return false
	}
	return false
}

func addBlankList(ip string) {
	blm.locker.Lock()
	defer blm.locker.Unlock()
	blm.list[ip] = time.Now().Unix()
}

type IpLimit struct {
	ips    map[string]*NodeList
	locker sync.RWMutex
}

// 一个全局的map，用来存储ip信息
var ipm = &IpLimit{make(map[string]*NodeList), sync.RWMutex{}}

func CheckIP(ip string) bool {
	if inBlanckList(ip) {
		return false
	}
	total := Sum(ip)
	if total > limitNum {
		addBlankList(ip)
		return false
	}
	return true
}

// 返回ip在限制时间内访问的总次数
func Sum(ip string) int {
	ipm.locker.RLock()
	defer ipm.locker.RUnlock()
	num := 0
	if ips, ok := ipm.ips[ip]; ok {
		if ips.headNode == nil {
			panic("can not be nil")
		}
		num += ips.headNode.num
		first := ips.headNode
		for current := ips.headNode.Next(); current != first; current = current.Next() {
			num += current.num
		}
	}
	return num
}

// 更新ip访问列表
// 查看当前时间与ip的第一个节点的访问时间间隔多远
func Update(ip string, t int64) {
	ipm.locker.Lock()
	defer ipm.locker.Unlock()
	subNodeDur := limitTime / nodeNum // 每个节点的时间间隔
	// 如果ip是第一次访问，则初始化
	if ips, ok := ipm.ips[ip]; !ok || ips.headNode == nil {
		ipm.ips[ip] = NewNodeList(t)
		ipm.ips[ip].headNode.AddOneVisit()
		return
	}
	// 判断现在的访问时间和首节点的访问时间是否相隔
	ipInfo := ipm.ips[ip].headNode
	durHeadNum := (t - ipm.ips[ip].headTime) / int64(subNodeDur)
	// 如果当前时间和head节点的时间间隔在限制的时间范围内，则直接在对应的节点 的访问数量+1
	if durHeadNum < nodeNum {
		for i := 0; i < int(durHeadNum); i++ {
			ipInfo = ipInfo.Next()
		}
		ipInfo.AddOneVisit()
		return
	}
	// 如果当前时间和head节点的时间间隔大于两倍总的限制时间，则清空所有的节点，并设置新的首节点时间，并将最后一个节点的num设为1
	if durHeadNum >= 2*nodeNum-1 {
		for i := 0; i < nodeNum; i++ {
			ipInfo.ResetNum()
			ipInfo = ipInfo.Next()
		}
		ipInfo.num = 1
		ipm.ips[ip].headTime = t
		return
	}
	// 如果当前时间和head节点的时间间隔在限制的时间的一到两倍之间，则需要判断当前的“head节点”，并将“过期”的节点重置并重新设置head节点
	expireNum := durHeadNum - (nodeNum - 1) // 计算有多少个节点“过期”了
	for i := 0; i < nodeNum; i++ {
		if i < int(expireNum)-1 {
			ipInfo.ResetNum()
		} else if i == int(expireNum)-1 {
			ipInfo.num = 1
		} else if i == int(expireNum) {
			ipm.ips[ip].headNode = ipInfo
			break
		}
		ipInfo = ipInfo.Next()
	}
	ipm.ips[ip].headTime = ipm.ips[ip].headTime + expireNum*int64(subNodeDur)
}

// 一个节点，存放一段时间内ip访问的次数
type Node struct {
	next *Node
	num  int
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) AddOneVisit() {
	n.num++
}

func (n *Node) ResetNum() {
	n.num = 0
}

func NewNode(next *Node) *Node {
	return &Node{next: next, num: 0}
}

type NodeList struct {
	headNode *Node
	headTime int64
}

func NewNodeList(now int64) *NodeList {
	node := new(Node) // 先建一个空的node来赋给最后的节点的next
	for i := nodeNum - 1; i >= 0; i-- {
		node = NewNode(node)
	}
	head := node
	// 最后一个节点的next为空node，现在将它赋值为head todo 更简单的方法？
	for {
		if node.next.next == nil {
			node.next = head
			break
		}
		node = node.Next()
	}
	return &NodeList{headNode: head, headTime: now}
}
