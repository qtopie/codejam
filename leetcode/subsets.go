package leetcode

import "sort"

// Decrease and Conquer
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	// sort nums
	sort.Ints(nums)

	// store value index mapping
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	result := [][]int{{}}
	lastSubsets := [][]int{{}}

	// element size goes from 1 to n
	for i := 1; i <= len(nums); i++ {
		newSubset := make([][]int, 0)

		for _, subset := range lastSubsets {
			// get start index
			start := 0
			if len(subset) > 0 {
				start = m[subset[len(subset)-1]] + 1
				if start >= len(nums) {
					// skip this subset since no element left on right
					continue
				}
			}

			for j := start; j < len(nums); j++ {
				s := make([]int, len(subset))
				copy(s[0:], subset)
				s = append(s, nums[j])
				newSubset = append(newSubset, s)
			}
		}

		lastSubsets = newSubset
		result = append(result, lastSubsets...)
	}

	return result
}

// Decrease and Conquer
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	// sort nums
	sort.Ints(nums)

	// store value index mapping
	m := make(map[int]int)
	for i, v := range nums {
		// get first one if override
		if _, ok := m[v]; !ok {
			m[v] = i
		}
	}

	result := [][]int{{}}
	lastSubsets := [][]int{{}}

	// element size goes from 1 to n
	for i := 1; i <= len(nums); i++ {
		newSubset := make([][]int, 0)

		for _, subset := range lastSubsets {
			// get start index
			start := 0
			if len(subset) > 0 {
				start = getStartOfSubset(subset, m)
				if start >= len(nums) {
					// skip this subset since no element left on right
					continue
				}
			}

			for j := start; j < len(nums); j++ {
				// skip duplicated nums
				if j > start && nums[j] == nums[j-1] {
					continue
				}

				s := make([]int, len(subset))
				copy(s[0:], subset)
				s = append(s, nums[j])
				newSubset = append(newSubset, s)
			}
		}

		lastSubsets = newSubset
		result = append(result, lastSubsets...)
	}

	return result
}

func getStartOfSubset(subset []int, m map[int]int) int {
	start := m[subset[len(subset)-1]] + 1

	for j := len(subset) - 2; j >= 0; j-- {
		if subset[len(subset)-1] == subset[j] {
			start++
		}
	}

	return start
}
