package leetcode

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	results := make([][]int, 0)

	for i, first := range nums {
		remaining := make([]int, len(nums)-1)
		copy(remaining, nums[:i])
		copy(remaining[i:], nums[i+1:])
		subresults := permute(remaining)
		for _, r := range subresults {
			result := []int{first}
			result = append(result, r...)
			results = append(results, result)
		}

	}

	return results
}
