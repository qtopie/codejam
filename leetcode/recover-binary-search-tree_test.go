package leetcode

import "testing"

func Test_recoverTreeOfString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//		{"case 1", args{"[3,2,1]"}, "[1,2,3]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recoverTreeOfString(tt.args.s); got != tt.want {
				t.Errorf("recoverTreeOfString() = %v, want %v", got, tt.want)
			}
		})
	}
}
