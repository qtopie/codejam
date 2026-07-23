package leetcode

import (
	"reflect"
	"testing"
)

func Test_getModifiedArray(t *testing.T) {
	type args struct {
		n       int
		updates [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty Array",
			args: args{n: 0, updates: nil},
			want: []int{},
		},
		{
			name: "Single Update",
			args: args{n: 5, updates: [][]int{{2, 10}}},
			want: []int{0, 0, 10, 0, 0},
		},
		{
			name: "Multiple Updates",
			args: args{n: 10, updates: [][]int{{1, 5}, {4, 20}, {7, 1}}},
			want: []int{0, 5, 0, 0, 20, 0, 1, 0, 0, 0},
		},
		{
			name: "Update to Existing Value",
			args: args{n: 3, updates: [][]int{{0, 10}, {0, 20}}},
			want: []int{20, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getModifiedArray(tt.args.n, tt.args.updates); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getModifiedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
