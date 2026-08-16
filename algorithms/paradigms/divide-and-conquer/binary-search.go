package divide_and_conquer

// BinarySearch searches for val in sorted ascending array nums.
// Returns the index or -1 if not found.
func BinarySearch(nums []int, val int, start, end int) int {
	if start > end {
		return -1
	}

	m := start + (end-start)/2

	if nums[m] == val {
		return m
	} else if nums[m] < val {
		return BinarySearch(nums, val, m+1, end)
	} else {
		return BinarySearch(nums, val, start, m-1)
	}
}
