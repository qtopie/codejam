package leetcode

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}

	carry := byte('0')
	result := []byte(b)

	for i := 0; i < len(b); i++ {
		if i >= len(a) {
			if carry == '0' {
				return string(result)
			}

			result[len(b)-i-1], carry = add('0', b[len(b)-i-1], carry)

		} else {
			result[len(b)-i-1], carry = add(a[len(a)-i-1], b[len(b)-i-1], carry)
		}
	}

	if carry == '1' {
		return "1" + string(result)
	}

	return string(result)
}

func add(x, y, carry byte) (sum byte, newCarry byte) {
	count := 0
	if x == '1' {
		count++
	}
	if y == '1' {
		count++
	}

	if carry == '1' {
		count++
	}

	switch count {
	case 0:
		return '0', '0'
	case 1:
		return '1', '0'
	case 2:
		return '0', '1'
	}

	return '1', '1'
}
