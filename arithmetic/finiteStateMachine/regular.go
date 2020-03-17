package main

import (
	"fmt"
)

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
https://leetcode-cn.com/problems/regular-expression-matching/solution/yi-bu-dao-wei-zhi-jie-an-zheng-ze-biao-da-shi-de-s/
*/


type Node struct {
	ele      byte             // 该节点字符
	children map[byte][]*Node // 子孙节点
	end      bool             // 是否为最终节点
	size     int              // 本节点的最小长度（主要是*可为0）
	parent   *Node            // 父节点
}

// 将 c=>child 加入n的子孙节点, 如果c中已存在child，那么直接返回，如果不存在，那么将child 加入到c的节点中
func (n *Node) Append(c byte, child *Node) {
	m := n.children
	if m == nil {
		m = make(map[byte][]*Node)
		n.children = m
	}
	list := m[c]
	if list == nil {
		list = make([]*Node, 0)
	}
	for _, v := range list {
		if v == child {
			return
		}
	}
	list = append(list, child)
	m[c] = list
}

func debug(data ...interface{})  {
	fmt.Println(data...)
}

// s: 要匹配的字符串  p：正则字符串
func isMatch(s string, p string) bool {
	begin := new(Node)
	begin.ele = '>'
	begin.size = generatePattern(begin, p, 0)
	debug(begin)
	return check(begin, s, 0)
}

// now：当前要匹配的节点 str：要匹配的字符串 i:要匹配的字符串的索引
func check(now *Node, str string, i int) bool {
	if len(str) <= i {
		return now.end
	}
	list := now.children['.'] // 如果当前节点元素中存在'.'，那么说明可以任意匹配，那么list长度不为0，直接check下个字符
	for _, v := range now.children[str[i]] { // 获取要匹配的字符串第i个元素，查看当前节点是否含有此元素的子孙节点
		list = append(list, v)
	}
	for _, v := range list { // list长度不为0，说明能够匹配到数据，则索引+1，进行下个元素匹配
		r := check(v, str, i+1)
		if r {
			return true
		}
	}
	return false
}

// now:当前节点 str：正则字符串 i：当前正则字符串的索引， 返回now的size
func generatePattern(now *Node, str string, i int) int {
	if len(str) <= i {
		now.end = true
		return now.size
	}
	next := now
	switch str[i] {
	case '*':
		now.size = 0
		now.Append(now.ele, now) // *代表任意数量的now，用闭合环表示
	default:
		node := new(Node)
		node.ele = str[i]
		now.Append(str[i], node)
		node.parent = now
		node.size = 1
		next = node
	}
	ret := generatePattern(next, str, i+1)
	if ret == 0 {
		now.end = true
	}
	addParent := now
	for addParent.parent != nil { // 查看当前节点的父节点，如果父节点size为0，即当前元素为'*'，那么给父节点增加子节点next => 情景为 ***b , 那么这几个*都有子节点b
		if addParent.size == 0 {
			debug(next, " -> ", addParent.parent)
			addParent.parent.Append(next.ele, next)
			addParent = addParent.parent
		}else {
			break
		}
	}
	return now.size + ret
}


