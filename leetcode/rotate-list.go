package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	size := 0

	p := head
	for {
		size++
		if p.Next == nil {
			// make it circle
			p.Next = head
			break
		}
		p = p.Next
	}

	length := size - k%size

	p = head
	for i := 1; i < length; i++ {
		p = p.Next
	}

	newHead := p.Next
	p.Next = nil
	return newHead
}
