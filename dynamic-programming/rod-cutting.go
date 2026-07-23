package dynamic_programming

// f(n) = max(p(n), f(n-j) + p[j]) j from 1 to n - 1
// Given a rod of length n inches and an array of prices that contains prices of all pieces of size smaller than n. Determine the locations where the cuts are to be made for maximum profit.
// https://www.geeksforgeeks.org/cutting-a-rod-dp-13/
func rodCutting(n int, profits []int) int {
	if n == 1 {
		return profits[0]
	}

	maxProfits := make([]int, n)
	maxProfits[0] = profits[0]

	for i := 2; i <= n; i++ {
		maxProfit := profits[i-1]
		for j := 1; j <= i-1; j++ {
			if maxProfit < maxProfits[i-j-1]+profits[j-1] {
				maxProfit = maxProfits[i-j-1] + profits[j-1]
				// cut point for i is determined at j
			}
		}

		maxProfits[i-1] = maxProfit
	}

	return maxProfits[len(maxProfits)-1]
}

// 把一根绳子剪成多段，并且使得每段的长度乘积最大
// f(i) = max(1*f(i-1), 2*f(i-2)...) or (1*(i-1), 2 * (i-2)...)
func cutRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	// from i to n
	for i := 2; i <= n; i++ {
		// cutting, at least one
		for j := 1; j < i; j++ {
			// try cut as j and sub cutting
			dp[i] = max(dp[i], dp[i-j]*j)
			// cut as i and i - j
			dp[i] = max(dp[i], (i-j)*j)
		}
	}

	return dp[n]
}
