package leetcode

// https://leetcode.com/problems/single-number-iii/
func singleNumber(nums []int) []int {
	xorAll := 0
	// xor to eliminate numbers appear twice
	for _, num := range nums {
		xorAll ^= num
	}

	rightMostSignificantBit := xorAll & -xorAll
	a, b := 0, 0

	for _, num := range nums {
		// a != b
		if num&rightMostSignificantBit != 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
