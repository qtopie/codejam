package leetcode

func minPathSum(grid [][]int) int {
	rowSize, colSize := len(grid), len(grid[0])
	if rowSize == 1 && colSize == 1 {
		return grid[0][0]
	}

	dp := make([]int, colSize)
	dp[0] = grid[0][0]
	for j := 1; j < colSize; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	for i := 1; i < rowSize; i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < colSize; j++ {
			dp[j] = grid[i][j] + min(dp[j-1], dp[j])
		}
	}

	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
