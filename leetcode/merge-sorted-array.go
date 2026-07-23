package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := 0

	for i >= 0 && j >= 0 {
		if nums2[j] > nums1[i] {
			nums1[len(nums1)-1-k] = nums2[j]
			j--
		} else {
			nums1[len(nums1)-1-k] = nums1[i]
			i--
		}
		k++
	}

	for ; i >= 0; i-- {
		nums1[len(nums1)-1-k] = nums1[i]
		k++
	}

	for ; j >= 0; j-- {
		nums1[len(nums1)-1-k] = nums2[j]
		k++
	}
}
