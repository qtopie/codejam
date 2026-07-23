package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var endA, endB *ListNode

	for endA = headA; endA.Next != nil; endA = endA.Next {
	}

	for endB = headB; endB.Next != nil; endB = endB.Next {
	}

	// no intersection
	if endA != endB {
		return nil
	}

	// set endA to headB
	endA.Next = headB

	var slow, fast *ListNode
	for slow, fast = headA.Next, headA.Next.Next; slow != fast; slow, fast = slow.Next, fast.Next.Next {
	}

	p1, p2 := slow, headA
	for ; p1 != p2; p1, p2 = p1.Next, p2.Next {
	}

	// break link
	endA.Next = nil

	return p1
}
