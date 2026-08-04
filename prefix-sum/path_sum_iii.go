package prefixsum

// https://leetcode.com/problems/path-sum-iii/
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 用前缀和，每条路径是一个单独的解空间不干扰
	// 前缀和是为了快速求“区间和”
	prefixSumMap := map[int]int{0: 1}

	// 回溯算法
	var dfs func(p *TreeNode, currSum int) int
	dfs = func(p *TreeNode, currSum int) int {
		if p == nil {
			return 0
		}

		// try next state
		currSum += p.Val
		cnt := prefixSumMap[currSum-targetSum] // 利用前缀和来获取之前解的个数
		prefixSumMap[currSum]++

		cnt += dfs(p.Left, currSum)
		cnt += dfs(p.Right, currSum)

		// restore
		prefixSumMap[currMap]--

		return cnt
	}

	return dfs(root, 0)
}
