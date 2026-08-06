package divide_and_conquer

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func partition(nums []int, l, r int) int {
	// init pivot, and the range for sorting
	i, j := l+1, r
	pivot := nums[l]

	// split into two parts until cross over
	for {
		// the left part < pivot
		for ; i < j && nums[i] <= pivot; i++ {
		}
		// the right part is >= pivot, let i <=j for cross over, since i++ for last run
		for ; i < j && nums[j] >= pivot; j-- {
		}

		if i >= j {
			break
		}

		// swap when find the pair of numbers, note that there may not exist
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	nums[l], nums[j] = nums[j], nums[l]
	return j
}

func findKlargest(nums []int, start, end, k int) int {
	if start > end || k <= 0 || k > end-start+1 {
		return -1
	}

	p := partition(nums, start, end)

	d := p - start + 1
	if d == k {
		return nums[p]
	} else if d > k {
		return findKlargest(nums, start, p-1, k)
	} else {
		return findKlargest(nums, p+1, end, k-(p-start+1))
	}

}
