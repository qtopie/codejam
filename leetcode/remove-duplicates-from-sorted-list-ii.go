package leetcode

// 首先考虑一个一般的情况， 1-2-2-4 我们需要删除2-2
// 这里我们考虑使用双指针， 第一个指针指向当前有效的节点，第2个指针遍历列表
// 只需要比较p1.Next.Val == p2.Val即可，如果相等，继续遍历，发现一个不相等的p2
// 特殊情况1， 开始的几个节点需要删除， 解决办法是增加一个前驱指针，而且能使每个元素的遍历保持一致pre->Next
// 特殊情况2， 最后的几个节点需要删除
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := &ListNode{Next: head}

	p1 := p           // pre of current node
	p2 := p.Next.Next // next node

	for p2 != nil {
		duplicated := false
		for ; p2 != nil && p1.Next.Val == p2.Val; p2 = p2.Next {
			duplicated = true
		}

		if duplicated {
			p1.Next = p2
		} else {
			p1 = p1.Next
		}

		if p2 != nil {
			p2 = p2.Next
		}
	}

	return p.Next
}
