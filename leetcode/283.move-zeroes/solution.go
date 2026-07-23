package main


// Category: algorithms
// Level: Easy
// Percent: 63.65157%

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
//
//
// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:
// Input: nums = [0]
// Output: [0]
//
//
// Constraints:
//
//
// 	1 <= nums.length <= 10⁴
// 	-2³¹ <= nums[i] <= 2³¹ - 1
//
//
//
// Follow up: Could you minimize the total number of operations done?

func moveZeroes(nums []int) {
	// no zero start index
	k := -1

	for i := 0; i < len(nums); i++ {
		// if non zeros element found, place it into pos
		if nums[i] != 0 {
			nums[i], nums[k+1] = nums[k+1], nums[i]
			k++
		}
	}
}
