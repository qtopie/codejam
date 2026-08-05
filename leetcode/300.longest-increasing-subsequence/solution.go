package main

// Category: algorithms
// Level: Medium
// Percent: 59.792007%

// Given an integer array nums, return the length of the longest strictly increasing subsequence.
//
//
// Example 1:
//
// Input: nums = [10,9,2,5,3,7,101,18]
// Output: 4
// Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
//
//
// Example 2:
//
// Input: nums = [0,1,0,3,2,3]
// Output: 4
//
//
// Example 3:
//
// Input: nums = [7,7,7,7,7,7,7]
// Output: 1
//
//
//
// Constraints:
//
//
// 	1 <= nums.length <= 2500
// 	-10⁴ <= nums[i] <= 10⁴
//
//
//
// Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?

// 利用单调性求解， 同时发现序列的长度由最后一个元素决定。
// 保留各个解的最后一个元素（对应的最小的值）在队列的位置, 我们就可以知道这个解的最大长度
func lengthOfLIS(nums []int) int {
	tails := make([]int, 0)
	for _, e := range nums {
		// get lower bound from l, 找现有队列里第一个>= e的数字
		l, r := 0, len(tails)
		for l < r {
			m := l + (r-l)/2
			if tails[m] >= e {
				r = m
			} else {
				l = m + 1
			}
		}

		if l == len(tails) {
			tails = append(tails, e)
		} else {
			tails[l] = e
		}
	}

	return len(tails)
}
