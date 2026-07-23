package leetcode

var (
	mapping = map[rune][]rune{
		'2': []rune{'a', 'b', 'c'},
		'3': []rune{'d', 'e', 'f'},
		'4': []rune{'g', 'h', 'i'},
		'5': []rune{'j', 'k', 'l'},
		'6': []rune{'m', 'n', 'o'},
		'7': []rune{'p', 'q', 'r', 's'},
		'8': []rune{'t', 'u', 'v'},
		'9': []rune{'w', 'x', 'y', 'z'},
	}
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := []string{""}

	for _, d := range digits {
		newResult := make([]string, 0)
		for _, v := range result {
			for _, ch := range mapping[d] {
				newCh := append([]rune(v), ch)
				newResult = append(newResult, string(newCh))
			}
		}

		result = newResult
	}

	return result
}
