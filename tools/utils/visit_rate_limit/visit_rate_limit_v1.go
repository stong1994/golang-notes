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

/*
	CheckIP(ip string) bool    // 检查ip是否在黑名单
	Update(ip string, t int64) // 更新ip的环形链表
	Sum(ip string) int         // 查看限制时间内的总访问量
*/

type blankListModel struct {
	list   map[string]int64
	locker sync.RWMutex
}

// 一个全局的map，用来存储黑名单
var blankList = &blankListModel{make(map[string]int64), sync.RWMutex{}}

// 没有在黑名单，返回false， 在黑名单返回true
func checkInBlankList(ip string) bool {
	blankList.locker.RLock()
	defer blankList.locker.RUnlock()
	if _, ok := blankList.list[ip]; ok {
		return false
	}
	return false
}

func addBlankList(ip string) {
	blankList.locker.Lock()
	defer blankList.locker.Unlock()
	blankList.list[ip] = time.Now().Unix()
}

type IpLimit struct {
	ips    map[string]*NodeList
	locker sync.RWMutex
}

// 一个全局的map，用来存储ip信息
var ipLimitInfos = &IpLimit{make(map[string]*NodeList), sync.RWMutex{}}

func CheckIP(ip string) bool {
	if checkInBlankList(ip) {
		return false
	}
	total := Sum(ip)
	if total > limitCount {
		addBlankList(ip)
		return false
	}
	return true
}

// 返回ip在限制时间内访问的总次数
func Sum(ip string) int {
	ipLimitInfos.locker.RLock()
	defer ipLimitInfos.locker.RUnlock()
	num := 0
	if ips, ok := ipLimitInfos.ips[ip]; ok {
		if ips.headNode == nil {
			panic("can not be nil")
		}
		num += ips.headNode.num
		first := ips.headNode
		for current := ips.headNode.next; current != first; current = current.next {
			num += current.num
		}
	}
	return num
}

// 更新ip访问列表
// 查看当前时间与ip的第一个节点的访问时间间隔多远
func Update(ip string) {
	ipLimitInfos.locker.Lock()
	defer ipLimitInfos.locker.Unlock()
	t := time.Now().UnixNano()/1e6
	subNodeDur := limitDuration / nodeCount // 每个节点的时间间隔
	// 如果ip是第一次访问，则初始化
	if ips, ok := ipLimitInfos.ips[ip]; !ok || ips.headNode == nil {
		ipLimitInfos.ips[ip] = NewNodeList(t)
		ipLimitInfos.ips[ip].headNode.num++
		return
	}
	// 判断现在的访问时间和首节点的访问时间是否相隔
	ipInfo := ipLimitInfos.ips[ip].headNode
	durHeadNum := (t - ipLimitInfos.ips[ip].headTime) / int64(subNodeDur)
	// 如果当前时间和head节点的时间间隔在限制的时间范围内，则直接在对应的节点 的访问数量+1
	if durHeadNum < nodeCount {
		for i := 0; i < int(durHeadNum); i++ {
			ipInfo = ipInfo.next
		}
		ipInfo.num++
		return
	}
	// 如果当前时间和head节点的时间间隔大于两倍总的限制时间，则清空所有的节点，并设置新的首节点时间，并将最后一个节点的num设为1
	if durHeadNum >= 2*nodeCount-1 {
		for i := 0; i < nodeCount; i++ {
			ipInfo.num = 0
			ipInfo = ipInfo.next
		}
		ipInfo.num = 1
		ipLimitInfos.ips[ip].headTime = t
		return
	}
	// 如果当前时间和head节点的时间间隔在限制的时间的一到两倍之间，则需要判断当前的“head节点”，并将“过期”的节点重置并重新设置head节点
	expireNum := durHeadNum - (nodeCount - 1) // 计算有多少个节点“过期”了
	for i := 0; i < nodeCount; i++ {
		if i < int(expireNum)-1 {
			ipInfo.num = 0
		} else if i == int(expireNum)-1 {
			ipInfo.num = 1
		} else if i == int(expireNum) {
			ipLimitInfos.ips[ip].headNode = ipInfo
			break
		}
		ipInfo = ipInfo.next
	}
	ipLimitInfos.ips[ip].headTime = ipLimitInfos.ips[ip].headTime + expireNum*int64(subNodeDur)
}
