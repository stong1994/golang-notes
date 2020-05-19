package num91_100

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tree := &TreeNode{
		Val:   10,
		Left:  &TreeNode{
			Val:   5,
		},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{
				Val:   6,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	
	res := isValidBST(tree)
	if res {
		t.Errorf("want false got true")
	}
}

/*给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, -1<<63, 1<<63-1)
}

func dfs(root *TreeNode, min, max int) bool {
	return root == nil || min < root.Val && root.Val < max &&
		dfs(root.Left, min, root.Val) &&
		dfs(root.Right, root.Val, max)
}
