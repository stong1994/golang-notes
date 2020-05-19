package num101_110

import "testing"

func TestMaxDepth(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	res := maxDepth(tree)
	if res != 3 {
		t.Errorf("want %d got %d", 3, res)
	}
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left+1
	}
	return right + 1
}