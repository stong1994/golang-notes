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

func TestForI(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		i++
		fmt.Println(arr[i])
	}
}

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
*/
// todo https://leetcode-cn.com/problems/regular-expression-matching/solution/yi-bu-dao-wei-zhi-jie-an-zheng-ze-biao-da-shi-de-s/
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
	begin.Size = generatePattern(begin, p, 0)
	debug(begin.String())
	return check(begin, s, 0)
}

type Node struct {
	C        byte
	Parent   *Node
	Children map[byte][]*Node
	End      bool
	Size     int
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
			m[c] = list
			return
		}
	}
	list = append(list, child)
	m[c] = list
}

func generatePattern(now *Node, str string, idx int) int {
	if len(str) <= idx {
		now.End = true
		return now.Size
	}
	vnow := now
	switch str[idx] {
	case '*':
		now.Size = 0
		now.Append(now.C, now)
	default:
		node := new(Node)
		node.C = str[idx]
		now.Append(str[idx], node)
		node.Parent = now
		node.Size = 1
		vnow = node
	}
	ret := generatePattern(vnow, str, idx+1)
	if ret == 0 {
		now.End = true
	}
	addParent := now
	for addParent.Parent != nil {
		if addParent.Size == 0 {
			debug(toString(vnow), " -> ", toString(addParent.Parent))
			addParent.Parent.Append(vnow.C, vnow)
			addParent = addParent.Parent
		} else {
			break
		}
	}
	return now.Size + ret
}

func check(now *Node, str string, idx int) bool {
	if len(str) <= idx {
		return now.End
	}
	list := now.Children['.']
	for _, v := range now.Children[str[idx]] {
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
