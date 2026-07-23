package leetcode

// dp[i,j] = dp[i-1][j] + dp[i][j-1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rowSize, colSize := len(obstacleGrid), len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 || obstacleGrid[rowSize-1][colSize-1] == 1 {
		return 0
	} else if rowSize == 1 && colSize == 1 {
		return 1
	}

	dp := make([]int, colSize)
	dp[0] = 1
	for j := 1; j < colSize; j++ {
		if obstacleGrid[0][j] == 1 {
			dp[j] = 0
		} else {
			dp[j] = dp[j-1] & 1
		}
	}

	for i := 1; i < rowSize; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[0] = 0
		} else {
			dp[0] = dp[0] & 1
		}

		for j := 1; j < colSize; j++ {
			if obstacleGrid[i][j] == 1 {
				// obstacle
				dp[j] = 0
			} else if dp[j] == 0 && dp[j-1] == 0 {
				// no path
				dp[j] = 0
			} else {
				dp[j] = dp[j-1] + dp[j]
			}
		}
	}

	return dp[len(dp)-1]
}
