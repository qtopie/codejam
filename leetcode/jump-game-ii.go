package leetcode

func jump(nums []int) int {
	// 0 -> 0
	minJumps := make([]int, len(nums))

	// set to len(nums) - 1, since we always have at least one solution
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] >= len(nums)-1 {
			return minJumps[i] + 1
		}

		for j := i + 1; j < len(nums) && j <= i+nums[i]; j++ {
			if minJumps[j] == 0 || minJumps[j] > minJumps[i] {
				minJumps[j] = minJumps[i] + 1
			}
		}
	}

	return 0
}
