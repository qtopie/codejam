package leetcode

// & | ~ ^ << >>
// 000 1 222
func singleNumberII(nums []int) int {
	once, twice := 0, 0
	for _, num := range nums {
		// eliminate the number that appear twice
		once = ^twice & (once ^ num)
		// elminate the number that appears 3 times
		twice = ^once & (twice ^ num)
	}
	return once
}
