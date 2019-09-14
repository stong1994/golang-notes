package num_1_10

import (
	"reflect"
	"testing"
)

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

链接：https://leetcode-cn.com/problems/add-two-numbers
*/
func TestAddTwoNumbers(t *testing.T) {

	n1 := &ListNode{3, nil}
	n2 := &ListNode{4, n1}
	n3 := &ListNode{2, n2}

	n11 := &ListNode{4, nil}
	n12 := &ListNode{6, n11}
	n13 := &ListNode{5, n12}

	n21 := &ListNode{8, nil}
	n22 := &ListNode{0, n21}
	n23 := &ListNode{7, n22}

	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			"test1",
			n3,
			n13,
			n23,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := addTwoNumbers(v.l1, v.l2)
			if !judgeNodeEqual(got, v.want) {
				t.Errorf("want %+v got %+v", v.want, got)
			}
		})
	}
}

func judgeNodeEqual(n1, n2 *ListNode) bool {
	return reflect.DeepEqual(n1, n2)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	p, q, cur := l1, l2, dummyHead
	carry := 0
	for p != nil || q != nil {
		x, y := 0, 0
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		sum := x + y + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}
