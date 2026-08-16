package divide_and_conquer

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		nums  []int
		val   int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"find element", args{[]int{1, 2, 3, 4, 5}, 4, 0, 4}, 3},
		{"not found", args{[]int{1, 2, 3, 4, 5}, 6, 0, 4}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.nums, tt.args.val, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
