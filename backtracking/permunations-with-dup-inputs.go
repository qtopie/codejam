package backtracking

import "sort"

func fullPermutationsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0)
	used := make([]bool, len(nums))
	path := make([]int, 0)

	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			results = append(results, append([]int{}, path...))
			return
		}

		for j := range nums {
			// nums[j] == nums[j-1] && !used[j-1] 如果包含重复， 只按顺序取一次排列
			if used[j] || (j > 0 && nums[j] == nums[j-1] && !used[j-1]) {
				continue
			}
			used[j] = true
			path = append(path, nums[j])
			backtracking()
			path = path[:len(path)-1]
			used[j] = false
		}
	}

	backtracking()

	return results
}

// arrangements without dup from input
func arrangements(nums []int) [][]int {
	results := make([][]int, 0)
	used := make([]bool, len(nums))
	path := make([]int, 0)

	var backtracking func()
	backtracking = func() {
		results = append(results, append([]int{}, path...))

		for i := range nums {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backtracking()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtracking()

	return results
}
