package leetcode

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 || len(nums) == 0 {
		return 0
	}

	num := 0
	i := 0
	product := 0

	for j := 0; j < len(nums); j++ {
		if product == 0 || i == j {
			product = nums[j]
		} else {
			product = product * nums[j]
		}

		for product >= k {
			if i == j {
				product = 0
			} else {
				num = num + j - i
				product = product / nums[i]
			}
			i++
		}
	}

	// i, i+1 , len(nums)-1
	// len(nums) - i         1

	if i < len(nums) {
		num = num + ((1+len(nums)-i)*(len(nums)-i))/2
	}

	return num
}

func numSubarrayProductLessThanK2(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	i := 0
	product := 1
	count := 0

	for j := 0; j < len(nums); j++ {
		product = product * nums[j]

		for product >= k {
			product = product / nums[i]
			i++
		}

		// for each fixed i,  j  to i is new solution, reverse order
		count = count + j - i + 1
	}

	return count
}
