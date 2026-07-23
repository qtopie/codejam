package leetcode

// input: s = "aaabb", k = 3; output: 3
func longestSubstring(s string, k int) int {
	longest := 0
	dict := make(map[rune]int)
	for _, c := range s {
		dict[c] = dict[c] + 1
	}
	chs := []rune(s)

	for size := 1; size <= len(dict); size++ {
		cnt := make([]int, 26)

		// m for number of types of characters
		// n is the count of satified characters
		m, n := 0, 0

		i := 0
		for j := 0; j < len(s); j++ {
			// move right to add new element
			cj := chs[j] - 'a'
			cnt[cj]++
			if cnt[cj] == 1 {
				m++
			}
			if cnt[cj] == k {
				n++
			}

			// if surpass ch types limit, move left pointer
			for ; i <= j && m > size; i++ {
				ci := chs[i] - 'a'
				cnt[ci]--
				if cnt[ci] == 0 {
					m--
				}
				if cnt[ci] == k-1 {
					n--
				}

			}

			if m == n && j-i+1 > longest {
				longest = j - i + 1
			}
		}

	}

	return longest
}
