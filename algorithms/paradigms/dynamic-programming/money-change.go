package dynamic_programming

// 换零钱问题 可转化成重复背包问题
// for coins from 1 to i, amount is j
// f[i, j] = min(f[i-1, j], f[i - 1, j - k * coins[i]] + k)
func calcSmallestCoins(coins []int, amount int) int {
	minAmountCoins := make([]int, amount+1)

	for i := 0; i < len(coins); i++ {
		for k := 1; k <= amount/coins[i]; k++ {
			for j := amount; j >= coins[i]; j-- {
				if minAmountCoins[j] > minAmountCoins[j-k*coins[i]]+k {
					minAmountCoins[j] = minAmountCoins[j-k*coins[i]] + k
				}
			}
		}
	}
	return minAmountCoins[amount]
}
