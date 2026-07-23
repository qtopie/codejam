package leetcode

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150
func maxProfitII(prices []int) int {
	result := 0

	for l, r := 0, 1; r < len(prices); {
		for ; r < len(prices) && prices[r] >= prices[r-1]; r++ {
		}
		result += max(0, prices[r-1]-prices[l])
		l = r
		r = r + 1
	}

	return result
}
