package leetcode

// C_n_k = C_n_k-1 + C_n-1_k-1
// c_1_1 = 1
// n as row, k as column
//
//	(n-1, k-1) --> k included,    (n-1, k) --> k not included
//	 																(n,k)  C_n^k = C_n-1^k + C_n-1^k-1
//
// f(n, k) = f(n-1, k) + f(n-1, k-1)   bottom up
//
// (1,1)
// (2,1) (2,2)
// (3,1) (3,2) (3,3)
// (4,1) (4,2) (4,3) (4,4)
func combine(n int, k int) [][]int {
	// initialize first line [] [][]
	current := make([][][]int, k)
	current[0] = [][]int{{1}}

	for i := 2; i <= n; i++ {
		j := k
		// dp
		for ; j > 1; j-- {
			if j == i {
				// calculate C_i^j, append j
				c := make([]int, j-1)
				copy(c, current[j-2][0]) // copy C_j-1^j-1, only one combination
				c = append(c, j)
				current[j-1] = [][]int{c}
				continue
			}

			cmb := current[j-1]
			for _, c := range current[j-2] {
				newC := make([]int, len(c))
				copy(newC, c)
				newC = append(newC, i)
				cmb = append(cmb, newC)
			}
			current[j-1] = cmb
		}

		// j == 1, add i to result
		current[j-1] = append(current[j-1], []int{i})
	}

	return current[k-1]
}
