package backtracking

import "sort"

func combinationsWithDup(input []int) [][]int {
	sort.Ints(input)
	results := [][]int{}
	path := make([]int, 0)

	var backtracking func(i int)
	backtracking = func(i int) {
		results = append(results, append([]int{}, path...))

		for j := i; j < len(input); j++ {
			if j > i && input[j] == input[j-1] {
				continue
			}
			// try next state
			path = append(path, input[j])
			backtracking(j + 1)
			// restore states
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	return results
}
