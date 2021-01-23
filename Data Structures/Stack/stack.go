package main

import (
	"fmt"
)

var p = fmt.Println

//Stack struct
type Stack struct {
	slice []int
}

//Push an item to the top
func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}

//Look at the top item
func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

//Remove the top item and return it
func (s *Stack) Pop() int {
	val := s.slice[len(s.slice)-1]
	s.slice = s.slice[0:(len(s.slice)-1)]
	return val
}

func main() {
	var stack *Stack = new(Stack)
	stack.Push(1)
	stack.Peek()
	p(stack.slice)
	stack.Push(3)
	stack.Peek()
	stack.Push(7)
	stack.Peek()
	stack.Push(5)
	p(stack.slice)
	stack.Peek()
	stack.Push(130)
	p(stack.slice)
	stack.Peek()
	stack.Pop()
	p(stack.slice)
	p(stack.Peek())
}