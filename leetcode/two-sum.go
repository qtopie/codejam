package leetcode

import (
	"sort"
)

// exactly one solution
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	m[target-nums[0]] = 1 // start from index 1
	for i := 1; i < len(nums); i++ {
		if m[nums[i]] != 0 {
			return []int{m[nums[i]] - 1, i}
		}

		m[target-nums[i]] = i + 1
	}

	return []int{}
}

// return all the triplets [nums[i], nums[j], nums[k]], no duplicates
func threeSum(nums []int, target int) [][]int {
	if len(nums) <= 2 {
		return [][]int{}
	}

	result := make([][]int, 0)

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		// skip same values
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// skip triples too large
		if nums[i] > 0 && 3*nums[i] > target {
			continue
		}

		newTarget := target - nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] == newTarget {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--

				for ; l < r && nums[l] == nums[l-1]; l++ {
				}
				for ; l < r && nums[r] == nums[r+1]; r-- {
				}

			} else if nums[l]+nums[r] > newTarget {
				r--
			} else {
				l++
			}
		}

	}

	return result

}

// unique quadruplets
func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)

	if len(nums) < 4 {
		return result
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 && 4*nums[i] > target {
			continue
		}

		newTarget := target - nums[i]
		subResult := threeSum(nums[i+1:], newTarget)
		for _, triplet := range subResult {
			record := []int{nums[i]}
			record = append(record, triplet...)
			result = append(result, record)
		}

	}

	return result
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	dict1 := twoSumCount(nums1, nums2)
	dict2 := twoSumCount(nums3, nums4)

	count := 0
	for k, c1 := range dict1 {
		if c2, ok := dict2[-k]; ok {
			count = count + c1*c2
		}
	}

	return count
}

func twoSumCount(nums1, nums2 []int) map[int]int {
	dict := make(map[int]int)
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i := 0; i < len(nums1); i++ {
		if i > 0 && nums1[i] == nums1[i-1] {
			continue
		}

		for j := 0; j < len(nums2); j++ {
			if j > 0 && nums2[j] == nums2[j-1] {
				continue
			}

			sum := nums1[i] + nums2[j]

			if c, ok := dict[sum]; ok {
				dict[sum] = c + 1
			} else {
				dict[sum] = 1
			}
		}
	}

	return dict
}
