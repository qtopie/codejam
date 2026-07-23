package leetcode

func generateParenthesis(n int) []string {
	return _generateParenthesis(n-1, n, "(")
}

func _generateParenthesis(l, r int, p string) []string {
	result := make([]string, 0)

	if l > r {
		return result
	} else if l == 0 {
		// fill with right brackets
		for i := 1; i <= r; i++ {
			p = p + ")"
		}

		result = append(result, p)
		return result
	}

	// try generating with (
	arr1 := _generateParenthesis(l-1, r, p+"(")
	result = append(result, arr1...)
	// try generating with )
	arr2 := _generateParenthesis(l, r-1, p+")")
	result = append(result, arr2...)

	return result
}

// 使用动态规划解决, 利用二维数组存储已有解 X_n = A + B
// 当i=n-1时,当使用()包住A的时候, 都会产生一个新的A,
// 考虑到与B的组合, 可以一直追溯到0
func _generateParens(n int) []string {
	dp := make([][]string, n+1)

	dp[0] = append(dp[0], "")

	// for each pair of parens starting from 1 to n
	for i := 1; i <= n; i++ {
		// for each of existed results
		for j := 0; j < i; j++ {
			// new results
			parens := make([]string, 0)

			// for each left part a
			for _, a := range dp[j] {
				// for each right part b
				for _, b := range dp[i-1-j] {
					parens = append(parens, "("+a+")"+b)
				}
			}

			dp[i] = append(dp[i], parens...)
		}
	}

	return dp[n]
}
