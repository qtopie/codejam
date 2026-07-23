package leetcode

func simplifyPath(path string) string {
	parts := make([]string, 0)

	chs := []rune(path)
	start := 0
	if chs[start] == '/' {
		start++
	}

	// split and remove all slashes
	for i := start; i < len(chs); i++ {
		if chs[i] == '/' {
			if start < i {
				parts = append(parts, string(chs[start:i]))
			}

			start = i + 1
		}
	}

	if start < len(chs) {
		parts = append(parts, string(chs[start:]))
	}

	dropCount := 0
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] == ".." {
			dropCount++
			copy(parts[i:], parts[i+1:])
			parts = parts[:len(parts)-1]
		} else if parts[i] == "." {
			copy(parts[i:], parts[i+1:])
			parts = parts[:len(parts)-1]
		} else if dropCount > 0 {

			if len(parts) < dropCount {
				// clear 0 to i
				if len(parts) == i+1 {
					return "/"
				} else {
					parts = parts[i+1:]
					break
				}
			}

			copy(parts[i:], parts[i+1:])
			parts = parts[:len(parts)-1]
			dropCount--
		}
	}

	if len(parts) == 0 {
		return "/"
	}

	result := ""
	for _, s := range parts {
		result = result + "/" + s
	}

	return result
}
