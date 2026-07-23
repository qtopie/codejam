package main

// Category: algorithms
// Level: Hard
// Percent: 42.082424%

// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.
//
// The path sum of a path is the sum of the node's values in the path.
//
// Given the root of a binary tree, return the maximum path sum of any non-empty path.
//
//
// Example 1:
//
// Input: root = [1,2,3]
// Output: 6
// Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
//
//
// Example 2:
//
// Input: root = [-10,9,20,null,null,15,7]
// Output: 42
// Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
//
//
//
// Constraints:
//
//
// 	The number of nodes in the tree is in the range [1, 3 * 10⁴].
// 	-1000 <= Node.val <= 1000
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	maxSum := root.Val

	var getGain func(*TreeNode) int
	getGain = func(p *TreeNode) int {
		if p == nil {
			return 0
		}

		// max自动减支负贡献
		leftGain := max(0, getGain(p.Left))
		rightGain := max(0, getGain(p.Right))

		// 合计贡献， 注意如果当前节点整体是负贡献，则会忽略，因为我们保存了全局最优解
		currentMax := p.Val + leftGain + rightGain
		maxSum = max(maxSum, currentMax)

		// 计算单个节点某条边的最大贡献
		return p.Val + max(leftGain, rightGain)
	}

	getGain(root)
	return maxSum
}
