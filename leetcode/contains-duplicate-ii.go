package leetcode

// 1,0,1,1
func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int)
	l := 0
	for r := 0; r < len(nums); r++ {
		dict[nums[r]]++

		if r-l > k {
			dict[nums[l]]--
			if dict[nums[l]] == 0 {
				delete(dict, nums[l])
			}
			l++
		}

		if dict[nums[r]] > 1 {
			return true
		}
	}

	return false
}
