/*
通过锁来安全的通信
 */
package mutex

import (
	"fmt"
	"sync"
)

// 节点
type Peer struct {
	ID string
}

// 给节点发送消息
func (p *Peer) WriteMsg(msg string) {
	fmt.Printf("send to %s : %s \n", p.ID, msg)
}

// 当前节点的管理
type Host struct {
	peers map[string]*Peer // 根据peer id索引
	lock sync.RWMutex
}

func NewHost() *Host {
	h := &Host{
		peers: make(map[string]*Peer),
	}
	return h
}

func (h *Host) GetPeer(pid string) *Peer {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return h.peers[pid]
}

func (h *Host) AddPeer(p *Peer) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.peers[p.ID] = p
}

func (h *Host) RemovePeer(pid string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.peers, pid)
}

func (h *Host) BroadcastMsg(msg string) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	for _, v := range h.peers {
		v.WriteMsg(msg)
	}
}