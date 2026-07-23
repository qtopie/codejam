package leetcode

import "fmt"

func solveNQueens(n int) [][]string {
	results := make([][]string, 0)
	cols := make([]int, n)

	// for each level of tree
	for row, col := 1, 1; row <= n; row++ {
		for ; col <= n; col++ { // for each child of parent
			if !conflict(row, col, cols) {
				cols[row-1] = col
				if row == n {
					// write results
					fmt.Println("found")
					results = append(results, writeToResult(cols))
				} else { // set child and continue
					col = 1
					break // next row
				}
			}

		}

		// backtracking
		if col > n {
			if row == 1 { // termination
				break
			}

			cols[row-1] = 0
			col = cols[row-2] + 1
			row -= 2 // r++ will exucute anyway
		}

	}

	return results
}

func conflict(row, col int, cols []int) bool {
	for i, v := range cols {
		if v == 0 {
			return false
		}

		if col == v || col-row == v-i-1 || col+row == v+i+1 {
			return true
		}
	}
	return false
}

func writeToResult(s []int) []string {
	result := make([]string, 0)
	for _, v := range s {
		chs := make([]byte, len(s))

		for j, _ := range chs {
			if j == v-1 {
				chs[j] = 'Q'
			} else {
				chs[j] = '.'
			}
		}

		result = append(result, string(chs))
	}

	return result
}
