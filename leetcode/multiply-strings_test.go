package leetcode

import (
	"testing"
)

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case 1", args{"3", "1"}, "3"},
		{"case 2", args{"1", "3"}, "3"},
		{"case 3", args{"2", "10"}, "20"},
		{"case 4", args{"10", "3"}, "30"},
		{"case 5", args{"11", "11"}, "121"},
		{"case 6", args{"123", "456"}, "56088"},
		{"case 7", args{"123456789", "987654321"}, "121932631112635269"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtractOne(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case 0", args{"100"}, "99"},
		{"case 1", args{"11"}, "10"},
		{"case 1", args{"1"}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtractOne(tt.args.a); got != tt.want {
				t.Errorf("subtractOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addString(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case 1", args{a: "12", b: "3"}, "15"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addString(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addString() = %v, want %v", got, tt.want)
			}
		})
	}
}
