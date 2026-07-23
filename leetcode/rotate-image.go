package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)

	// for each layer
	for i := 0; i < n/2; i++ {
		// for each group
		// width=n-2*i, a[i,i] b[i,i+width-1],c[i+width-1,i+width-1],d[i+width-1,i]
		for j := 0; j < n-i*2-1; j++ {
			// set vars
			a := matrix[i][i+j]
			b := matrix[i+j][n-i-1]
			c := matrix[n-i-1][n-i-1-j]
			d := matrix[n-i-1-j][i]

			// do switch
			matrix[i][i+j] = d
			matrix[i+j][n-i-1] = a
			matrix[n-i-1][n-i-1-j] = b
			matrix[n-i-1-j][i] = c
		}
	}
}
