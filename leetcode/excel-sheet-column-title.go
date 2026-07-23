package leetcode

import "slices"

// 1 -> A
// 701 -> ZY
func convertToTitle(columnNumber int) string {
	n := columnNumber

	chs := make([]rune, 0)

	for n > 0 {
		if n%26 == 0 {
			ch := 'Z'
			chs = slices.Insert(chs, 0, ch)
			// 高位26的个数
			n = n/26 - 1
		} else {
			ch := rune(n%26-1) + 'A'
			chs = slices.Insert(chs, 0, ch)
			n = n / 26
		}
	}

	return string(chs)
}
