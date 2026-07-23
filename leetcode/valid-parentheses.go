package leetcode

// push slice
// s := append(arr, e)
// arr = arr[:len(arr) -1 ]
func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	stack := make([]rune, 0)

	for _, v := range s {
		if len(stack) == 0 {
			// push
			stack = append(stack, v)
		} else if bracketMatch(stack[len(stack)-1], v) {
			// pop
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	return len(stack) == 0
}

func bracketMatch(b1, b2 rune) bool {
	if b1 == '(' && b2 == ')' || b1 == '[' && b2 == ']' || b1 == '{' && b2 == '}' {
		return true
	}

	return false
}
