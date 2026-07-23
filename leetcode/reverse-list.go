package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 反转下一部分
	tmp := reverseList(head.Next)
	// 将末尾指回head节点
	head.Next.Next = head
	// head Next断开
	head.Next = nil

	return tmp
}
