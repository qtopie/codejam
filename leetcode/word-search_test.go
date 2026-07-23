package leetcode

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{board: [][]byte{
			{byte('A'), byte('B'), byte('C'), byte('E')},
			{byte('S'), byte('F'), byte('C'), byte('S')},
			{byte('A'), byte('D'), byte('E'), byte('E')},
		}, word: "SEE"}, true},
		{"test2", args{board: [][]byte{
			{byte('A'), byte('B'), byte('C'), byte('E')},
			{byte('S'), byte('F'), byte('C'), byte('S')},
			{byte('A'), byte('D'), byte('E'), byte('E')},
		}, word: "ABCB"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
