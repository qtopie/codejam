package leetcode

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)

	var backtracking func(p *TreeNode, states *[]int, sum int, target int)
	backtracking = func(p *TreeNode, states *[]int, sum int, target int) {
		if p == nil {
			return
		}

		*states = append(*states, p.Val)
		sum += p.Val

		if p.Left == nil && p.Right == nil && sum == target {
			paths := make([]int, len(*states))
			copy(paths, *states)
			result = append(result, paths)
		} else {
			if p.Left != nil {
				backtracking(p.Left, states, sum, target)
			}

			if p.Right != nil {
				backtracking(p.Right, states, sum, target)
			}
		}

		sum -= p.Val
		*states = (*states)[:len(*states)-1]
	}

	states := []int{}
	backtracking(root, &states, 0, targetSum)
	return result
}
