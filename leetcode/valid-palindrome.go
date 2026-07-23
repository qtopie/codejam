package leetcode

func isPalindrome(s string) bool {
	chs := convertToChs(s)
	return isPalindromeChs(chs)
}

func convertToChs(s string) []rune {
	result := make([]rune, 0)
	for _, r := range s {
		if !(r >= rune('A') && r <= rune('Z') || r >= rune('a') && r <= rune('z') || (r >= rune('0') && r <= rune('9'))) {
			continue
		} else if r > rune('Z') {
			result = append(result, 'A'+r-'a')
		} else {
			result = append(result, r)
		}
	}

	return result
}
