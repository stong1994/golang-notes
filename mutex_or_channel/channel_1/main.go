// sharing memory by  communicating
//  don't communicate by sharing memory

// 缺点，没有实现get方法，因为有返回值需要接收
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

// 当前节点的管理
type Host struct {
	add chan *Peer
	broadcast chan string
	remove chan string
	stop chan struct{}
}

func NewHost() *Host {
	return &Host{
		add: make(chan *Peer),
		broadcast: make(chan string),
		remove: make(chan string),
		stop: make(chan struct{}),
	}
}

func (h *Host) Start() {
	go h.loop()
}

func (h *Host) Stop() {
	close(h.stop)
}

func (h *Host) AddPeer(p *Peer) {
	h.add <- p
}

func (h *Host) RemovePeer(pid string) {
	h.remove <- pid
}

func (h *Host) BroadcastMsg(msg string) {
	h.broadcast <- msg
}

func (h *Host) loop() {
	peers := make(map[string]*Peer)

	for {
		select {
		case p:= <- h.add:
			peers[p.ID] = p
		case pid := <- h.remove:
			delete(peers, pid)
		case msg := <- h.broadcast:
			for _, p := range peers {
				p.WriteMsg(msg)
			}
		case <- h.stop:
			h.Stop()
		}
	}
}