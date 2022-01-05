package main

import (
	"errors"
	"fmt"
)

type Stack struct{
	Elements []int
}

// Push - handles pushing new elements on top of stack
func (s *Stack) Push (elem int) {

	s.Elements = append(s.Elements, elem)
}

// Pop - retrieve and remove an element from top of stack
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	lastIndex := len(s.Elements) - 1
	var elem int
	elem, s.Elements = s.Elements[lastIndex], s.Elements[:lastIndex]
	return elem, nil
}

//Peek - returnes top eleemnet from stack without removing it
func (s *Stack) Peek() (int, error){
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	} else{
		return s.Elements[len(s.Elements) -1 ], nil
	}
}

// Length - returns the size of the stack
func (s *Stack) Length() int{
	return len(s.Elements)
}

func (s *Stack ) IsEmpty() bool{
	return len(s.Elements) == 0
}

func main(){
	fmt.Println("Stack structure")

	stack := Stack{}
	stack.Push(1)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	e1, _ :=stack.Pop()
	fmt.Println(e1)
	e2, _ :=stack.Pop()
	fmt.Println(e2)
	e3, _ :=stack.Pop()
	fmt.Println(e3)
}
