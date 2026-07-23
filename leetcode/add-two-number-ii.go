package leetcode

import "strconv"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	s1 := []byte{}
	s2 := []byte{}

	for p := l1; p != nil; p = p.Next {
		s1 = append(s1, strconv.Itoa(p.Val)[0])
	}

	for p := l2; p != nil; p = p.Next {
		s2 = append(s2, strconv.Itoa(p.Val)[0])
	}

	s := addString(string(s1), string(s2))

	head := &ListNode{Val: int(s[0] - '0')}
	p := head
	for i := 1; i < len(s); i++ {
		p.Next = &ListNode{Val: int(s[i] - '0')}
		p = p.Next
	}

	return head
}
