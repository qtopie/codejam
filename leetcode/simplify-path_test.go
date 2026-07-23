package leetcode

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//		{"test1", args{"/../"}, "/"},
		//		{"test2", args{"/.../a/../b/c/../d/./"}, "/.../b/d"},
		//	{"test2", args{"/home//foo/"}, "/home/foo"},
		// {"test2", args{"/home/user/Documents/../Pictures"}, "/home/user/Pictures"},
		{"test5", args{"/VcbgE///../../../OV///WRGq/..///.//"}, "/OV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
