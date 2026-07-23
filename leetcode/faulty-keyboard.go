package leetcode

func finalString(s string) string {
	chs := make([]rune, 0)

	for _, c := range s {
		if c == 'i' {
			_reverse(chs)
		} else {
			chs = append(chs, c)
		}
	}

	return string(chs)
}

func _reverse(chs []rune) {
	size := len(chs)

	for i := 0; i < size/2; i++ {
		chs[i], chs[size-1-i] = chs[size-1-i], chs[i]
	}
}
