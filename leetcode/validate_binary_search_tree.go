package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return _isValidBST(root, -1<<31, 1<<31-1)
}

func _isValidBST(root *TreeNode, start, end int) bool {
	if root == nil {
		return true
	}

	if root.Val < start || root.Val > end {
		return false
	}

	return _isValidBST(root.Left, start, root.Val-1) && _isValidBST(root.Right, root.Val+1, end)
}
