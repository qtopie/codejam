package math

// NumOfOnes returns the number of set bits (1s) in binary representation of n (Hamming Weight).
func NumOfOnes(n int) int {
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
