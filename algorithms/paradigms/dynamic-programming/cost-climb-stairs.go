package dynamic_programming

// 题目：爬楼梯的最小代价
// 题目描述：
// 给定一个整数数组 cost，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付了费用，可以爬一个或两个台阶。你可以选择从索引为 0 或索引为 1 的台阶开始。请问到达楼梯顶部的最低花费是多少？
func minCostToClimbStairs(costs []int) int {
	if len(costs) == 0 {
		return 0
	}

	if len(costs) == 1 {
		return 0
	}

	if len(costs) == 2 {
		return min(costs[0], costs[1])
	}
	// dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	dp := make([]int, len(costs))
	dp[0] = costs[0]
	dp[1] = costs[1]

	for i := 2; i < len(costs); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + costs[i]
	}

	return min(dp[len(costs)-1], dp[len(costs)-2])
}
