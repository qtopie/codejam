package main

// Category: algorithms
// Level: Hard
// Percent: 51.61526%

// Given two strings s and t, return the number of distinct subsequences of s which equals t.
//
// The test cases are generated so that the answer fits on a 32-bit signed integer.
//
//
// Example 1:
//
// Input: s = "rabbbit", t = "rabbit"
// Output: 3
// Explanation:
// As shown below, there are 3 ways you can generate "rabbit" from s.
// rabbbit
// rabbbit
// rabbbit
//
//
// Example 2:
//
// Input: s = "babgbag", t = "bag"
// Output: 5
// Explanation:
// As shown below, there are 5 ways you can generate "bag" from s.
// babgbag
// babgbag
// babgbag
// babgbag
// babgbag
//
//
// Constraints:
//
//
// 	1 <= s.length, t.length <= 1000
// 	s and t consist of English letters.
//

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}

	dp := make([]int, n+1)
	// 初始状态：空字符串始终能被匹配 1 次
	dp[0] = 1

	for i := 1; i <= m; i++ {
		for j := n; j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] = dp[j-1] + dp[j]
			}
		}

	}

	return dp[n]
}

func _numDistinct(chs []rune, cht []rune) int {
	if len(chs) < len(cht) || len(cht) == 0 {
		return 0
	}

	if chs[len(chs)-1] == cht[len(cht)-1] {
		// [s-1, t] or [s-1, t-1]
		return _numDistinct(chs[:len(chs)-1], cht) + _numDistinct(chs[:len(chs)-1], cht[:len(cht)-1])
	} else {
		// [s-1, t]
		return _numDistinct(chs[:len(chs)-1], cht)
	}

}
