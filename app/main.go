package main

import "fmt"

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	setZeroes(matrix)
	fmt.Println(matrix)
}

func totalNQueensIteratively(n int) int {
	count := 0
	cols := make([]int, n)

	for row, col := 1, 1; row <= n; row++ {
		for ; col <= n; col++ {
			if !conflict(row, col, cols) {
				cols[row-1] = col
				if row == n {
					count++
				} else {
					col = 1
					break
				}
			}
		}

		// when full go back to parent n - 1 row
		if col > n {
			if row == 1 {
				break
			}

			cols[row-1] = 0 // reset column
			col = cols[row-2] + 1
			row = row - 2 // row++ will be executed in next loop
		}
	}

	return count
}

func conflict(row, col int, cols []int) bool {
	for i, v := range cols {
		if v == 0 {
			return false
		}

		if col == v || col-row == cols[i]-i-1 || col+row == cols[i]+i+1 {
			return true
		}
	}

	return false
}

func setZeroes(matrix [][]int) {
	rowSize, columnSize := len(matrix), len(matrix[0])

	for i := 0; i < rowSize; i++ {
		for j := 0; j < columnSize; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 0; i < rowSize; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < columnSize; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < columnSize; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < rowSize; i++ {
				matrix[i][j] = 0
			}
		}
	}
}
