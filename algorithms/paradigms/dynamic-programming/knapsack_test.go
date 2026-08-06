package dynamic_programming

import "testing"

func Test_knapsack(t *testing.T) {
	type args struct {
		weights  []int
		values   []int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{
			weights:  []int{2, 1, 3, 2},
			values:   []int{12, 10, 20, 15},
			capacity: 5,
		},
			37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack(tt.args.weights, tt.args.values, tt.args.capacity); got != tt.want {
				t.Errorf("knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}
