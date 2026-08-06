package divide_and_conquer

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
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
		{"test0", args{[]int{2, 1}, 0, 1}, []int{1, 2}},
		{"test1", args{[]int{13, 4, 3, 2}, 0, 3}, []int{2, 3, 4, 13}},
		{"test2", args{[]int{2, 3, 3, 3, 3}, 0, 4}, []int{2, 3, 3, 3, 3}},
		{"test3", args{[]int{7, 6, 5, 4, 3, 2, 1}, 0, 6}, []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeSort(tt.args.nums, tt.args.start, tt.args.end)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
