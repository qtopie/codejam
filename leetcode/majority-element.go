package leetcode

func majorityElement(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	cnt1, cdt1 := 1, nums[0]
	cnt2, cdt2 := 0, 0

	for i := 1; i < len(nums); i++ {
		if cnt1 == 0 && nums[i] != cdt2 {
			cnt1 = 1
			cdt1 = nums[i]
		} else if cnt2 == 0 && nums[i] != cdt1 {
			cnt2 = 1
			cdt2 = nums[i]
		} else if nums[i] == cdt1 {
			cnt1++
		} else if nums[i] == cdt2 {
			cnt2++
		} else {
			cnt1--
			cnt2--
		}
	}

	cnt1, cnt2 = 0, 0
	for i := 0; i < len(nums); i++ {
		if cdt1 == nums[i] {
			cnt1++
		} else if cdt2 == nums[i] {
			cnt2++
		}
	}

	result := make([]int, 0)
	threshold := len(nums) / 3

	if cnt1 > threshold {
		result = append(result, cdt1)
	}

	if cnt2 > threshold {
		result = append(result, cdt2)
	}

	return result
}
