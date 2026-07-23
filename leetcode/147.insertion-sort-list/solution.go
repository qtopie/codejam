package main

// Category: algorithms
// Level: Medium
// Percent: 58.69165%

// Given the head of a singly linked list, sort the list using insertion sort, and return the sorted list's head.
//
// The steps of the insertion sort algorithm:
//
//
// 	Insertion sort iterates, consuming one input element each repetition and growing a sorted output list.
// 	At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list and inserts it there.
// 	It repeats until no input elements remain.
//
//
// The following is a graphical example of the insertion sort algorithm. The partially sorted list (black) initially contains only the first element in the list. One element (red) is removed from the input data and inserted in-place into the sorted list with each iteration.
//
//
// Example 1:
//
// Input: head = [4,2,1,3]
// Output: [1,2,3,4]
//
//
// Example 2:
//
// Input: head = [-1,5,3,4,0]
// Output: [-1,0,3,4,5]
//
//
//
// Constraints:
//
//
// 	The number of nodes in the list is in the range [1, 5000].
// 	-5000 <= Node.val <= 5000
//

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 使用虚拟头节点
	dummy := &ListNode{Next: head}
	// 需要另一个指针来记录尾部的位置
	sortedTail := head

	// loop iterate head
	for p := head.Next; p != nil; {
		if sortedTail.Val <= p.Val {
			sortedTail = p
			p = p.Next
		} else {
			// find next element is nil or >= p
			prev := dummy
			for ; prev.Next != nil && prev.Next.Val < p.Val; prev = prev.Next {
			}

			// keep remaining list
			sortedTail.Next = p.Next

			p.Next = prev.Next
			prev.Next = p

			p = sortedTail.Next
		}
	}

	return dummy.Next
}

func insertionSortList2(head *ListNode) *ListNode {
	dummy := &ListNode{}

	// loop iterate head
	for p := head; p != nil; {
		// keep remaining list
		tmp := p.Next

		// break old link
		// p.Next = nil

		// find next element is nil or >= p
		prev := dummy
		for ; prev.Next != nil && prev.Next.Val < p.Val; prev = prev.Next {
		}

		p.Next = prev.Next
		prev.Next = p

		p = tmp
	}

	return dummy.Next
}
