package leetcode

import "testing"

func Test_isValidRow(t *testing.T) {
	type args struct {
		row []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case 1", args{[]byte{'1', '2'}}, true},
		{"case 1", args{[]byte{'1', '1'}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidRow(tt.args.row); got != tt.want {
				t.Errorf("isValidRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
