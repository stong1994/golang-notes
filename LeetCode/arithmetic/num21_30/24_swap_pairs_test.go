package num21_30

import (
	"fmt"
	"testing"
)

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
func TestSwapPairs(t *testing.T) {
	head := &ListNode{
		Val:1,
		Next: &ListNode{
			Val:2,
			Next:&ListNode{
				Val:3,
				Next:&ListNode{
					Val:4,
					Next:nil,
				},
			},
		},
	}
	head = swapPairs(head)
	for head.Next != nil {
		fmt.Println(head.Val)
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 一般思路,将节点放到数组中,然后遍历数组,更改节点位置,再遍历数组,重新排列节点(遍历数组可用分治)
//func swapPairs(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	var arr []*ListNode
//	for head != nil {
//		arr = append(arr, head)
//		head = head.Next
//	}
//	curIdx := 0
//	for i := 0; i < len(arr)-4; i+=4 {
//		arr[i], arr[i+1] = arr[i+1], arr[i]
//		arr[i+2], arr[i+3] = arr[i+3], arr[i+2]
//		curIdx+=4
//	}
//	if curIdx+1 < len(arr) {
//		arr[curIdx], arr[curIdx+1] = arr[curIdx+1], arr[curIdx]
//		if curIdx+3 < len(arr) {
//			arr[curIdx+2], arr[curIdx+3] = arr[curIdx+3], arr[curIdx+2]
//		}
//	}
//
//	head = arr[len(arr)-1]
//	head.Next = nil
//	for i:=len(arr)-2; i >=0; i-- {
//		arr[i].Next = head
//		head = arr[i]
//	}
//	return head
//}
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	if next.Next != nil {
		head.Next = swapPairs(next.Next)
	}else {
		head.Next = nil
	}
	next.Next = head
	return next
}