package stringalgo

import "testing"

func TestCalcStringDistance(t *testing.T) {
	s1 := "kitten"
	s2 := "sitting"
	dist := CalcStringDistance(s1, s2, 0, len(s1)-1, 0, len(s2)-1)
	if dist != 3 {
		t.Errorf("CalcStringDistance(%q, %q) = %d, want 3", s1, s2, dist)
	}
}
