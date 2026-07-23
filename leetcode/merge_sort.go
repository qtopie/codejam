package leetcode

func mergeSort(nums []int) []int {
	_mergeSort(nums, 0, len(nums)-1)
	return nums
}

func _mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	middle := start + (end-start+1)/2
	_mergeSort(nums, start, middle-1)
	_mergeSort(nums, middle, end)

	_merge(nums, start, middle, end)
}

func _merge(nums []int, start, middle, end int) {
	leftParts, rightParts := make([]int, 0), make([]int, 0)
	leftParts = append(leftParts, nums[start:middle]...)
	rightParts = append(rightParts, nums[middle:end+1]...)

	i := 0
	j := 0
	k := start

	for ; i < len(leftParts) && j < len(rightParts); k++ {
		if leftParts[i] <= rightParts[j] {
			nums[k] = leftParts[i]
			i++
		} else {
			nums[k] = rightParts[j]
			j++
		}
	}

	if i >= len(leftParts) {
		// copy remaining right parts to nums[k:end+1]
		copy(nums[k:], rightParts[j:])
		return
	}

	if j >= len(rightParts) {
		// copy remaining left parts to nums[k, end+1]
		copy(nums[k:], leftParts[i:])
	}

}
