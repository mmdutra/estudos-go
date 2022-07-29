package linkedlist

import (
	"errors"
	"fmt"
)

type node struct {
	Value int8
	next  *node
}

type Linkedlist struct {
	head   *node
	length int
}

func (l *Linkedlist) Insert(value int8) {
	second := l.head
	l.head = &node{
		Value: value,
		next:  nil,
	}
	l.head.next = second
	l.length++
}

func (l *Linkedlist) Remove(value int8) error {
	node := l.head
	for {
		if node == nil {
			return errors.New("Node not found")
		}

		if node.next != nil && node.next.Value == value {
			next_next := node.next.next
			node.next = next_next
			return nil
		}
		node = node.next
	}
}

func (l Linkedlist) Print() string {
	node := l.head
	var str string
	for {
		if node == nil {
			return str
		}

		str += fmt.Sprint(node.Value) + " "
		node = node.next
	}
}
