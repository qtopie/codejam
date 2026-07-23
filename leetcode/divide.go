package leetcode

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	positive := true
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		dividend = -dividend
		positive = false
	}

	// make it positive
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < divisor {
		return 0
	}

	quotient, _ := positiveDivide(dividend, divisor)

	if positive {
		if quotient > 2147483647 {
			return 2147483647
		}
		return quotient
	}

	if quotient > 2147483648 {
		return 2147483647
	}
	return 0 - quotient
}

func positiveDivide(dividend int, divisor int) (int, bool) {
	if dividend == 0 {
		return 0, false
	}
	if divisor == 1 {
		return dividend, true
	}
	if dividend < divisor {
		return 0, true
	}
	if dividend == divisor {
		return 1, false
	}

	approx := divisor
	count := 1

	//upper bound
	for approx < dividend {
		approx += approx
		count += count
	}

	quotient, left := positiveDivide(approx-dividend, divisor)
	if left {
		return count - quotient - 1, true
	}

	return count - quotient, false
}
