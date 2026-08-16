package stringalgo

import "testing"

func TestCanRotate(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{"different length", "A", "AA", false},
		{"valid rotation", "ABCD", "CDAB", true},
		{"not a rotation", "ABCD", "ACBD", false},
		{"empty strings", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanRotate(tt.s1, tt.s2); got != tt.want {
				t.Errorf("CanRotate(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
