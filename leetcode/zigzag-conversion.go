package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	colSize := calcColumnSize(len(s), numRows)

	// init table
	tab := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		tab[i] = make([]byte, colSize)
	}

	// fill tables
	quotient := len(s) / (numRows*2 - 2)
	k := 0
	row := 0
	col := 0

	for i := 1; i <= quotient; i++ {
		// row
		for row = 0; row < numRows; row++ {
			tab[row][col] = s[k]
			k++
		}

		// slash
		for row, col = numRows-2, col+1; row >= 1; row, col = row-1, col+1 {
			tab[row][col] = s[k]
			k++
		}
	}

	if k < len(s) {
		for k < len(s) && k < (numRows*2-2)*quotient+numRows {
			tab[row][col] = s[k]
			row++
			k++
		}

		for row, col = numRows-2, col+1; k < len(s); row, col = row-1, col+1 {
			tab[row][col] = s[k]
			k++
		}
	}

	// copy result
	result := make([]byte, len(s))
	for k, i := 0, 0; i < numRows && k < len(s); i++ {
		for j := 0; j < colSize && k < len(s); j++ {
			if tab[i][j] != 0 {
				result[k] = tab[i][j]
				k++
			}
		}
	}
	return string(result)
}

func calcColumnSize(n, rowSize int) int {
	// step: row * 2 - 2, columns: 1 + row - 2 = row - 1
	quotient := n / (rowSize*2 - 2)
	col := quotient * (rowSize - 1)

	if n <= (rowSize*2-2)*quotient+rowSize {
		return col + 1
	}

	return col + 1 + (n%(rowSize*2-2) - rowSize)
}
