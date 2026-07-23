package leetcode

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := max(len1, len2)
		if len > end-start+1 {
			start = i - (len-1)/2
			end = i + len/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, start, end int) int {
	for (start >= 0 && end < len(s)) && s[start] == s[end] {
		start--
		end++
	}

	// end - start - 2 + 1
	return end - start - 1
}
