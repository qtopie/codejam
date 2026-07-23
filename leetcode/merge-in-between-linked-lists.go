package leetcode

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	prev := &ListNode{
		Next: list1,
	}

	// fix index
	a = a + 1
	b = b + 1

	p := prev
	for i := 1; i < a; i++ {
		p = p.Next
	}

	// record outer pointers p1, p2
	p1 := p
	for i := a; i <= b; i++ {
		p = p.Next
	}
	p2 := p.Next

	p1.Next = list2
	for p = list2; p.Next != nil; p = p.Next {
	}
	p.Next = p2

	return prev.Next
}
