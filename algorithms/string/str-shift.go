package stringalgo

import "strings"

// CanRotate checks if s2 can be obtained by shifting/rotating s1.
func CanRotate(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	return strings.Contains(s1+s1, s2)
}
