package leetcode

func setZeroes(matrix [][]int) {
	rowSize, columnSize := len(matrix), len(matrix[0])
	r, c := false, false

	for i := 0; i < rowSize; i++ {
		for j := 0; j < columnSize; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0

				if i == 0 {
					r = true
				}
				if j == 0 {
					c = true
				}
			}
		}
	}

	for i := 1; i < rowSize; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < columnSize; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < columnSize; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < rowSize; i++ {
				matrix[i][j] = 0
			}
		}
	}

	if r {
		for j := 1; j < columnSize; j++ {
			matrix[0][j] = 0
		}
	}

	if c {
		for i := 1; i < rowSize; i++ {
			matrix[i][0] = 0
		}
	}
}
