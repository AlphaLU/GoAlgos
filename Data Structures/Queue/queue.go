package main

import (
	"fmt"
)

var p = fmt.Println


//Struct of queue object
type Queue struct {
	slice []int
}

//Add an element
func (q *Queue) Enqueue(i int) {
	q.slice = append(q.slice, i)
}

//Removes the first element
func (q *Queue) Dequeue() int {
	val := q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return val
}


func main() {
	p("Started")
	var q *Queue = new(Queue)
	q.Enqueue(5)
	q.Enqueue(9)
	q.Enqueue(3)
	q.Enqueue(2)
	p(q.slice)
	q.Dequeue()
	p(q.slice)
	
}