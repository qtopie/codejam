package leetcode

func minimumIndex(nums []int) int {
	candidate, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count = 1
		} else if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}

	count = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		}
	}

	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		// check very i
		if nums[i] == candidate {
			cnt++
			if cnt > (i+1)/2 && (count-cnt) > (len(nums)-i-1)/2 {
				return i
			}
		}
	}

	return -1
}
