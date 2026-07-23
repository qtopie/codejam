package leetcode

import (
	"testing"
)

func Test_longestCommonSequence(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{s1: "abc", s2: "def"}, 0},
		{"case 2", args{s1: "abcd", s2: "cdef"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSequence(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("longestCommonSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubstring(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{s1: "abc", s2: "def"}, 0},
		{"case 2", args{s1: "abcd", s2: "cdef"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubstring(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("longestCommonSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
