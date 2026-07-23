package leetcode

func maximumSubarraySum(nums []int, k int) int64 {
	result := 0
	sum := 0
	dict := make(map[int]int)

	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		dict[nums[r]]++

		// filter duplicates and move forward
		for dict[nums[r]] > 1 {
			sum -= nums[l]
			dict[nums[l]]--
			l++
		}

		// meet conditions
		if r-l+1 == k {
			result = max(sum, result)
			sum -= nums[l]
			dict[nums[l]]--
			l++
		}

	}

	return int64(result)
}
