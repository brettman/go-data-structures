package main

import (
	"errors"
	"fmt"
)

// Queue - queue data structure representation
type PriorityQueue struct{
	High []int
	Low  []int
}

// Enqueue - add an eleemnt ot type int to end of queue
func (q *PriorityQueue) Enqueue(elem int, highPriority bool){
	if highPriority {
		q.High = append(q.High, elem)
	}else{
		q.Low = append(q.Low, elem)
	}
}

//Dequeue -
func (q *PriorityQueue) Dequeue() (int, error) {
	// if len of high priority is not zero, return first elem from high
	if len(q.High) != 0 {
		var firstElement int
		firstElement, q.High= q.High[0], q.High[1:]
		return firstElement, nil
	}
	if len(q.Low) != 0 {
		var firstElement int
		firstElement, q.Low= q.Low[0], q.Low[1:]
		return firstElement, nil
	}

	return 0, errors.New("all queues empty")
}

// Length - return length of queue
func(q *PriorityQueue) Length()int{
	return len(q.High) + len(q.Low)
}

// IsEmpty - returns true if queue is empty
func(q *PriorityQueue) IsEmpty() bool{
	return q.Length() == 0
}

// Peek - returns first element from queue without updating queue
func (q *PriorityQueue) Peek() (int, error) {
	if len(q.High) != 0 {
		return q.High[0], nil
	}
	if len(q.Low) != 0 {
		return q.Low[0], nil
	}

	return 0, errors.New("Empty queue")
}

func main(){
	fmt.Println("queues section")

	q := PriorityQueue{}
	q.Enqueue(1, true)
	fmt.Println(q)
	q.Enqueue(10, false)
	fmt.Println(q)
	elem, _ := q.Dequeue()
	fmt.Println(elem)
	q.Enqueue(2, true)
	elem, _ = q.Dequeue()
	fmt.Println(elem)
	elem, _ = q.Dequeue()
	fmt.Println(elem)
	fmt.Println(q)
}
