package dynamic_programming

import "testing"

func Test_computeBinomialCoefficient(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{5, 3}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeBinomialCoefficient(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("computeBinomialCoefficient() = %v, want %v", got, tt.want)
			}
		})
	}
}
