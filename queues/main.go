package main

import (
	"errors"
	"fmt"
)

// Queue - queue data structure representation
type Queue struct{
	Elements []int
}

// Enqueue - add an eleemnt ot type int to end of queue
func (q *Queue) Enqueue(elem int){
	q.Elements = append(q.Elements, elem)
}

//Dequeue - returns first element from queue
func (q *Queue) Dequeue() (int, error) {
	// return first element
	// remove the element
	// validate queue is not empty
	if len(q.Elements) == 0 {
		return 0, errors.New("empty queue")
	}
	var firstElement int
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	return firstElement, nil
}

// Length - return length of queue
func(q *Queue) Length()int{
	return len(q.Elements)
}

// IsEmpty - returns true if queue is empty
func(q *Queue) IsEmpty() bool{
	return q.Length() == 0
}

// Peek - returns first element from queue without updating queue
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.Elements[0], nil
}

func main(){
	fmt.Println("queues section")

	q := Queue{}
	fmt.Println(q)
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q)
	elem, _ := q.Dequeue()
	fmt.Println(elem)
	fmt.Println(q)
}
