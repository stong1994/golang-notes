package num141_150

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val:1,
		Right:&TreeNode{
			Val:2,
			Left: &TreeNode{
				Val:3,
			},
		},
	}
	result := preorderTraversal2(root)
	want := []int{1,2,3}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("want %v got %v", want, result)
	}
}

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

// 给定一个二叉树，返回它的 前序 遍历。
// 递归
func preorderTraversal(root *TreeNode) []int {
	var arr []int

	if root == nil {
		return arr
	}
	arr = append(arr, root.Val)
	left := preorderTraversal(root.Left)
	arr = append(arr, left...)
	right := preorderTraversal(root.Right)
	arr = append(arr, right...)
	return arr
}

// 迭代算法
// 栈用来保存待处理的右节点，先处理左节点
func preorderTraversal2(root *TreeNode) []int {
	var arr []int
	var stack []*TreeNode
	if root == nil {
		return arr
	}

	for {
		// 将当前节点的右节点都放到栈中
		for root != nil {
			arr = append(arr, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}
		// 出栈
		index := len(stack) - 1
		if index < 0 {
			break
		}
		root = stack[index]
		stack = stack[:index]
	}
	return arr
}