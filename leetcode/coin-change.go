package leetcode

// https://leetcode.com/problems/coin-change/description/?envType=study-plan-v2&envId=top-interview-150
// f(j) = min(f(j), 1+f(j-coins[i]))
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	upperBound := 1 << 32
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = upperBound
	}

	for j := 0; j <= amount; j++ {
		for i := 0; i < len(coins); i++ {
			if j >= coins[i] {
				dp[j] = min(dp[j], 1+dp[j-coins[i]])
			}
		}
	}

	result := dp[amount]
	if result == upperBound {
		return -1
	}

	return result
}
