package leetcode

// n = 1,   0 1
// n = 2,   00 01 11 10
// n = 3,   000 ->  001 ->  011 ->  010
//
//	100 <- 101 <- 111 <- 110 <-
func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}

	// calculate last
	codes := grayCode(n - 1)

	for i := len(codes) - 1; i >= 0; i-- {
		codes = append(codes, codes[i]|1<<(n-1))
	}

	return codes
}
