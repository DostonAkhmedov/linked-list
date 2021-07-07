package main

import (
	linkedlist "github.com/DostonAkhmedov/linked-list/list"
)

func main() {
	list := linkedlist.NewLinkedList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushFront(3)
	list.PushFront(2)
	list.Display()
}