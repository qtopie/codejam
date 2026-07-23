package leetcode

func lengthOfLongestSubstring(s string) int {
	chs := []rune(s)
	longest := 0
	dict := make(map[rune]int)
	l := 0

	for i, c := range s {
		count, _ := dict[c]

		// check duplicated
		if count > 0 {
			for j := l; j <= i; j++ {
				if c == chs[j] {
					l++
					break
				} else {
					dict[chs[j]] = 0
					l++
				}
			}
		} else {
			dict[c] = 1
		}

		if i-l+1 > longest {
			longest = i - l + 1
		}
	}

	return longest
}
