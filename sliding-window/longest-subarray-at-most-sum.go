package sliding_window

// 如果包含负数怎么处理? 算区间和累计成非负数
// Longest subarray having sum of elements atmost K, elements are positive or 0
// https://www.geeksforgeeks.org/longest-subarray-sum-elements-atmost-k/
func longestSubarrayAtMostSum(nums []int, limit int) int {
	maxLen := 0
	sum := 0

	for l, r := 0, 0; r < len(nums); r++ {
		if nums[r] > limit {
			sum = 0
			l = r + 1
			continue
		}

		sum += nums[r]

		// move left point when not match
		for sum > limit {
			sum -= nums[l]
			l++
		}

		if maxLen < r-l+1 {
			maxLen = r - l + 1
		}

	}

	return maxLen
}
