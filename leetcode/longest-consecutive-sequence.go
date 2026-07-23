package leetcode

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dict := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		dict[nums[i]]++
	}

	count, maxCount := 0, 0
	for i := 0; i < len(nums); i++ {
		count = 0
		// 两边展开匹配
		for j := nums[i]; dict[j] > 0; j, count = j+1, count+1 {
			// reset
			dict[j] = 0
		}
		for j := nums[i] - 1; dict[j] > 0; j, count = j-1, count+1 {
			dict[j] = 0
		}

		maxCount = max(maxCount, count)
		if maxCount*2 >= len(nums) {
			return maxCount
		}
	}

	return maxCount
}
