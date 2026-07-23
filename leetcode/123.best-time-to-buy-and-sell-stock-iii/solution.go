package main

// Category: algorithms
// Level: Hard
// Percent: 53.29058%

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// Find the maximum profit you can achieve. You may complete at most two transactions.
//
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
//
//
// Example 1:
//
// Input: prices = [3,3,5,0,0,3,1,4]
// Output: 6
// Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
// Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
//
// Example 2:
//
// Input: prices = [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
// Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.
//
//
// Example 3:
//
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transaction is done, i.e. max profit = 0.
//
//
//
// Constraints:
//
//
// 	1 <= prices.length <= 10⁵
// 	0 <= prices[i] <= 10⁵
//

// maxProfit = f[0,k] + f[k, size-1]
func maxProfit(prices []int) int {

	buy1 := -prices[0]
	sell1 := 0
	buy2 := -prices[0]
	sell2 := 0

	for i := 1; i < len(prices); i++ {
		sell2 = max(sell2, buy2+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy1 = max(buy1, -prices[i])
	}

	return max(sell1, sell2)
}
