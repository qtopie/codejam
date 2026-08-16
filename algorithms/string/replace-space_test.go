package stringalgo

import "testing"

func TestReplaceSpace(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"single space", "A B", "A%20B"},
		{"multiple spaces", "We are happy.", "We%20are%20happy."},
		{"no space", "Hello", "Hello"},
		{"empty string", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpace(tt.s); got != tt.want {
				t.Errorf("ReplaceSpace(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
