package sliding_window

// 单调队列
// 239 https://leetcode.cn/problems/sliding-window-maximum/description/
func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	result := make([]int, 0)
	// desc
	idxQueue := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		j := len(idxQueue) - 1
		for j >= 0 && nums[idxQueue[j]] <= nums[i] {
			j--
		}
		idxQueue = append(idxQueue[:j+1], i)

		// skip first k
		if i < k-1 {
			continue
		}

		// check full
		if len(idxQueue) > 0 && idxQueue[0] < i-k+1 {
			idxQueue = idxQueue[1:]
		}

		result = append(result, nums[idxQueue[0]])
	}

	return result
}
