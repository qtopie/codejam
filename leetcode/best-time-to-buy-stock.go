package leetcode

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/submissions/487267984/
func maxProfit(prices []int) int {
	result := 0
	lowest := prices[0]

	for i := 1; i < len(prices); i++ {
		if result < prices[i]-lowest {
			result = prices[i] - lowest
		} else if prices[i] < lowest {
			lowest = prices[i]
		}
	}

	return result
}
