package math

// GetBit returns whether the k-th bit (0-indexed from right) of n is 1.
func GetBit(n int, k int) bool {
	return (n & (1 << k)) != 0
}

// SetBit sets the k-th bit of n to 1.
func SetBit(n int, k int) int {
	return n | (1 << k)
}

// ClearBit clears the k-th bit of n to 0.
func ClearBit(n int, k int) int {
	return n & ^(1 << k)
}

// ToggleBit toggles the k-th bit of n.
func ToggleBit(n int, k int) int {
	return n ^ (1 << k)
}
