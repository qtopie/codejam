package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])

	l := 0
	r := m*n - 1

	for l <= r {
		p := l + (r-l)/2
		pivot := matrix[p/n][p%n]
		if pivot == target {
			return true
		} else if pivot < target {
			// update l
			l = p + 1
		} else {
			// update r
			r = p - 1
		}
	}

	return false
}
