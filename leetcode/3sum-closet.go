package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)

	minDiff := 1 << 32
	closest := target

	for i := 0; i+2 < len(nums); i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			diff := target - sum
			if diff == 0 {
				return target
			}

			if diff < 0 {
				diff = -diff
			}
			if diff < minDiff {
				minDiff = diff
				closest = sum
			}

			if sum < target {
				l++
			} else {
				r--
			}

		}
	}

	return closest
}
