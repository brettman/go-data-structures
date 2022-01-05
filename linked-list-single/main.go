package main

import "fmt"

type LinkedList struct {
	Head *Node
	Size int
}

// Node - represents a value in our linked list
type Node struct{
	Value string
	Next *Node
}

// Insert - adds a new element to the start of the linked list
func (l *LinkedList) Insert (elem string){
	node := &Node{
		Value: elem,
		Next: l.Head,
	}
	l.Head = node
	l.Size++
}

// DeleteFirst - removes first element from linked list
func (l *LinkedList) DeleteFirst(){
	l.Head = l.Head.Next
	l.Size--
}

// List - iterates through all elements in linked list and prints
func (l *LinkedList) List(){
	current := l.Head
	for current != nil {
		fmt.Printf("%+v\n", current)
		current = current.Next
	}
}

// Search - searchs for a value in the linked list and
//  returns the first node if it finds one, otherwise nil
func (l *LinkedList) Search(elem string) *Node{
	current:= l.Head
	for current != nil {
		if current.Value == elem{
			return current
		}
		current = current.Next
	}
	return nil
}

// Delete - removes an element if it matches the value
func (l *LinkedList) Delete(elem string){
	previous := l.Head
	current := l.Head
	for current != nil{
		if current.Value == elem{
			previous.Next = current.Next
			l.Size--
		}
		previous = current
		current = current.Next
	}
}

func main(){
	fmt.Println("Singly linked list")
	var ll LinkedList

	ll.Insert("one")
	ll.Insert("two")
	ll.Insert("three")

	ll.Delete("two")
	ll.List()
	fmt.Println(ll.Size)
}
