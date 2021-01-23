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
	value    int
	previous *Node
	next     *Node
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
		node.previous = l.tail
		l.tail.next = &node
	}
	l.tail = &node
}

func (l *List) GenerateList() {
	for i := 0; i < 20; i++ {
		l.Push(i)
	}
}

func (l *List) GetNodeByValue(i int) *Node {
	currentNode := l.head
	for currentNode.next != nil {
		if currentNode.value == i {
			break
		}
		currentNode = currentNode.next
	}

	return currentNode
}

func (l *List) DeleteNodeByValue(i int) *List {
	currentNode := l.head
	for currentNode.next != nil {
		if currentNode.value == i {
			if currentNode.next != nil {
				currentNode.previous.next = currentNode.next
				currentNode.next.previous = currentNode.previous
			}
		}
		currentNode = currentNode.next
	}
	if l.head.value == i {
		if l.head.next != nil {
			l.head = l.head.next
		}
	}
	if l.tail.value == i {
		if l.tail.previous != nil {
			l.tail = l.tail.previous
			l.tail.next = nil

		}
	}
	return l
}

func (l *List) PrintAllValues() {
	currentNode := l.head
	for currentNode != nil {
		p(currentNode.value)
		currentNode = currentNode.next
	}
}

func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) Previous() *Node {
	return node.previous
}

func main() {
	linkedlist := &List{}
	linkedlist.GenerateList()

	linkedlist.DeleteNodeByValue(17)
	linkedlist.DeleteNodeByValue(19)
	linkedlist.PrintAllValues()
	p(linkedlist.tail)

}

