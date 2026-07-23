package leetcode

func getModifiedArray(n int, updates [][]int) []int {
	if n <= 0 || updates == nil {
		return []int{}
	}

	result := make([]int, n)
	diff := make([]int, n)

	for _, update := range updates {
		i, j := update[0], update[1]
		delta := update[2]

		diff[i] = diff[i] + delta
		if j+1 < n {
			diff[j+1] = diff[j+1] - delta
		}
	}

	result[0] = diff[0]
	for i := 1; i < n; i++ {
		result[i] = result[i-1] + diff[i]
	}

	return result
}
