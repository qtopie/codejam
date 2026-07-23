package dynamic_programming

import (
	"testing"
)


func TestMinCostToClimbStairs(t *testing.T) {
	tests := []struct {
		name  string
		costs []int
		want  int
	}{
		{
			name:  "Example 1",
			costs: []int{10, 15, 20},
			want:  15,
		},
		{
			name:  "Example 2",
			costs: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want:  6,
		},
		{
			name:  "Single step",
			costs: []int{5},
			want:  0,
		},
		{
			name:  "Two steps",
			costs: []int{5, 10},
			want:  5,
		},
		{
			name:  "All zeros",
			costs: []int{0, 0, 0, 0},
			want:  0,
		},
		{
			name:  "Increasing costs",
			costs: []int{1, 2, 3, 4, 5},
			want:  6,
		},
		{
			name:  "Decreasing costs",
			costs: []int{5, 4, 3, 2, 1},
			want:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCostToClimbStairs(tt.costs)
			if got != tt.want {
				t.Errorf("minCostToClimbStairs(%v) = %v, want %v", tt.costs, got, tt.want)
			}
		})
	}
}