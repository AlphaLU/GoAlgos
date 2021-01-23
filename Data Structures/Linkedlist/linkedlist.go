package main

import (
	"fmt"
)

var p = fmt.Println

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	next *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(i int) {
	node := Node{value: i}
	if l.head == nil {
		l.head = &node
	} else {
		l.tail.next = &node
	}
	l.tail = &node
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	list := &List{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	p(list.head.next.next)

	p(list.head.Next())
}