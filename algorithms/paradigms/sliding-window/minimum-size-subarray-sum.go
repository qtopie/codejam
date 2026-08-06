package sliding_window

// 窗口的单调性, 通过窗口只需遍历一遍数组, 起始位置为l
// 209 https://leetcode.cn/problems/minimum-size-subarray-sum/description/?envType=study-plan-v2&envId=top-interview-150
func minSubArrayLen(target int, nums []int) int {
	minLen := len(nums) + 1
	sum := 0

	for l, r := 0, 0; r < len(nums); r++ {
		// window calculation
		sum += nums[r]

		for sum >= target {
			if l == r {
				return 1
			}

			if r-l+1 < minLen {
				minLen = r - l + 1
			}

			sum -= nums[l]
			l++
		}

	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}
