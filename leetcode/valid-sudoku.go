package leetcode

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 && len(board[0]) != 9 {
		return false
	}

	// rows
	for i := 0; i < 9; i++ {
		if !isValidRow(board[i]) {
			return false
		}
	}

	// columns
	for j := 0; j < 9; j++ {
		column := make([]byte, 9)
		for i := 0; i < 9; i++ {
			column[i] = board[i][j]
		}
		if !isValidRow(column) {
			return false
		}
	}

	// blocks
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// for each block
			block := make([]byte, 9)
			z := 0
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					block[z] = board[3*i+x][3*j+y]
					z++
				}
			}

			if !isValidRow(block) {
				return false
			}
		}
	}

	return true
}

func isValidRow(row []byte) bool {
	exists := 0

	for _, c := range row {
		if c != '.' {
			if c < '1' || c > '9' {
				return false
			}

			i := c - '0'
			if exists>>(i-1)&1 == 1 {
				return false
			}

			// update exists
			exists |= (1 << (i - 1))
		}
	}

	return true
}
