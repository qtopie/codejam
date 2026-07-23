package main

// Category: algorithms
// Level: Medium
// Percent: 51.354916%

// Given two integer arrays nums1 and nums2, return the maximum length of a subarray that appears in both arrays.
//
//
// Example 1:
//
// Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
// Output: 3
// Explanation: The repeated subarray with maximum length is [3,2,1].
//
//
// Example 2:
//
// Input: nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
// Output: 5
// Explanation: The repeated subarray with maximum length is [0,0,0,0,0].
//
//
//
// Constraints:
//
//
// 	1 <= nums1.length, nums2.length <= 1000
// 	0 <= nums1[i], nums2[i] <= 100
//

// 最优子结构： 最长的公共子序列由终止位i, j决定， 对于给定的终止位置，其向左边能匹配的最长长度是确定的
// if nums[i] == nums[j] , f[i,j] = 1 + f[i-1][j-1]
// else f[i,j] = 0
// memorize function: right to left
func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}

	m, n := len(nums1), len(nums2)
	dp := make([]int, n+1)

	maxLen := 0
	for i := 1; i <= m; i++ {
		for j := n; j >= 1; j-- {
			if nums1[i-1] != nums2[j-1] {
				dp[j] = 0
			} else {
				dp[j] = dp[j-1] + 1
				maxLen = max(maxLen, dp[j])
			}
		}
	}

	return maxLen
}
