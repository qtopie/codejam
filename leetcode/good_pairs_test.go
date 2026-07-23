package leetcode

import (
	"testing"
)

func Test_strongPasswordChecker(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"case 1", args{"a"}, 5},
		// {"case 2", args{"aA1"}, 3},
		// {"case 3", args{"1337C0d3"}, 0},
		// {"case 4", args{"ABABABABABABABABABAB1"}, 2},
		// {"case 5", args{"bbaaaaaaaaaaaaaaacccccc"}, 8},
		// {"case 6", args{"FFFFFFFFFFFFFFF11111111111111111111AAA"}, 23},
		{"case 7", args{"aaaaAAAAAA000000123456"}, 5},
		// {"case 10", args{"Abcdef1"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strongPasswordChecker(tt.args.password); got != tt.want {
				t.Errorf("strongPasswordChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repeatingParts(t *testing.T) {
	type args struct {
		arr []rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]rune("abccdefffff")}, 1},
		{"case 2", args{[]rune("1337C0d3")}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatingParts(tt.args.arr); got != tt.want {
				t.Errorf("repeatingParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
