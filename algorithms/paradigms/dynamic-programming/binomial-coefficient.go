package dynamic_programming

// 计算二项展开式系数
// 递推公式, 也是考虑第n个元素是否在选择的k个里的情况
// c(n, k) = c(n-1, k-1) + c(n-1, k)  for n > k > 0
// c(n,0) = c(n,n) = 1
// 0 <= k <= n
func computeBinomialCoefficient(n, k int) int {
	if k == 0 || n == k {
		return 1
	}

	// normalize k
	if k*2 > n {
		k = n - k
	}

	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, i+1)

		for j := 0; j <= minInt(i, k); j++ {
			if j == 0 || j == i {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
	}

	return dp[n][k]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}
