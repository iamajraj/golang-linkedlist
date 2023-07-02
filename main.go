package main

import (
	"fmt"
)

type Node struct {
	next *Node
	data any
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) push(data any) {
	n := &Node{
		data: data,
		next: nil,
	}

	current_tail := l.tail

	if l.head == nil {
		l.head = n
	}
	if l.tail == nil {
		l.tail = n
	} else {
		current_tail.next = n
		l.tail = n
	}
}

func (l *LinkedList) pop() {
	var current = l.head
	var previous = l.head

	if l.head == l.tail {
		l.head = nil
		return
	}

	for current != nil {
		if current == l.tail {
			previous.next = nil
			l.tail = previous
		}
		previous = current
		current = current.next
	}
}

func (l *LinkedList) getAll() []any {
	var current = l.head
	var data []any
	for current != nil {
		data = append(data, current.data)
		current = current.next
	}
	return data
}

func (l *LinkedList) at(idx int) any {
	var current = l.head
	var count = 0
	for current != nil {
		if count == idx {
			return current.data
		}
		current = current.next
		count++
	}
	return nil
}

func (l *LinkedList) indexOf(data any) int {
	var current = l.head
	var count = 0
	for current != nil {
		if data == current.data {
			return count
		}
		current = current.next
		count++
	}
	return -1
}

func NewLinkedList() LinkedList {
	return LinkedList{
		head: nil,
		tail: nil,
	}
}

func main() {
	myl := NewLinkedList()

	myl.push(1)
	myl.push(2)
	myl.push(3)
	myl.push(4)
	myl.push(10)
	myl.pop()
	fmt.Println(myl.at(3))
	fmt.Println(myl.getAll())
	fmt.Println(myl.indexOf(3))
}
