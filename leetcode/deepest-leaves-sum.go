package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	sum := 0
	maxDepth := 0

	_dfs_leaves(&sum, &maxDepth, root, 1)
	return sum
}

func _dfs_leaves(sum *int, maxDepth *int, p *TreeNode, depth int) {
	if p == nil {
		return
	}

	if depth > *maxDepth {
		*sum = 0
		*maxDepth = depth
	}

	if p.Left == nil && p.Right == nil {
		if depth >= *maxDepth {
			*sum += p.Val
		}
	}

	if p.Left != nil {
		_dfs_leaves(sum, maxDepth, p.Left, depth+1)
	}

	if p.Right != nil {
		_dfs_leaves(sum, maxDepth, p.Right, depth+1)
	}
}
