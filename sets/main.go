package main

import (
	"errors"
	"fmt"
)


// Set = representation of a set data structure
type Set struct {
	Elements map[string]struct{}
}

//NewSet - Create a new set
func NewSet() *Set{
	var set Set
	set.Elements = make (map[string]struct{})
	return &set
}

//Add - add an element to the set
func (s *Set) Add (elem string) {
	s.Elements[elem] = struct{}{}
}

// Delete - removes an element from the set, if it exists
func (s *Set) Delete(elem string) error {
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("element not present in set")
	}

	delete(s.Elements, elem)
	return nil
}

//Contains - checks to see if the element exists in the set
func (s *Set) Contains(elem string) bool{
	_, exists := s.Elements[elem]
	return exists
}

//List - prints out the elements in the set
func (s *Set) List(){
	for key, _ := range s.Elements{
		fmt.Println(key)
	}
}

func main(){
	var set = NewSet()
	set.Add("earth")
	set.Add("venus")
	set.Add("mars")
	set.Add("mars")

	set.Delete("mars")
	set.List()

	fmt.Println(set.Contains("venus"))
	fmt.Println(set.Contains("mars"))
}

