package list

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value int
	Next *Node
}

type linkedList struct {
	head *Node
	length int
}

type LinkedList interface {
	PushBack(v int)
	PushBackMultiple(arr ...int)
	PushFront(v int)
	PushFrontMultiple(arr ...int)
	Remove(index int)
	Display()
	Len() int
	Head() *Node
}

func NewLinkedList() LinkedList {
	return &linkedList{
		nil,
		0,
	}
}

func (l *linkedList) Head() *Node {
	return l.head
}

func (l *linkedList) Len() int {
	return l.length
}

func (l *linkedList) PushBack(v int) {
	if l.head == nil {
		l.head = &Node{Value: v, Next: nil}
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &Node{Value: v, Next: nil}
	}
	l.length++
}

func (l *linkedList) PushFront(v int) {
	if l.head == nil {
		l.head = &Node{Value: v, Next: nil}
	} else {
		l.head = &Node{Value: v, Next: l.head}
	}
	l.length++
}

func (l *linkedList) Remove(index int) {
	if l.head != nil {
		idx := 1
		current := l.head
		for current.Next != nil {
			if idx + 1 >= index {
				break
			}
			idx++
			current = current.Next
		}

		if index == 1 {
			l.head = current.Next
		} else {
			current.Next = current.Next.Next
		}

		l.length--
	}
}

func (l *linkedList) PushBackMultiple(arr ...int) {
	for _, v := range arr {
		l.PushBack(v)
	}
}

func (l *linkedList) PushFrontMultiple(arr ...int) {
	for _, v := range arr {
		l.PushFront(v)
	}
}

func (l *linkedList) Display() {
	fmt.Println(l)
}

func (l *linkedList) String() string {
	result := "["
	if l.head != nil {
		current := l.head
		for current != nil {
			result += strconv.Itoa(current.Value)
			if current.Next != nil {
				result += ", "
			}

			current = current.Next
		}
	}
	result += "]"

	return result
}

