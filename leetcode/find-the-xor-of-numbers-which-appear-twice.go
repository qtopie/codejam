package leetcode

// https://leetcode.com/problems/find-the-xor-of-numbers-which-appear-twice/
func duplicateNumbersXOR(nums []int) int {
	result := 0
	dict := make(map[int]int)

	for _, n := range nums {
		dict[n]++
		if dict[n] == 2 {
			result = result ^ n
		}
	}

	return result
}
