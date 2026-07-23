package leetcode

import "sort"

func strongPasswordChecker(password string) int {

	// if  n < 6, reparing with adding operation: max(repeats, missing, remaining size)
	// if n >=6 && n <= 20, reparing with replacing operation: max(repeats, missing)
	// if n > 20 repeats: replacing or removing, missing: replacing, removing size: removing
	// 	return max(removing) + max(replacing)

	missing := 0
	arr := []rune(password)
	if !upperCaseFound(arr) {
		missing++
	}
	if !lowerCaseFound(arr) {
		missing++
	}
	if !digitFound(arr) {
		missing++
	}

	repeats := repeatingParts(arr)

	if len(arr) < 6 {
		return maxTripple(repeats, missing, 6-len(arr))
	} else if len(arr) >= 6 && len(arr) <= 20 {
		return max(repeats, missing)
	}

	chs, nums := repeatingChs(arr)

	// initialize table
	size := maxOfSlice(nums) / 3
	t := make([][]int, size)
	for i := 0; i < size; i++ {
		t[i] = make([]int, 0)
	}

	for i := 0; i < len(chs); i++ {
		n := nums[i]
		for j := 0; n >= 3; j++ {
			r := n%3 + 1
			t[j] = append(t[j], r)
			n = n - r
		}
	}

	for i := 0; i < size; i++ {
		sort.Ints(t[i])
	}

	removed := len(arr) - 20
	i, j := 0, 0
	for removed > 0 && i < size {
		if j >= len(t[i]) {
			i++
			j = 0
			continue
		}

		removed -= t[i][j]
		if removed >= 0 {
			j++
		}
	}

	if removed > 0 {
		return len(arr) - 20 + missing
	}

	replacing := 0
	if j >= len(t[i]) {
		i++
		j = 0
	} else {
		replacing += len(t[i]) - j
		i++
	}

	for ; i < size; i++ {
		replacing += len(t[i])
	}

	// removed + replacing
	return len(arr) - 20 + max(replacing, missing)
}

func maxTripple(a, b, c int) int {
	return max(max(a, b), max(b, c))
}

func maxOfSlice(nums []int) int {
	max := -1
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	return max
}

func upperCaseFound(arr []rune) bool {
	for _, a := range arr {
		if a >= 'A' && a <= 'Z' {
			return true
		}
	}

	return false
}

func lowerCaseFound(arr []rune) bool {
	for _, a := range arr {
		if a >= 'a' && a <= 'z' {
			return true
		}
	}

	return false
}

func digitFound(arr []rune) bool {
	for _, a := range arr {
		if a >= '0' && a <= '9' {
			return true
		}
	}

	return false
}

func repeatingParts(arr []rune) int {
	count := 0
	start := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[start] {
			if (i - start) >= 3 {
				count += (i - start) / 3
			}

			start = i
		}
	}

	if len(arr)-start >= 3 {
		count += (len(arr) - start) / 3
	}

	return count
}

func repeatingChs(arr []rune) ([]rune, []int) {
	start := 0

	chs := make([]rune, 0)
	nums := make([]int, 0)

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[start] {
			if (i - start) >= 3 {
				chs = append(chs, arr[start])
				nums = append(nums, i-start)
			}

			start = i
		}
	}

	if len(arr)-start >= 3 {
		chs = append(chs, arr[start])
		nums = append(nums, len(arr)-start)
	}

	return chs, nums
}
