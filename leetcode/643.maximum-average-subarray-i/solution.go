package main

// Category: algorithms
// Level: Easy
// Percent: 48.12459%

// You are given an integer array nums consisting of n elements, and an integer k.
//
// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
//
//
// Example 1:
//
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
//
//
// Example 2:
//
// Input: nums = [5], k = 1
// Output: 5.00000
//
//
//
// Constraints:
//
//
// 	n == nums.length
// 	1 <= k <= n <= 10⁵
// 	-10⁴ <= nums[i] <= 10⁴
//

func findMaxAverage(nums []int, k int) float64 {
	l, r := 0, k-1
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum

	for r < len(nums)-1 {
		l, r = l+1, r+1
		sum = sum - nums[l-1] + nums[r]
		maxSum = max(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}
