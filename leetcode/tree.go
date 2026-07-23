package leetcode

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (p *TreeNode) ToArray() []string {
	return toArray(p)
}

func toArray(root *TreeNode) []string {
	if root == nil {
		return []string{"null"}
	}

	result := []string{strconv.Itoa(root.Val)}

	if root.Left == nil && root.Right == nil {
		return result
	}

	left := toArray(root.Left)
	right := toArray(root.Right)
	result = append(result, left...)
	result = append(result, right...)

	return result
}

func clone(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	newRoot := &TreeNode{Val: root.Val}
	newRoot.Left = clone(root.Left)
	newRoot.Right = clone(root.Right)
	return newRoot
}

func preorderTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)

	preorderTraverse(root.Left)
	preorderTraverse(root.Right)
}

func inorderTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	inorderTraverse(root.Left)

	fmt.Println(root.Val)

	inorderTraverse(root.Right)
}
