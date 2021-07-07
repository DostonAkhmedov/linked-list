package list

import "fmt"

type node struct {
	value interface{}
	next *node
}

type linkedList struct {
	head *node
	length int
}

type LinkedList interface {
	PushBack(v interface{})
	PushFront(v interface{})
	Remove(index int)
	Display()
	Len() int
}

func NewLinkedList() LinkedList {
	return &linkedList{
		nil,
		0,
	}
}

func (l *linkedList) PushBack(v interface{}) {
	if l.head == nil {
		l.head = &node{value: v, next: nil}
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = &node{value: v, next: nil}
	}
	l.length++
}

func (l *linkedList) PushFront(v interface{}) {
	if l.head == nil {
		l.head = &node{value: v, next: nil}
	} else {
		l.head = &node{value: v, next: l.head}
	}
	l.length++
}

func (l *linkedList) Remove(index int) {
	if l.head != nil {
		idx := 1
		current := l.head
		for current.next != nil {
			if idx + 1 >= index {
				break
			}
			idx++
			current = current.next
		}

		if index == 1 {
			l.head = current.next
		} else {
			current.next = current.next.next
		}
	}
}

func (l *linkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Print(current.value)
		if current.next != nil {
			fmt.Print(", ")
		}
		current = current.next
	}
	fmt.Println()
}

func (l *linkedList) Len() int {
	return l.length
}