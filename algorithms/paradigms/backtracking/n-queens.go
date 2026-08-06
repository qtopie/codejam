package backtracking

func nQueens(n int) int {
	count := 0

	var backtrack func(n int, states []int, row int)
	backtrack = func(n int, states []int, row int) {
		if row > n {
			count++
			return
		}

		// dfs
		for col := 1; col <= n; col++ {
			if checkStates(states, row, col) {
				states[row-1] = col
				backtrack(n, states, row+1)

				// restore
				states[row-1] = 0
			}
		}
	}

	n := 8
	backtrack(n, make([]int, n), 1)

	return count
}

func checkStates(states []int, row, col int) bool {
	for i := 1; i < row; i++ {
		if states[i-1] == col {
			return false
		}

		// y = x || y = -x
		if states[i-1]-i == col-row || states[i-1]+i == row+col {
			return false
		}
	}

	return true
}
