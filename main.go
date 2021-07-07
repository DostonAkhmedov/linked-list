package main

import (
	"github.com/DostonAkhmedov/linked-list/list"
)

func main() {
	list1 := list.NewLinkedList()
	list1.PushFrontMultiple(3, 4, 2)
	list1.Display()

	list2 := list.NewLinkedList()
	list2.PushFrontMultiple(4, 6, 5)
	list2.Display()

	result := sum(list1, list2)
	result.Display()
}
