package num21_30

import "testing"

func TestReverseKGroup(t *testing.T) {
	head := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:3,
				Next:&ListNode{
					Val:4,
					Next:&ListNode{
						Val:5,
						Next:nil,
					},
				},
			},
		},
	}
	reverseKGroup(head, 3)
}

// 一般解法,将节点放入数组中,然后更换位置,再重新生成链表
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	var arr []*ListNode
//	for head != nil {
//		arr = append(arr, head)
//		head = head.Next
//	}
//	arrLen := len(arr)
//	for i := 0; i <= arrLen-k; i+=k {
//		reverseArr(arr, i, i+k-1)
//	}
//	head = arr[arrLen-1]
//	head.Next = nil
//	for i := len(arr)-2; i>=0; i-- {
//		arr[i].Next = head
//		head = arr[i]
//	}
//	return head
//}

func reverseArr(arr []*ListNode, m, n int)  {
	if m >= n {
		return
	}
	arr[m], arr[n] = arr[n], arr[m]
	reverseArr(arr, m+1, n-1)
}


func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nex := head.Next
	for i := 0; i < k-1; i++ {
		if nex == nil {
			return head
		}
		nex = nex.Next
	}

	last := head
	next := last.Next
	cur := next
	for i := 0; i < k-1; i++ {
		cur = next.Next
		next.Next = last
		last = next
		next = cur
	}

	head.Next = reverseKGroup(next, k)
	return last
}