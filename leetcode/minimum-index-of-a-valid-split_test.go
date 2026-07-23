package leetcode

import "testing"

func Test_minimumIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 1, 1, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumIndex(tt.args.nums); got != tt.want {
				t.Errorf("minimumIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
