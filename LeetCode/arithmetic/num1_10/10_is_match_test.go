package num_1_10

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			"match0",
			"abc",
			"a.c",
			true,
		},
		{
			"match1",
			"aa",
			"a",
			false,
		}, {
			"match2",
			"aa",
			"a*",
			true,
		}, {
			"match3",
			"ab",
			".*",
			true,
		},
		{
			"match4",
			"aab",
			"c*a*b",
			true,
		},
		{
			"match5",
			"mississippi",
			"mis*is*p*.",
			false,
		},
		{
			"match6",
			"",
			"c*",
			true,
		},
		{
			"math6.5",
			"baabbbaccbccacacc",
			"c*..b*a*a.*a..*c",
			true,
		},
		{
			"match7",
			"mississippi",
			"mis*is*ip*.",
			true,
		},
		{
			"match8",
			"aaa",
			"a*a",
			true,
		},
		{
			"match9",
			"aaa",
			"ab*ac*a",
			true,
		},
		{
			"match10",
			"aaa",
			"ab*a*c*a",
			true,
		},
		{
			"match11",
			"a",
			"ab*",
			true,
		},
		{
			"match12",
			"bbbba",
			".*a*a",
			true,
		},
		{
			"match13",
			"ab",
			".*..",
			true,
		},
		{
			"match14",
			"a",
			".*..a*",
			false,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := isMatch(v.s, v.p)
			if got != v.want {
				t.Errorf("want %v but got %v", v.want, got)
			}
		})
	}
}

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
*/
// https://leetcode-cn.com/problems/regular-expression-matching/solution/yi-bu-dao-wei-zhi-jie-an-zheng-ze-biao-da-shi-de-s/
func debug(v ...interface{}) {
	log.Println(v...)
}

func toString(i interface{}) string {
	switch i.(type) {
	case int:
		return fmt.Sprintf("%v", i)
	case string:
		return fmt.Sprintf("%v", i)
	case bool:
		return fmt.Sprintf("%v", i)
	default:
		return fmt.Sprintf("%p", i)
	}
}

func isMatch(s string, p string) bool {
	begin := new(Node)
	begin.C = '>'
	begin.Size = generatePattern(begin, p, 0) // 设置有限状态机的长度
	debug(begin.String())
	return check(begin, s, 0)
}

type Node struct {
	C        byte             // 当前字符
	Parent   *Node            // 父节点
	Children map[byte][]*Node // 子孙节点
	End      bool             // 本节点是否为最终节点
	Size     int              // 本节点的最小长度，如果后面携带 *，则自由长度为 0; 否则为 1
}

func (n *Node) String() string {
	return n.StringLevel(0, make(map[*Node]bool))
}

func (n *Node) StringLevel(level int, finishNodes map[*Node]bool) string {
	r := make([]string, 0)
	if n.End {
		r = append(r, fmt.Sprintf("  id%s{%v};", toString(n), string(n.C)))
	} else {
		r = append(r, fmt.Sprintf("  id%s(%v);", toString(n), string(n.C)))
	}
	finishNodes[n] = true
	for k, v := range n.Children {
		for _, c := range v {
			if _, ok := finishNodes[c]; !ok {
				r = append(r, c.StringLevel(level+1, finishNodes))
			}
			r = append(r, fmt.Sprintf("  id%s -- %s --> id%s;", toString(n), string(k), toString(c)))
		}
	}
	return strings.Join(r, "\n")
}

// 判断n的子节点中是否有c以及child，如果没有，添加
func (n *Node) Append(c byte, child *Node) {
	m := n.Children
	if m == nil {
		m = make(map[byte][]*Node)
		n.Children = m
	}
	list := m[c]
	if list == nil {
		list = make([]*Node, 0)
	}
	for _, v := range list {
		if v == child {
			// m[c] = list
			return
		}
	}
	list = append(list, child)
	m[c] = list
}

// now为str[idx]的父节点，str为正则字符串，idx为当前字符串索引，返回状态机中当前节点后的的长度
func generatePattern(now *Node, str string, idx int) int {
	if len(str) <= idx {
		now.End = true
		return now.Size
	}
	vnow := now
	switch str[idx] {
	case '*': // *代表零个或多个now.C，指定size为0并在now节点的子孙节点中增加本身
		now.Size = 0
		now.Append(now.C, now)
	default: // 为now的子孙节点中添加一个普通节点
		node := new(Node)
		node.C = str[idx]
		now.Append(str[idx], node)
		node.Parent = now
		node.Size = 1
		vnow = node
	}
	ret := generatePattern(vnow, str, idx+1) // 递归剩余的字符
	if ret == 0 {
		now.End = true
	}
	addParent := now
	for addParent.Parent != nil { // 如果父节点为*，那么增加当前节点为爷节点的子孙节点
		if addParent.Size == 0 {
			debug((vnow), " -> ", toString(addParent.Parent))
			addParent.Parent.Append(vnow.C, vnow)
			addParent = addParent.Parent
		}else {
			break
		}
	}
	return now.Size + ret
}

// 检测str是否符合有限状态机
func check(now *Node, str string, idx int) bool {
	if len(str) <= idx {
		return now.End
	}
	list := now.Children['.'] // 如果子孙节点中包含. 那么在for循环的idx+1,相当于匹配到了任何数
	for _, v := range now.Children[str[idx]] { // 获取
		list = append(list, v)
	}
	for _, v := range list {
		r := check(v, str, idx+1)
		if r {
			return true
		}
	}
	return false
}
