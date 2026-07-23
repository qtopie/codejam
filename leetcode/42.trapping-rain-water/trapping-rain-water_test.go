package leetcode

import "testing"

func Test_trapRainWater(t *testing.T) {
	type args struct {
		heightMap [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"test1", args{[][]int{
		// 	{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1},
		// }}, 4},
		{"test2", args{[][]int{
			{5, 5, 5, 1}, {5, 1, 1, 5}, {5, 1, 5, 5}, {5, 2, 5, 8},
		}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapRainWater(tt.args.heightMap); got != tt.want {
				t.Errorf("trapRainWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
