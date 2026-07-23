package leetcode

func deleteDuplicates(head *ListNode) *ListNode {

	p := head

	for p != nil {
		if p.Next == nil {
			break
		}

		if p.Next.Val == p.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return head
}
