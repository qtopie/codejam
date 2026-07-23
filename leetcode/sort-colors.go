package leetcode

func sortColors(nums []int) {
	// 0, i, j, n
	i, j := 0, len(nums)-1

	for k := i; k <= j && i <= j; {
		if nums[k] == 0 {
			nums[k], nums[i] = nums[i], nums[k]
			i++
			k++
		} else if nums[k] == 2 {
			nums[k], nums[j] = nums[j], nums[k]
			j--
		} else {
			k++
		}
	}
}
