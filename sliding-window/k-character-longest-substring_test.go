package sliding_window

import "testing"

func TestLongestKSubstring(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		expect int
	}{
		{"eceba", 2, 3},         // "ece"
		{"aa", 1, 2},            // "aa"
		{"a", 1, 1},             // "a"
		{"abcabcabc", 2, 2},     // "ab", "bc", etc.
		{"abcabcabc", 3, 9},     // whole string
		{"", 2, 0},              // empty string
		{"aabbcc", 1, 2},        // "aa", "bb", "cc"
		{"aabbcc", 2, 4},        // "aabb", "bbcc"
		{"aabbcc", 3, 6},        // whole string
		{"abaccc", 2, 4},        // "accc"
		{"abaccc", 1, 3},        // "ccc"
		{"abc", 0, 0},           // k = 0
	}

	for _, tt := range tests {
		got := longestKSubstring(tt.s, tt.k)
		if got != tt.expect {
			t.Errorf("longestKSubstring(%q, %d) = %d; want %d", tt.s, tt.k, got, tt.expect)
		}
	}
}