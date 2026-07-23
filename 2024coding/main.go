package main

import "fmt"

func binSearch(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		if target == nums[m] {
			return m
		} else if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	m := partition0(nums, l, r)
	quickSort(nums, l, m-1)
	quickSort(nums, m+1, r)
}

func partition0(nums []int, l, r int) int {
	if l >= r {
		return l
	}

	// 4, 3, 5, 2
	p := l
	for i := l + 1; i <= r; i++ {
		if nums[i] < nums[l] {
			nums[i], nums[p+1] = nums[p+1], nums[i]
			p++
		}
	}
	nums[l], nums[p] = nums[p], nums[l]

	return p
}

func partition(nums []int, l, r int) int {
	if l >= r {
		return l
	}

	i, j := l+1, r
	p := l

	// split into two parts by pivot
	for i < j {
		// find first left one >= pivot
		for ; i < j && nums[i] < nums[p]; i++ {
		}

		// find first right one < pivot
		for ; i < j && nums[j] >= nums[p]; j-- {
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[i], nums[j] = nums[j], nums[i]

	if nums[j] < nums[p] {
		nums[p], nums[j] = nums[j], nums[p]
	}

	return j
}

func backtrack() {
}

// 4,1,3,5,7,8
func main() {
	nums := []int{1, 3, 5, 7, 10}

	i := binSearch(nums, 7)
	fmt.Println(i)

	// nums2 := []int{1, 8, 5, 7, 10}
	// quickSort(nums2, 0, len(nums2)-1)
	// fmt.Println(nums2)

	nums3 := []int{5, 2, 3, 1}
	quickSort(nums3, 0, len(nums3)-1)
	fmt.Println(nums3)
}
