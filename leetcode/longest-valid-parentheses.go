package leetcode

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	digits := make([]int, len(s))

	stack := make([]rune, 0)
	indices := make([]int, 0)

	for i, v := range s {
		if len(stack) == 0 {
			// push
			stack = append(stack, v)
			indices = append(indices, i)
		} else if validParentheses(stack[len(stack)-1], v) {
			// set bits
			digits[indices[len(stack)-1]], digits[i] = 1, 1

			// pop
			stack = stack[:len(stack)-1]
			indices = indices[:len(indices)-1]
		} else {
			// push
			stack = append(stack, v)
			indices = append(indices, i)
		}
	}

	return getLongestContinuousOne(digits)
}

func validParentheses(b1, b2 rune) bool {
	if b1 == '(' && b2 == ')' {
		return true
	}

	return false
}

func getLongestContinuousOne(digits []int) int {
	max := 0
	for i := 0; i < len(digits); i++ {
		// set start
		start := i
		for i < len(digits) && digits[i] == 0 {
			i++
		}
		if i >= len(digits) {
			return max
		} else {
			start = i
		}

		end := i
		for i < len(digits) && digits[i] == 1 {
			i++
		}
		end = i - 1

		delta := end - start + 1
		if delta > max {
			max = delta
		}
	}

	return max
}
