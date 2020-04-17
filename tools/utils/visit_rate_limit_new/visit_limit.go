package limiter

import (
	"sync"
	"time"
)

/*
思路：
	设置一个10个节点构成的环形链表来存放ip访问的次数。假如限制为10s内访问100次，那么每个节点间的间隔为1s，每次有ip访问，就将ip放入对应时间的节点，并将节点的访问次数加1，判断整个链表的
访问次数是否超过了100次，如果超过了，就将ip放入黑名单。如果有节点的访问时间超过了10s，那么当用户访问时，更新这些过期的节点
*/

/*
	CheckIP(ip string) bool    // 检查ip是否在黑名单
	Update(ip string, t int64) // 更新ip的环形链表
	Sum(ip string) int         // 查看限制时间内的总访问量
*/

type visitLimit struct {
	limitCount int // 限制规则的访问次数
	limitDuration int // 限制规则的时间间隔 单位：毫秒
	nodeCount int
	blankList *blankListModel
	ipLimit *ipLimit
}

func NewVisitLimit(limitCount, limitDuration, nodeCount int) *visitLimit {
	return &visitLimit{
		limitCount:limitCount,
		limitDuration:limitDuration,
		nodeCount:nodeCount,
		blankList:newBlockList(),
		ipLimit: newIpLimit(),
	}
}

// 黑名单
type blankListModel struct {
	list   map[string]int64 // key => 用户表示，例如：ip， value => 访问的开始时间
	sync.RWMutex
}

func newBlockList() *blankListModel {
	return &blankListModel{list: make(map[string]int64)}
}

// 没有在黑名单，返回false， 在黑名单返回true
func (b *blankListModel) CheckInBlankList(ip string) bool {
	b.RLock()
	defer b.RUnlock()
	if _, ok := b.list[ip]; ok {
		return true
	}
	return false
}

func (b *blankListModel) AddBlankList(ip string) {
	b.Lock()
	defer b.Unlock()
	b.list[ip] = time.Now().Unix()
}

// 访问的ip信息
type ipLimit struct {
	ips    map[string]*NodeList
	sync.RWMutex
}

func newIpLimit() *ipLimit {
	return &ipLimit{
		ips:     make(map[string]*NodeList, 10000),
	}
}

func (v *visitLimit) CheckIP(ip string) bool {
	if v.blankList.CheckInBlankList(ip) {
		return false
	}
	total := v.VisitSum(ip)
	if total > v.limitCount {
		v.blankList.AddBlankList(ip)
		return false
	}
	return true
}

// 返回ip在限制时间内访问的总次数
func (v *visitLimit) VisitSum(ip string) int {
	v.ipLimit.RLock()
	defer v.ipLimit.RUnlock()
	num := 0
	if ips, ok := v.ipLimit.ips[ip]; ok {
		if ips.headNode == nil {
			panic("head node can not be nil")
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
func (v *visitLimit) UpdateIp(ip string) {
	v.ipLimit.Lock()
	defer v.ipLimit.Unlock()
	t := time.Now().UnixNano()/1e6
	subNodeDur := v.limitDuration / v.nodeCount // 每个节点的时间间隔
	// 如果ip是第一次访问，则初始化
	if ips, ok := v.ipLimit.ips[ip]; !ok || ips.headNode == nil {
		v.ipLimit.ips[ip] = NewNodeList(v.nodeCount)
		v.ipLimit.ips[ip].headNode.num++
		return
	}
	// 判断现在的访问时间和首节点的访问时间是否相隔
	ipInfo :=v.ipLimit.ips[ip].headNode
	durHeadNum := (t - v.ipLimit.ips[ip].headTime) / int64(subNodeDur)
	// 如果当前时间和head节点的时间间隔在限制的时间范围内，则直接在对应的节点 的访问数量+1
	if durHeadNum < int64(v.nodeCount) {
		for i := 0; i < int(durHeadNum); i++ {
			ipInfo = ipInfo.next
		}
		ipInfo.num++
		return
	}
	// 如果当前时间和head节点的时间间隔大于两倍总的限制时间，则清空所有的节点，并设置新的首节点时间，并将最后一个节点的num设为1
	if durHeadNum >= 2*int64(v.nodeCount)-1 {
		for i := 0; i < v.nodeCount; i++ {
			ipInfo.num = 0
			ipInfo = ipInfo.next
		}
		ipInfo.num = 1
		v.ipLimit.ips[ip].headTime = t
		return
	}
	// 如果当前时间和head节点的时间间隔在限制的时间的一到两倍之间，则需要判断当前的“head节点”，并将“过期”的节点重置并重新设置head节点
	expireNum := durHeadNum - int64(v.nodeCount - 1) // 计算有多少个节点“过期”了
	for i := 0; i < v.nodeCount; i++ {
		if i < int(expireNum)-1 {
			ipInfo.num = 0
		} else if i == int(expireNum)-1 {
			ipInfo.num = 1
		} else if i == int(expireNum) {
			v.ipLimit.ips[ip].headNode = ipInfo
			break
		}
		ipInfo = ipInfo.next
	}
	v.ipLimit.ips[ip].headTime = v.ipLimit.ips[ip].headTime + expireNum*int64(subNodeDur)
}
