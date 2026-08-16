package backtracking

import "testing"

func TestNQueens(t *testing.T) {
	// N=4 -> 2, N=8 -> 92
	tests := []struct {
		n    int
		want int
	}{
		{4, 2},
		{8, 92},
	}

	for _, tt := range tests {
		gotRec := NQueens(tt.n)
		if gotRec != tt.want {
			t.Errorf("NQueens(%d) = %d, want %d", tt.n, gotRec, tt.want)
		}

		gotIter := TotalNQueensIteratively(tt.n)
		if gotIter != tt.want {
			t.Errorf("TotalNQueensIteratively(%d) = %d, want %d", tt.n, gotIter, tt.want)
		}
	}
}
