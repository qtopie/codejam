package main

import "strconv"

// Category: algorithms
// Level: Medium
// Percent: 35.213764%

// Given a positive integer n, find the smallest integer which has exactly the same digits existing in the integer n and is greater in value than n. If no such positive integer exists, return -1.
//
// Note that the returned integer should fit in 32-bit integer, if there is a valid answer but it does not fit in 32-bit integer, return -1.
//
//
// Example 1:
// Input: n = 12
// Output: 21
// Example 2:
// Input: n = 21
// Output: -1
//
//
// Constraints:
//
//
// 	1 <= n <= 2³¹ - 1
//

func nextGreaterElement(n int) int {
	if n < 10 {
		return -1
	}

	// itoa
	s := strconv.Itoa(n)
	digits := []rune(s)
	found := false

	for i := len(digits) - 2; i >= 0; i-- {
		if digits[i] < digits[i+1] {
			found = true

			// find and insert
			j := 0
			for ; digits[len(digits)-1-j] <= digits[i]; j++ {
			}
			digits[i], digits[len(digits)-1-j] = digits[len(digits)-1-j], digits[i]

			// reverse digits i+1 parts
			l, r := i+1, len(digits)-1
			for k := 0; k <= (r-l-1)/2; k++ {
				digits[l+k], digits[r-k] = digits[r-k], digits[l+k]
			}
			break
		}
	}

	val, _ := strconv.Atoi(string(digits))
	if !found || val > (1<<31-1) {
		return -1
	}
	return val
}
