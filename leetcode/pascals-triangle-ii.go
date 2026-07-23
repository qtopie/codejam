package leetcode

// 1            0
// 1            1
// 1 2(2*1) rowIndex event number 2
// 1 3          3
// 1 4 6(2 * 3) 4

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	// prevRow := getRow(rowIndex - 1)
	// row := append([]int{1}, prevRow...)
	// // dont count first one and last one
	// for i := 1; i < rowIndex; i++ {
	// 	row[i] = prevRow[i-1] + prevRow[i]
	// }
	row := make([]int, rowIndex+1)
	row[0] = 1

	// iteration
	for i := 2; i <= rowIndex; i++ {

		// fill digits
		for j := 1; j < (rowIndex+1)/2; j++ {
			row[j] = row[j-1] + row[j]
		}

		if i%2 == 0 {
			row[i/2] = 2 * row[i/2-1]
		}
	}

	// fill digits with mirror
	for j := 0; j < (rowIndex+1)/2; j++ {
		row[rowIndex-j] = row[j]
	}

	return row
}
