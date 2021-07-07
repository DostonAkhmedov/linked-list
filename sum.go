package main

import linkedlist "github.com/DostonAkhmedov/linked-list/list"

func sum(l1 linkedlist.LinkedList, l2 linkedlist.LinkedList) linkedlist.LinkedList {
	resultList := linkedlist.NewLinkedList()

	var digit1, digit2, div, mod, sum int
	node1, node2 := l1.Head(), l2.Head()

	for node1 != nil || node2 != nil {
		digit1, digit2 = 0, 0
		if node1 != nil {
			digit1 = node1.Value
			node1 = node1.Next
		}
		if node2 != nil {
			digit2 = node2.Value
			node2 = node2.Next
		}

		sum = digit1 + digit2 + div
		div, mod = sum / 10, sum % 10

		resultList.PushBack(mod)
	}

	if div > 0 {
		resultList.PushBack(div)
	}

	return resultList
}
