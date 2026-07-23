package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func NewTreeFromString(s string) *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	return root
}

func recoverTreeOfString(s string) string {
	root := NewTreeFromString(s)
	recoverTree(root)
	return inorderTraverseTreeToString(root)
}

// 1, 6, 3, 4, 5, 2
// first find large , than find small, use in-order visiting to make traverse asending
func recoverTree(root *TreeNode) {
	// declare variables, 注意这里不能将下面的变量通过参数传递，实际上是传指针的值，而非引用
	var (
		pre   *TreeNode
		large *TreeNode
		small *TreeNode
	)

	// use inner functions so we can use vars above as global variable in method recoverTree
	var _inorderTraverse func(p *TreeNode)
	_inorderTraverse = func(p *TreeNode) {
		if p == nil {
			return
		}

		_inorderTraverse(p.Left)

		if pre != nil && pre.Val > p.Val {
			if large == nil {
				// needed when reversed nodes are adjacent to each other
				large = pre
				small = p
			} else {
				small = p
				return
			}
		}

		if p != nil {
			pre = p
		}

		_inorderTraverse(p.Right)
	}

	// found not valid nodes
	_inorderTraverse(root)
	// swap off
	large.Val, small.Val = small.Val, large.Val
}

func inorderTraverseTreeToString(root *TreeNode) string {
	return ""
}
