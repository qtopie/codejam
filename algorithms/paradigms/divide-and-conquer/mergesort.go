package divide_and_conquer

func mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	m := start + (end-start)/2

	// divide into two parts
	mergeSort(nums, start, m)
	mergeSort(nums, m+1, end)

	// merging into one  start <= m < k
	merge(nums, start, m, end)
}

func merge(nums []int, l, m, r int) {
	i, j := 0, 0
	k := 0

	leftParts := make([]int, m-l+1)
	rightParts := make([]int, r-m)
	copy(leftParts, nums[l:m+1])
	copy(rightParts, nums[m+1:r+1])

	for ; i < len(leftParts) && j < len(rightParts); k++ {
		// compare and set values
		if leftParts[i] <= rightParts[j] {
			nums[l+k] = leftParts[i]
			i++
		} else {
			nums[l+k] = rightParts[j]
			j++
		}
	}

	if i >= len(leftParts) {
		copy(nums[l+k:], rightParts[j:])
	}

	if j >= len(rightParts) {
		copy(nums[l+k:], leftParts[i:])
	}
}
