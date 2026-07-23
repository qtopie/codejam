package leetcode

func removeDuplicates2(nums []int) int {
	i, j := 0, 0
	for i+j < len(nums) {
		k := i + 1

		for k+j < len(nums) && nums[k] == nums[i] {
			k++
		}

		// 1,1,1
		// 0,1,2,3
		if k > i + 2 {
			// move i+2 to k to last, keep i+1
			copy(nums[i+2:], nums[k:])
			j = j + k - i - 2
			i = i + 2
		} else {
			i = k
		}
	}

	return i
}
