package main

// Category: algorithms
// Level: Medium
// Percent: 59.122913%

// Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.
//
// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
//
//
// 	For example, "ace" is a subsequence of "abcde".
//
//
// A common subsequence of two strings is a subsequence that is common to both strings.
//
//
// Example 1:
//
// Input: text1 = "abcde", text2 = "ace"
// Output: 3
// Explanation: The longest common subsequence is "ace" and its length is 3.
//
//
// Example 2:
//
// Input: text1 = "abc", text2 = "abc"
// Output: 3
// Explanation: The longest common subsequence is "abc" and its length is 3.
//
//
// Example 3:
//
// Input: text1 = "abc", text2 = "def"
// Output: 0
// Explanation: There is no such common subsequence, so the result is 0.
//
//
//
// Constraints:
//
//
// 	1 <= text1.length, text2.length <= 1000
// 	text1 and text2 consist of only lowercase English characters.
//

// if s1[i] == s2[j], f[i,j] = f[i-1, j-1] + 1
// else f[i,j] = max(f[i-1][j], f[i, j-1])
func longestCommonSubsequence(s1 string, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	m, n := len(s1), len(s2)

	//  1, 2
	//  11, 22
	// dp move from left to right, and use corner to remeber last element
	corner := 0
	dp := make([]int, n+1)

	for i := 1; i <= m; i++ {
		corner = 0
		for j := 1; j <= n; j++ {
			// store next corner
			tmp := dp[j]
			if s1[i-1] == s2[j-1] {
				dp[j] = corner + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}

			corner = tmp
		}
	}

	return dp[n]
}
