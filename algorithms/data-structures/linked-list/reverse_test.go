package linkedlist

import (
	"reflect"
	"testing"
)

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var res []int
	for curr := head; curr != nil; curr = curr.Next {
		res = append(res, curr.Val)
	}
	return res
}

func TestReverse(t *testing.T) {
	list := sliceToList([]int{1, 2, 3, 4})
	reversed := Reverse(list)
	got := listToSlice(reversed)
	want := []int{4, 3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Reverse() = %v, want %v", got, want)
	}
}

func TestDeleteRandomNode(t *testing.T) {
	list := sliceToList([]int{1, 2, 3, 4})
	// delete node 2 (list.Next)
	DeleteRandomNode(list.Next)
	got := listToSlice(list)
	want := []int{1, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DeleteRandomNode() = %v, want %v", got, want)
	}
}
