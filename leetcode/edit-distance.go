package leetcode

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	} else if len(word2) == 0 {
		return len(word1)
	}

	if word1[0] == word2[0] {
		return minDistance(word1[1:], word2[1:])
	}

	// replace
	min := minDistance(word1[1:], word2[1:])

	// insert
	ins := minDistance(word1, word2[1:])

	// delete
	del := minDistance(word1[1:], word2)

	if ins < min {
		min = ins
	}
	if del < min {
		min = del
	}
	return 1 + min
}
