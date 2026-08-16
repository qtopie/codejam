package math

import "testing"

func TestBitOperators(t *testing.T) {
	if !GetBit(5, 0) || GetBit(5, 1) || !GetBit(5, 2) {
		t.Errorf("GetBit failed on 5")
	}

	if SetBit(4, 0) != 5 {
		t.Errorf("SetBit failed")
	}

	if ClearBit(5, 0) != 4 {
		t.Errorf("ClearBit failed")
	}

	if ToggleBit(5, 0) != 4 || ToggleBit(4, 0) != 5 {
		t.Errorf("ToggleBit failed")
	}
}
