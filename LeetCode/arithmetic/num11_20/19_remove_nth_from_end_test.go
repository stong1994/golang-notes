package main

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head1 := &ListNode{1, nil}
	now1 := head1
	for i := 0; i < 1; i++ {
		now1.Next = &ListNode{i + 2, nil}
		now1 = now1.Next
	}
	tests := []struct {
		name string
		head *ListNode
		n    int
	}{
		{
			"test1",
			head1,
			2,
		},
	}
	now := head1
	for {
		fmt.Printf("%+v\n", now)
		if now.Next == nil {
			break
		}
		now = now.Next
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := removeNthFromEnd(v.head, v.n)
			gotNow := got
			for {
				fmt.Printf("%+v\n", gotNow)
				if gotNow.Next == nil {
					break
				}
				gotNow = gotNow.Next
			}
		})
	}
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

// 删除倒数第n个节点
// 双指针法（https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/solution/dong-hua-tu-jie-leetcode-di-19-hao-wen-ti-shan-chu/）
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeP, nodeQ := new(ListNode), new(ListNode)
	nodeP.Next = head // 有些需要删除第一个节点，因此在第一个节点前设置一个空节点
	nodeQ.Next = head
	first := nodeP
	for i := 0; i < n; i++ { // p和q本来是同步的，让q领先p n个节点，那么当q到达尾部时，p的下一个节点就是要删除的节点
		nodeQ = nodeQ.Next
	}
	for {
		if nodeQ.Next == nil {
			nodeP.Next = nodeP.Next.Next
			break
		}
		nodeQ = nodeQ.Next
		nodeP = nodeP.Next
	}

	return first.Next
}
