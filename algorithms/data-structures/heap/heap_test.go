package heap

import (
	"container/heap"
	"testing"
)

func TestIntHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)

	var got []int
	for h.Len() > 0 {
		got = append(got, heap.Pop(h).(int))
	}

	want := []int{1, 2, 3, 5}
	for i, v := range got {
		if v != want[i] {
			t.Fatalf("heap pop sequence got %v, want %v", got, want)
		}
	}
}

func TestPriorityQueue(t *testing.T) {
	pq := PriorityQueue{
		&Cell{Dx: 0, Dy: 0, Height: 7},
		&Cell{Dx: 0, Dy: 1, Height: 2},
		&Cell{Dx: 0, Dy: 2, Height: 5},
	}
	heap.Init(&pq)

	first := heap.Pop(&pq).(*Cell)
	if first.Height != 2 {
		t.Errorf("expected min height 2, got %d", first.Height)
	}
}
