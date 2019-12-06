package num21_30

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {
	v1 := []int{1, 4, 5}
	v2 := []int{1, 3, 4}
	v3 := []int{2, 6}
	mergeV := []int{1, 1, 2, 3, 4, 4, 5, 6}
	var (
		l1 *ListNode
		l2 *ListNode
		l3 *ListNode
	)
	for i := len(v1) - 1; i >= 0; i-- {
		if l1 == nil {
			l1 = &ListNode{
				Val:  v1[i],
				Next: nil,
			}
			continue
		}
		l1 = &ListNode{
			Val:  v1[i],
			Next: l1,
		}
	}
	for i := len(v2) - 1; i >= 0; i-- {
		if l2 == nil {
			l2 = &ListNode{
				Val:  v2[i],
				Next: nil,
			}
			continue
		}
		l2 = &ListNode{
			Val:  v2[i],
			Next: l2,
		}
	}

	for i := len(v3) - 1; i >= 0; i-- {
		if l3 == nil {
			l3 = &ListNode{
				Val:  v3[i],
				Next: nil,
			}
			continue
		}
		l3 = &ListNode{
			Val:  v3[i],
			Next: l3,
		}
	}

	nodeList := []*ListNode{l1, l2, l3}

	mergeL := mergeKLists(nodeList)

	for {
		if mergeL == nil {
			if len(mergeV) == 0 {
				t.Log("success")
				return
			}
			t.Error("result can not be nil")
			return
		}
		if len(mergeV) <= 0 {
			t.Error("result length over than mergeV")
			return
		}
		if mergeL.Val != mergeV[0] {
			t.Errorf("expect %d got %d", mergeV[0], mergeL.Val)
			return
		}
		mergeL = mergeL.Next
		mergeV = mergeV[1:]
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 递归 时间复杂度 O(m) m为链表中总的元素个数
// func mergeKLists(lists []*ListNode) *ListNode {
// 	if len(lists) == 0 {
// 		return nil
// 	}
// 	minIdx := -1
// 	minVal := math.MaxInt64
// 	for i, list := range lists {
// 		if list == nil {
// 			continue
// 		}
// 		if list.Val < minVal {
// 			minVal = list.Val
// 			minIdx = i
// 		}
// 	}
// 	if minIdx == - 1 {
// 		return nil
// 	}
// 	node := &ListNode{
// 		Val:  minVal,
// 		Next: nil,
// 	}
// 	lists[minIdx] = lists[minIdx].Next
// 	return Append(node, lists)
// }

// func Append(now *ListNode, lists []*ListNode) *ListNode {
// 	now.Next = mergeKLists(lists)
// 	return now
// }

// 迭代
// func mergeKLists(lists []*ListNode) *ListNode {
// 	if len(lists) == 0 {
// 		return nil
// 	}
// 	preNode := &ListNode{}
// 	nodeList := preNode
//
// 	ended := func() bool {
// 		for _, v := range lists {
// 			if v != nil {
// 				return false
// 			}
// 		}
// 		return true
// 	}
//
// 	for {
// 		if ended() {
// 			break
// 		}
// 		minIdx := -1
// 		minVal := math.MaxInt64
// 		for i, list := range lists {
// 			if list == nil {
// 				continue
// 			}
// 			if list.Val < minVal {
// 				minVal = list.Val
// 				minIdx = i
// 			}
// 		}
// 		if minIdx == - 1 {
// 			break
// 		}
// 		node := &ListNode{
// 			Val:  minVal,
// 			Next: nil,
// 		}
// 		nodeList.Next = node
// 		nodeList = nodeList.Next
// 		lists[minIdx] = lists[minIdx].Next
// 	}
//
// 	return preNode.Next
// }

// 分治 (时间复杂度，两个链表合并为O(min(m,n))（递归）,将链表数组分治O(log(k))(分治). 相乘。空间复杂度为O(z)(z为链表中总的元素个数))
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) == 2 {
		return mergeTwoList(lists)
	}

	left := mergeKLists(lists[:len(lists)/2])
	right := mergeKLists(lists[len(lists)/2:])
	return mergeKLists([]*ListNode{left, right})
}

func mergeTwoList(list []*ListNode) *ListNode {
	if len(list) != 2 {
		panic("length must be 2")
	}
	if list[0] == nil {
		return list[1]
	}
	if list[1] == nil {
		return list[0]
	}
	var node *ListNode
	if list[0].Val <= list[1].Val {
		node = list[0]
		list[0] = list[0].Next
		node.Next = mergeTwoList(list)
	} else {
		node = list[1]
		list[1] = list[1].Next
		node.Next = mergeTwoList(list)
	}
	return node
}
