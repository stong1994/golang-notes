package num21_30

import "testing"

/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func TestMergeTwoLists(t *testing.T) {
	v1 := []int{1, 2, 4}
	v2 := []int{1, 3, 4}
	mergeV := []int{1, 1, 2, 3, 4, 4}
	var (
		l1 *ListNode
		l2 *ListNode
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

	mergeL := mergeTwoLists(l1, l2)

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
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	mergeL := &ListNode{0, nil}
// 	start := mergeL
//
// 	for {
// 		if l1 == nil {
// 			mergeL.Next = l2
// 			break
// 		}
// 		if l2 == nil {
// 			mergeL.Next = l1
// 			break
// 		}
// 		if l1.Val <= l2.Val {
// 			mergeL.Next = l1
// 			mergeL = mergeL.Next
// 			l1 = l1.Next
// 			continue
// 		}
// 		mergeL.Next = l2
// 		l2 = l2.Next
// 		mergeL = mergeL.Next
// 	}
// 	return start.Next
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归 (递归比上边的迭代空间复杂度高，迭代只用到了4个指针，O(1);递归会遍历l1和l2，O(m+n))
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
