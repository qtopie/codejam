package leetcode

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var (
		left      *ListNode
		right     *ListNode
		rightHead *ListNode
	)

	p := head

	for p != nil {
		if p.Val >= x {
			if right == nil {
				right = p
				rightHead = p
			} else {
				right.Next = p
				right = right.Next
			}

			p = p.Next
			right.Next = nil
		} else {
			if left == nil {
				left = p
				head = p
			} else {
				left.Next = p
				left = left.Next
			}

			p = p.Next
			left.Next = nil
		}
	}

	if left == nil {
		return rightHead
	}

	left.Next = rightHead
	return head
}
