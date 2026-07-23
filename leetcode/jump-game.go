package leetcode

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	i := len(nums) - 1
	j := i - 1
	for ; j >= 0 && i-j > nums[j]; j-- {
	}
	if j < 0 {
		return false
	}

	for i = len(nums) - 2; i > 0; i-- {
		if nums[i] != 0 {
			continue
		}

		for j = i - 1; j >= 0 && i-j >= nums[j]; j-- {
		}
		if j < 0 {
			return false
		}
	}

	return true
}
