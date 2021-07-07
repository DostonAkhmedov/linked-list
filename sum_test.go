package main

import (
	"github.com/DostonAkhmedov/linked-list/list"
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		list1, list2, missing []int
	}{
		{
			[]int{2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 0, 7},
		},
		{
			[]int{0},
			[]int{0},
			[]int{0},
		},
		{
			[]int{9, 9, 9, 9, 9, 9, 9},
			[]int{9, 9, 9, 9},
			[]int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, test := range testCases {
		list1, list2, missing := list.NewLinkedList(), list.NewLinkedList(), list.NewLinkedList()
		list1.PushBackMultiple(test.list1...)
		list2.PushBackMultiple(test.list2...)
		missing.PushBackMultiple(test.missing...)

		result := sum(list1, list2)
		if !AreEqual(missing, result) {
			t.Logf("l1 = %s\n", list1)
			t.Logf("l2 = %s\n", list2)
			t.Errorf("\nmissed = %s\nresult = %s", missing, result)
		}
	}
}

func AreEqual(l1 list.LinkedList, l2 list.LinkedList) bool {
	if l1.Len() != l2.Len() {
		return false
	}

	node1, node2 := l1.Head(), l2.Head()
	for node1 != nil && node2 != nil {
		if node1.Value != node2.Value {
			return false
		}

		node1 = node1.Next
		node2 = node2.Next
	}

	return true
}
