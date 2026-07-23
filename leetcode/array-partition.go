package leetcode

// 1 2, 2 5, 6 6
func arrayPairSum(nums []int) int {
	insertionSort(nums)

	sum := 0
	for i := 0; i < len(nums); i = i + 2 {
		sum = sum + nums[i]
	}

	return sum
}

func insertionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
