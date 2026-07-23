package leetcode

// 最长公共子序列与最长公共子串，如果枚举，则需要2^n时间复杂度
// 转换思维，考虑每一个字符是否在最长公共子序列或者最长公共子串中， 字符串，从0到n-1开始扩展

// input: abcd, fbacd
// output: 3  (acd or bcd)
// f[i][j] represents max common lengths of s1[0:i+1] s2[0:j+1]
// Recurrence:
//		f[i][j] = max(f[i-1][j], f[i][j-1]) if not matched
//                f[i-1][j-1] + 1  if s1[i] == s2[j]
// reference: https://www.cnblogs.com/7explore-share/p/5927292.html
// https://the-art-of-programming-by-july.readthedocs.io/en/latest/ebook/zh/%E6%9C%80%E9%95%BF%E5%85%AC%E5%85%B1%E5%AD%90%E5%BA%8F%E5%88%97/
func longestCommonSequence(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	chs1 := []byte(s1) // column
	chs2 := []byte(s2) // row

	maxLens := make([]int, len(s2))

	// initialize first row
	for j := 0; j < len(s2); j++ {
		if chs1[0] == chs2[j] {
			maxLens[j] = 1
			// init rest of this row
			for j < len(s2) {
				maxLens[j] = 1
				j++
			}
		}
	}

	for i := 1; i < len(s1); i++ {
		// init f[i][0], that is maxLens[0]
		if maxLens[0] != 1 && s1[i] == s2[0] {
			maxLens[0] = 1
		}

		nextTopLeft := maxLens[0]
		for j := 1; j < len(s2); j++ {
			currentTopLeft := nextTopLeft
			nextTopLeft = maxLens[j]

			if s2[j] == s1[i] {
				maxLens[j] = currentTopLeft + 1
			} else {
				maxLens[j] = max(maxLens[j], maxLens[j-1])
			}
		}
	}

	return maxLens[len(maxLens)-1]
}

// input: abc, abcd
// output: 3 (abc)
// f[i][j] represents max common substring matched for string s1, s2 (ending must match)
// Recurrence:
//		f[i][j] = 0 if not matched
//                f[i-1][j-1] + 1  if s1[i] == s2[j]
func longestCommonSubstring(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	chs1 := []byte(s1) // column
	chs2 := []byte(s2) // row

	// start from 0 to size to make code better to read
	maxLens := make([]int, len(s2)+1)
	maxMatched := 0

	for i := 1; i <= len(s1); i++ {
		for j := len(s2); j > 0; j-- {
			if chs1[i-1] == chs2[j-1] {
				maxLens[j] = maxLens[j-1] + 1
				if maxMatched < maxLens[j] {
					maxMatched = maxLens[j]
				}
			} else {
				maxLens[j] = 0
			}
		}

	}

	return maxMatched
}
