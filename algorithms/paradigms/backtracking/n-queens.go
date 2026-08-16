package backtracking

// NQueens solves the N-Queens problem recursively and returns total distinct solutions count.
func NQueens(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	states := make([]int, n)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row > n {
			count++
			return
		}

		for col := 1; col <= n; col++ {
			if checkStates(states, row, col) {
				states[row-1] = col
				backtrack(row + 1)
				states[row-1] = 0
			}
		}
	}

	backtrack(1)
	return count
}

func checkStates(states []int, row, col int) bool {
	for i := 1; i < row; i++ {
		if states[i-1] == col {
			return false
		}
		if states[i-1]-i == col-row || states[i-1]+i == row+col {
			return false
		}
	}
	return true
}

// TotalNQueensIteratively solves the N-Queens problem iteratively with explicit backtracking.
func TotalNQueensIteratively(n int) int {
	if n <= 0 {
		return 0
	}

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

		// When no column is feasible, backtrack to row - 1
		if col > n {
			if row == 1 {
				break
			}
			cols[row-1] = 0
			col = cols[row-2] + 1
			row = row - 2
		}
	}

	return count
}

func conflict(row, col int, cols []int) bool {
	for i := 0; i < row-1; i++ {
		v := cols[i]
		if v == 0 {
			continue
		}
		if col == v || col-row == v-(i+1) || col+row == v+(i+1) {
			return true
		}
	}
	return false
}
