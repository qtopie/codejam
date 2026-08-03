package main

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		p    string
		want bool
	}{
		{
			name: "adceb matches *a*b",
			s:    "ad",
			p:    "*a*",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isMatch(tt.s, tt.p)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
