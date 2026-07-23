package leetcode

// input: 8 output: 2
// x' [a,b] f(a) < 0, f(b) > 0 |a-b| <= 1 取a
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	low, high := 1, x

	// low 单调增加, 永远小于x', high动态变化,保持>=low
	for i := low + (high-low+1)/2; low < high; i = low + (high-low+1)/2 {
		result := i * i
		if result == x {
			return i
		} else if result > x {
			high = i - 1
		} else {
			low = i
		}
	}

	return low
}
