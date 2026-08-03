package main

//
// [10] Regular Expression Matching is on the run...
//
//
// Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:
//
//
// 	'.' Matches any single character.​​​​
// 	'*' Matches zero or more of the preceding element.
//
//
// Return a boolean indicating whether the matching covers the entire input string (not partial).
//
//
// Example 1:
//
// Input: s = "aa", p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".
//
//
// Example 2:
//
// Input: s = "aa", p = "a*"
// Output: true
// Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
//
//
// Example 3:
//
// Input: s = "ab", p = ".*"
// Output: true
// Explanation: ".*" means "zero or more (*) of any character (.)".
//
//
//
// Constraints:
//
//
// 	1 <= s.length <= 20
// 	1 <= p.length <= 20
// 	s contains only lowercase English letters.
// 	p contains only lowercase English letters, '.', and '*'.
// 	It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.

// dp s[0:i] p[0:j] match
// if p[j] != '*', and if p[j] == '.' || p[j] == s[i], then sp[i,j] = sp[i-1, j-1]
// else sp[i,j] = sp[i, j-2], or if p[j-1] == '.' || p[j-1] == s[i], then sp[i,j] = sp[i-1, j]

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)

	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	// init for empty values, * must have previous valid ch
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] != '*' {
				if p[j-1] == '.' || p[j-1] == s[i-1] {
					dp[i][j] = dp[i-1][j-1]
				}
				// else remain false
			} else if j >= 2 {
				// try zero match first
				dp[i][j] = dp[i][j-2]

				// match .* for s[i-1]
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					// remove one match or zero match
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[m][n]
}
