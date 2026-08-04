package main

// Category: algorithms
// Level: Hard
// Percent: 44.101707%

// We can scramble a string s to get a string t using the following algorithm:
//
//
// 	If the length of the string is 1, stop.
// 	If the length of the string is > 1, do the following:
//
// 		Split the string into two non-empty substrings at a random index, i.e., if the string is s, divide it to x and y where s = x + y.
// 		Randomly decide to swap the two substrings or to keep them in the same order. i.e., after this step, s may become s = x + y or s = y + x.
// 		Apply step 1 recursively on each of the two substrings x and y.
//
//
//
//
// Given two strings s1 and s2 of the same length, return true if s2 is a scrambled string of s1, otherwise, return false.
//
//
// Example 1:
//
// Input: s1 = "great", s2 = "rgeat"
// Output: true
// Explanation: One possible scenario applied on s1 is:
// "great" --> "gr/eat" // divide at random index.
// "gr/eat" --> "gr/eat" // random decision is not to swap the two substrings and keep them in order.
// "gr/eat" --> "g/r / e/at" // apply the same algorithm recursively on both substrings. divide at random index each of them.
// "g/r / e/at" --> "r/g / e/at" // random decision was to swap the first substring and to keep the second substring in the same order.
// "r/g / e/at" --> "r/g / e/ a/t" // again apply the algorithm recursively, divide "at" to "a/t".
// "r/g / e/ a/t" --> "r/g / e/ a/t" // random decision is to keep both substrings in the same order.
// The algorithm stops now, and the result string is "rgeat" which is s2.
// As one possible scenario led s1 to be scrambled to s2, we return true.
//
//
// Example 2:
//
// Input: s1 = "abcde", s2 = "caebd"
// Output: false
//
//
// Example 3:
//
// Input: s1 = "a", s2 = "a"
// Output: true
//
//
//
// Constraints:
//
//
// 	s1.length == s2.length
// 	1 <= s1.length <= 30
// 	s1 and s2 consist of lowercase English letters.
//
// 找切割点
//

func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	// dp[k][i][j] 表示 len=k, s1[i...], s2[j...]
	dp := make([][][]bool, n+1)
	for k := 0; k <= n; k++ {
		dp[k] = make([][]bool, n)
		for i := 0; i < n; i++ {
			dp[k][i] = make([]bool, n)
		}
	}

	// 1. 初始化长度为 1 的情况
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[1][i][j] = (s1[i] == s2[j])
		}
	}

	// 2. 枚举子串长度 k (从 2 到 n)
	for k := 2; k <= n; k++ {
		// 枚举 s1 的起点 i
		for i := 0; i <= n-k; i++ {
			// 枚举 s2 的起点 j
			for j := 0; j <= n-k; j++ {
				// 枚举分割点长度 p (1 <= p < k)
				for p := 1; p < k; p++ {
					// 情况 A: 不交换
					if dp[p][i][j] && dp[k-p][i+p][j+p] {
						dp[k][i][j] = true
						break
					}
					// 情况 B: 交换
					if dp[p][i][j+k-p] && dp[k-p][i+p][j] {
						dp[k][i][j] = true
						break
					}
				}
			}
		}
	}

	return dp[n][0][0]
}
