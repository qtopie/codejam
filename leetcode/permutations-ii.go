package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0)
	states := make([]int, 0)
	visited := make([]bool, len(nums))

	var backtrack func(states []int, nums []int, visited []bool)
	backtrack = func(states []int, nums []int, visited []bool) {
		if len(states) == len(nums) {
			result := make([]int, len(nums))
			copy(result, states)
			results = append(results, result)
			return
		}

		for i := 0; i < len(nums); i++ {
			// cannot change order for repeated numbers
			if visited[i] || i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}

			visited[i] = true
			states = append(states, nums[i])
			backtrack(states, nums, visited)
			// undo
			visited[i] = false
			states = states[:len(states)-1]
		}
	}

	backtrack(states, nums, visited)

	return results
}
