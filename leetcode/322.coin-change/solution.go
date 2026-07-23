package main

import "sort"

// Category: algorithms
// Level: Medium
// Percent: 48.62034%

// coinChange returns the fewest number of coins needed to make up the amount.
// Uses classic DP: dp[i] = min(dp[i], dp[i-coin] + 1)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// coinChangeSkip is an optimized version that limits the DP range.
// Since dp[amount] only depends on dp[amount-coin] for each coin,
// and the smallest coin is `smallest`, we only need dp values up to
// amount-smallest. Once we have those, we compute dp[amount] directly.
func coinChangeSkip(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	smallest := coins[0]

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	// Only compute intermediate states up to amount - smallest.
	// Any i + smallest > amount means dp[i] can't contribute to dp[amount].
	limit := amount - smallest
	for i := 1; i <= limit; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// Compute dp[amount] from cached values.
	for _, coin := range coins {
		if amount >= coin {
			dp[amount] = min(dp[amount], dp[amount-coin]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
