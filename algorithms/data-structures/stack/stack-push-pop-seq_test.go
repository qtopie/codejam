package stack

import "testing"

func TestIsPopOrder(t *testing.T) {
	pushSeq := []int{1, 2, 3, 4, 5}
	validPopSeq := []int{4, 5, 3, 2, 1}
	invalidPopSeq := []int{4, 3, 5, 1, 2}

	if !IsPopOrder(pushSeq, validPopSeq) {
		t.Errorf("expected valid pop order, got false")
	}

	if IsPopOrder(pushSeq, invalidPopSeq) {
		t.Errorf("expected invalid pop order, got true")
	}
}
