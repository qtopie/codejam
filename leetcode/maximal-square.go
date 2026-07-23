package leetcode

// https://leetcode.com/problems/maximal-square/description/?envType=study-plan-v2&envId=top-interview-150
func maximalSquare(matrix [][]byte) int {
	maxLen := 0
	row, col := len(matrix), len(matrix[0])

	dp := make([]int, col+1)

	for i := 0; i < row; i++ {
		for tmp, j := dp[0], 0; j < col; j++ {
			if matrix[i][j] == '1' {
				tmp, dp[j+1] = dp[j+1], 1+min(min(dp[j], tmp), dp[j+1])
				maxLen = max(maxLen, dp[j+1])
			} else {
				dp[j+1] = 0
			}
		}

	}

	return maxLen * maxLen
}
