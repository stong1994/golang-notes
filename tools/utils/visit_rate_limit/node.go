package main

import "time"

// 一个节点，存放一段时间内ip访问的次数
type Node struct {
	next *Node
	num  int
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
	for i := nodeCount - 1; i >= 0; i-- {
		node = NewNode(node)
	}
	head := node
	// 最后一个节点的next为空node，现在将它赋值为head todo 更简单的方法？
	for {
		if node.next.next == nil {
			node.next = head
			break
		}
		node = node.next
	}
	return &NodeList{headNode: head, headTime: now}
}

func (n *NodeList) init() {
	n.headTime = time.Now().UnixNano() / 1e6
	n.headNode.num = 1
}
