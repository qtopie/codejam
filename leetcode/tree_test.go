package leetcode

import (
	"reflect"
	"testing"
)

func Test_toArray(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	got := root.ToArray()
	want := []string{"1", "2", "3"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("toArray = %v, want %v", got, want)
	}
}

func Test_clone(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	newRoot := clone(root)

	got := newRoot.ToArray()
	want := []string{"1", "2", "3"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("toArray = %v, want %v", got, want)
	}

}
