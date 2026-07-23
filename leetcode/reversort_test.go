package leetcode

import (
	"reflect"
	"testing"
)

func Test_resortCost(t *testing.T) {
	type args struct {
		L []int
	}
	tests := []struct {
		name     string
		args     args
		wantCost int
	}{
		{"case 1", args{[]int{4, 2, 1, 3}}, 6},
		{"case 2", args{[]int{1, 2}}, 1},
		{"case 3", args{[]int{7, 6, 5, 4, 3, 2, 1}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCost := resortCost(tt.args.L); gotCost != tt.wantCost {
				t.Errorf("resortCost() = %v, want %v", gotCost, tt.wantCost)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		nums  []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test case1", args{[]int{1,2,3,4},0,3}, []int{4,3,2,1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.nums, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
