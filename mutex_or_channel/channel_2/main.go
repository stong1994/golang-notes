// sharing memory by  communicating
//  don't communicate by sharing memory

// 实现了Get()方法
// 缺点：不能同时读
package main

import (
	"fmt"
)

// 节点
type Peer struct {
	ID string
}

// 给节点发送消息
func (p *Peer) WriteMsg(msg string) {
	fmt.Printf("send to %s : %s \n", p.ID, msg)
}


type Operation func(peers map[string]*Peer)
// 当前节点的管理
type Host struct {
	opCh chan Operation
	stop chan struct{}
}

func NewHost() *Host {
	return &Host{
		opCh: make(chan Operation),
		stop: make(chan struct{}),
	}
}

func (h *Host) Start() {
	go h.loop()
}

func (h *Host) Stop() {
	close(h.stop)
}

func (h *Host) GetPeer(pid string) *Peer {
	recevCh := make(chan *Peer)
	op := func(peers map[string]*Peer) {
		recevCh <- peers[pid]
	}

	// 传送数据
	go func() {
		h.opCh <- op
	}()

	// 等待接收数据
	return <- recevCh
}

func (h *Host) SendToOne(pid, str string) {
	p := h.GetPeer(pid)
	p.WriteMsg(str)
}


func (h *Host) AddPeer(p *Peer) {
	op := func(peers map[string]*Peer) {
		peers[p.ID] = p
	}
	h.opCh <- op
}

func (h *Host) RemovePeer(pid string) {
	op := func(peers map[string]*Peer) {
		delete(peers, pid)
	}
	h.opCh <- op
}

func (h *Host) BroadcastMsg(msg string) {
	op := func(peers map[string]*Peer) {
		for _, p := range peers {
			p.WriteMsg(msg)
		}
	}
	h.opCh <- op
}

func (h *Host) loop() {
	peers := make(map[string]*Peer)

	for {
		select {
		case op:= <- h.opCh:
			op(peers)
		case <- h.stop:
			h.Stop()
		}
	}
}