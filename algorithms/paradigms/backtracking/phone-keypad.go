package backtracking

var numberAlphabets = []string{
	"",
	"",
	"ABC",
	"DEF",
	"GHI",
	"JKL",
	"MNO",
	"PQRS",
	"TUV",
	"WXYZ",
}

// LookupWordByPhoneNumber returns all possible letter combinations for given digits.
func LookupWordByPhoneNumber(number string) []string {
	if len(number) == 0 {
		return nil
	}

	results := make([]string, 0)

	var searchRecursively func(numbers []int, cols []int, i int)
	searchRecursively = func(numbers []int, cols []int, i int) {
		if i == len(numbers) {
			chs := make([]byte, len(numbers))
			for idx := 0; idx < len(numbers); idx++ {
				n := numbers[idx]
				chs[idx] = numberAlphabets[n][cols[idx]]
			}
			results = append(results, string(chs))
			return
		}

		s := numberAlphabets[numbers[i]]
		for j := 0; j < len(s); j++ {
			cols[i] = j
			searchRecursively(numbers, cols, i+1)
		}
	}

	numbers := make([]int, len(number))
	cols := make([]int, len(number))

	for i, c := range number {
		numbers[i] = int(c - '0')
	}

	searchRecursively(numbers, cols, 0)
	return results
}
