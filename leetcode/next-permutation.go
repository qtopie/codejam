package leetcode

// 找到最后的升序部分，1243 -> 1342
// 首先，找到24，下一个序列应该刚好大于2, 先找到恰好大于2的数字3
// 然后，将2和3交换, 这里可以发现，3的右边为42是最大的结果，因此需要交换一次变为最小的
func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}

	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i < 0 {
		reverseArr(nums, 0)
		return
	}

	j := len(nums) - 1
	for nums[j] <= nums[i] {
		j--
	}

	nums[i], nums[j] = nums[j], nums[i]
	reverseArr(nums, i+1)
}

func reverseArr(nums []int, start int) {
	for l, r := start, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
