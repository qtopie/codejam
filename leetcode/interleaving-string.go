package leetcode

// input: s1: aabcc, s2: dbbca, s3: aadbbcbcac
// 分析： s3由s1和s2组成， 位置可以任意, 通过考察s3的最后一个元素，可以得到以下递推公式
// recurrence:
//    f(i, j) = f(i-1,j) and s1[i] == s3[i+j]		or		f(i, j-1) and s2[j] == s3[i+j]
// 这里不能使用双指针，参考用例4，通过递推公式可以看出，对于f(i,j)要考察两种不同的情况，而双指针办法只能排列出一种
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	if len(s1) == 0 {
		return s2 == s3
	}

	if len(s2) == 0 {
		return s1 == s3
	}

	matches := make([]bool, len(s2)+1)

	matches[0] = true // ("","","")
	for j := 1; j <= len(s2); j++ {
		matches[j] = matches[j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= len(s1); i++ {
		matches[0] = matches[0] && s1[i-1] == s3[i-1]
		for j := 1; j <= len(s2); j++ {
			// upper or left
			matches[j] = (s1[i-1] == s3[i+j-1] && matches[j]) || (s2[j-1] == s3[i+j-1] && matches[j-1])
		}
	}

	return matches[len(s2)]
}
