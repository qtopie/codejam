package linkedlist

// ListNode represents a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Reverse reverses a singly linked list recursively.
func Reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := Reverse(head.Next)
	head.Next.Next = head
	head.Next = nil

	return tmp
}
