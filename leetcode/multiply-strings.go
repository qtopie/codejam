package leetcode

// 123 * 123
func multiply(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	result := "0"

	for i := len(num1) - 2; i >= 0; i-- {
		zeros := make([]byte, len(num1)-1-i)
		for j := 0; j < len(num1)-1-i; j++ {
			zeros[j] = '0'
		}

		tmp := multiply(string(num1[i]), num2) + string(zeros)
		result = addString(result, tmp)
	}

	for t := string(num1[len(num1)-1]); t != "0"; t = subtractOne(t) {
		result = addString(result, num2)
	}

	return result
}

func addString(a, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	carry := byte('0')
	r := []byte(b)
	for i := 0; i < len(b); i++ {
		if i >= len(a) {
			if carry == '0' {
				return string(r)
			}
			r[len(r)-1-i], carry = addDigit('0', b[len(b)-1-i], carry)
		} else {
			r[len(r)-1-i], carry = addDigit(a[len(a)-1-i], b[len(b)-1-i], carry)
		}
	}

	if carry == '1' {
		return "1" + string(r)
	}

	return string(r)
}

func addDigit(a, b, c byte) (result, carry byte) {
	result = a - '0' + b - '0' + c
	if result > '9' {
		return result - 10, '1'
	}

	return result, '0'
}

func subtractOne(a string) string {

	// clone into slice
	digits := []byte(a)

	for i := len(a) - 1; i >= 0; i-- {
		// find last non-zero
		if a[i] != '0' {
			// when reaching leftmost
			if a[i] == '1' && i == 0 {
				// 1 - 1 = 0
				if len(a) == 1 {
					return "0"
				}

				// 1000 -1 = 999
				s := make([]byte, len(a)-1)
				for j := 0; j < len(a)-1; j++ {
					s[j] = '9'
				}

				return string(s)
			} else {
				digits[i] = a[i] - 1
				for j := i + 1; j < len(a); j++ {
					digits[j] = '9'
				}
				return string(digits)
			}
		}
	}

	return "-1"
}
