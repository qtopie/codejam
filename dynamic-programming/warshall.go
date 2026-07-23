package dynamic_programming

func warshall(matrix [][]int) [][]int {
	r := matrix
	n := len(matrix)
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if matrix[i-1][j-1] != 1 {
					matrix[i-1][j-1] = matrix[i-1][k-1] | matrix[k-1][j-1]
				}
			}
		}
	}

	return r
}

func floyd(w [][]int) [][]int {
	d := w
	n := len(w)
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if d[i-1][j-1] > d[i-1][k-1]+d[k-1][j-1] {
					d[i-1][j-1] = d[i-1][k-1] + d[k-1][j-1]
				}
			}
		}
	}

	return d
}
