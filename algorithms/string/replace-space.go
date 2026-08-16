package stringalgo

// ReplaceSpace replaces space characters in s with "%20".
func ReplaceSpace(s string) string {
	chs := []rune(s)
	spaceCount := 0
	for _, c := range chs {
		if c == ' ' {
			spaceCount++
		}
	}

	if spaceCount == 0 {
		return s
	}

	oldLen := len(chs)
	newLen := oldLen + spaceCount*2
	res := make([]rune, newLen)

	i1 := oldLen - 1
	i2 := newLen - 1

	for i1 >= 0 {
		if chs[i1] == ' ' {
			res[i2] = '0'
			res[i2-1] = '2'
			res[i2-2] = '%'
			i2 -= 3
		} else {
			res[i2] = chs[i1]
			i2--
		}
		i1--
	}

	return string(res)
}
