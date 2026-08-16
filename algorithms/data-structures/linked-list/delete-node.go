package linkedlist

// DeleteRandomNode deletes the node pointed to by ptr in O(1) by copying next node's value.
// Note: ptr must not be the tail node.
func DeleteRandomNode(ptr *ListNode) {
	if ptr == nil || ptr.Next == nil {
		return
	}

	ptr.Val = ptr.Next.Val
	ptr.Next = ptr.Next.Next
}
