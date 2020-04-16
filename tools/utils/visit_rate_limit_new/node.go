package visit_limit

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

func NewNodeList(nodeCount int) *NodeList {
	node := new(Node) // 节点链表中最后一个节点
	end := node
	// 已初始化最后一个节点，从倒数第二个节点开始，逆序不断得到新的节点。
	for i := nodeCount - 1; i > 0; i-- {
		node = NewNode(node)
	}
	end.next = node // 遍历完后的node为节点链表中的第一个节点
	return &NodeList{headNode: node, headTime: time.Now().UnixNano()/1e6}
}

func (n *NodeList) init() {
	n.headTime = time.Now().UnixNano() / 1e6
	n.headNode.num = 1
}
