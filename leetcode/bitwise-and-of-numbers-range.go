package leetcode

func rangeBitwiseAnd(left int, right int) int {
	if left == 0 {
		return 0
	}
	if left == right {
		return left
	}

	// get high bits
	n := left & right

	// find delta
	k := right - left

	// delta bits should be set to zeros
	for i := 1; i < k; i = (i << 1) + 1 {
		k = k | i
	}

	return n & (^k)
}
