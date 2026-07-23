package leetcode

func countVowelStrings(n int) int {
	switch n {
	case 1:
		return 5
	case 2:
		return 15
	case 3:
		return 35
	case 4:
		return 70
	}

	old := &[]int{1, 4, 10, 20, 35}
	curr := &[]int{1, 1, 1, 1, 1}

	for k := 5; k <= n; k++ {
		(*curr)[1] = k
		(*curr)[2] = (*old)[2] + (*curr)[1]
		(*curr)[3] = (*old)[3] + (*curr)[2]
		(*curr)[4] = (*old)[4] + (*curr)[3]

		old, curr = curr, old
	}

	s := *old
	return s[0] + s[1] + s[2] + s[3] + s[4]
}
