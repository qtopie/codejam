package dynamic_programming

// 01背包
// 枚举 2^N, 利用复合子问题简化为N*C的复杂度
// i for item, j for capacity limit
// dp[i, j] = max(dp[i-1, j], dp[i-1, j - weights[i]] + values[i])

func knapsack(weights []int, values []int, capacity int) int {
	maxVals := make([]int, capacity+1)

	for i := 0; i < len(weights); i++ {
		for j := capacity; j >= weights[i]; j-- {
			if maxVals[j-weights[i]]+values[i] > maxVals[j] {
				maxVals[j] = maxVals[j-weights[i]] + values[i]
			}
		}
	}

	return maxVals[capacity]
}

// 可重复背包可以转化为01背包
// dp[i, j] = max ( dp [i-1, j], dp[i-1, j - k * weights[i]] + k * values[i])
func multiKnapsack(weights []int, counts []int, values []int, capacity int) int {
	maxVals := make([]int, capacity+1)

	for i := 0; i < len(weights); i++ {
		for k := 0; k < counts[i]; k++ {
			for j := capacity; j >= weights[i]; j-- {
				if maxVals[j-weights[i]]+values[i] > maxVals[j] {
					maxVals[j] = maxVals[j-weights[i]] + values[i]
				}
			}
		}
	}

	return maxVals[capacity]
}
