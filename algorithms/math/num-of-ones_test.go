package math

import "testing"

func TestNumOfOnes(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{3, 2},
		{7, 3},
		{16, 1},
	}
	for _, tt := range tests {
		if got := NumOfOnes(tt.n); got != tt.want {
			t.Errorf("NumOfOnes(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
