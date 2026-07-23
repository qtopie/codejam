package leetcode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// check root
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil || q == nil && p != nil || p.Val != q.Val {
		return false
	}

	// check left and child
	if !isSameTree(p.Left, q.Left) || !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
