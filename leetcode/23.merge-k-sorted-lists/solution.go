package main

// Category: algorithms
// Level: Hard
// Percent: 59.056843%

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
//
// Merge all the linked-lists into one sorted linked-list and return it.
//
//
// Example 1:
//
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted linked list:
// 1->1->2->3->4->4->5->6
//
//
// Example 2:
//
// Input: lists = []
// Output: []
//
//
// Example 3:
//
// Input: lists = [[]]
// Output: []
//
//
//
// Constraints:
//
//
// 	k == lists.length
// 	0 <= k <= 10⁴
// 	0 <= lists[i].length <= 500
// 	-10⁴ <= lists[i][j] <= 10⁴
// 	lists[i] is sorted in ascending order.
// 	The sum of lists[i].length will not exceed 10⁴.
//

type MinHeap []*ListNode

func (h *MinHeap) Init(lists []*ListNode) {
	*h = (*h)[:0]
	for _, list := range lists {
		if list != nil {
			*h = append(*h, list)
		}
	}

	for i := len(*h)/2 - 1; i >= 0; i-- {
		siftDown(*h, i)
	}
}

func (h *MinHeap) Push(val *ListNode) {
	*h = append(*h, val)

	siftUp(*h, len(*h)-1)
}

func (h *MinHeap) Pop() *ListNode {
	if len(*h) == 0 {
		return nil
	}

	smallest := (*h)[0]
	last := len(*h) - 1
	(*h)[0] = (*h)[last]
	*h = (*h)[:last]
	if len(*h) > 0 {
		siftDown(*h, 0)
	}

	return smallest
}

// from parent to down
func siftDown(h MinHeap, i int) {
	size := len(h)

	for {
		l, r := 2*i+1, 2*i+2
		smallest := i

		if l < size && h[l].Val < h[smallest].Val {
			smallest = l
		}

		if r < size && h[r].Val < h[smallest].Val {
			smallest = r
		}

		if smallest == i {
			return
		}

		h[i], h[smallest] = h[smallest], h[i]
		i = smallest
	}
}

// from leaf to up
func siftUp(h MinHeap, i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h[parent].Val <= h[i].Val {
			return
		}

		h[i], h[parent] = h[parent], h[i]
		i = parent
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	// create a dummy pointer
	dummy := &ListNode{}
	p := dummy

	// build a heap with each smallest elements in lists
	var h MinHeap
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			h.Push(lists[i])
		}
	}

	// loop to pop and add new possible smallest elements util all empty
	for smallest := h.Pop(); smallest != nil; smallest = h.Pop() {
		// remember next
		sNext := smallest.Next
		// cleanup current smallest
		smallest.Next = nil

		// add new smallest and move forward
		p.Next = smallest
		p = p.Next

		// set new smallest candidate for this list
		if sNext != nil {
			h.Push(sNext)
		}
	}

	return dummy.Next
}
