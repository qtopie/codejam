package leetcode

import "testing"

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case 1", args{"aabcc", "dbbca", "aadbbcbcac"}, true},
		{"case 2", args{"", "", ""}, true},
		{"case 3", args{"aabcc", "dbbca", "aadbbaccc"}, false},
		{"case 4", args{"c", "ca", "cac"}, true},
		{"case 5", args{"aacaac", "aacaaeaac", "aacaaeaaeaacaac"}, false},
		{"case 6", args{"db", "b", "cbb"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
