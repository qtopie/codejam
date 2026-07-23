package leetcode

// max substring of numbers
//
// recurrence formula:
// curr = curr + nums[i] if curr > 0
//
//	nums[i] other
//
// max(curr, oldCurr)
//
// e.g, input: [1, 2, 3, 4]
//
//	output: [1, 2, 3, 4]
func maxNums(nums []int) []int {
	size := len(nums)
	if size == 0 {
		return []int{}
	}

	current := nums[0]
	start, end := 0, 0
	max := current
	maxStart, maxEnd := start, end

	for i := 1; i < size; i++ {
		if current > 0 {
			current += nums[i]
			end++
		} else {
			current = nums[i]
			start, end = i, i
		}

		if current > max {
			max = current
			maxStart, maxEnd = start, end
		}
	}

	return nums[maxStart : maxEnd+1]
}

// divide and conquer in python
// https://leetcode.com/problems/maximum-subarray/discuss/1453074/Divide-and-Conquer-in-Python
