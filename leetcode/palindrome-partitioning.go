package leetcode

func palindromePartition(s string) [][]string {
	results := make([][]string, 0)
	states := make([]string, 0)

	var backtrack func(chs []rune, states []string, start int)
	backtrack = func(chs []rune, states []string, start int) {
		if len(chs) == start {
			result := make([]string, len(states))
			copy(result, states)
			results = append(results, result)
			return
		}

		for i := start; i < len(s); i++ {
			if isPalindromeChs(chs[start : i+1]) {
				states = append(states, string(chs[start:i+1]))
				backtrack(chs, states, i+1)
				states = states[:len(states)-1]
			}
		}
	}

	backtrack([]rune(s), states, 0)
	return results
}

func isPalindromeChs(chs []rune) bool {
	for i, j := 0, len(chs)-1; i < j; i, j = i+1, j-1 {
		if chs[i] != chs[j] {
			return false
		}
	}

	return true
}
