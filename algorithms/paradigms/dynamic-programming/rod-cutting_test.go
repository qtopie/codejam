package dynamic_programming

import "testing"

func Test_rodCutting(t *testing.T) {
	type args struct {
		n       int
		profits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{2, []int{1,7}}, 7},
		{"test2", args{5, []int{1,2,3,4,2}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rodCutting(tt.args.n, tt.args.profits); got != tt.want {
				t.Errorf("rodCutting() = %v, want %v", got, tt.want)
			}
		})
	}
}
