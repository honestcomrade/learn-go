package main

import (
	"container/list"
	"fmt"
)

// Item is the basic item we want to push to the linked list
type Item struct {
	Name     string
	Category string
	ID       int
}

// RemoveDupes ...
// write an algorithm to remove dupes from an unsorted singly linked list
func RemoveDupes(l *list.List) {
	for i := l.Back(); i != nil; i = i.Prev() {
		for j := i; j != nil; j = j.Prev() {
			if i.Value == j.Value {
				l.Remove(j)
			}
		}
	}
}

// PrintList is just for printing the items in the list
func PrintList(l *list.List) {
	for temp := l.Back(); temp != nil; temp = temp.Prev() {
		fmt.Println(temp.Value)
	}
}

func main() {
	items := list.New()
	items.PushBack(Item{"Joel", "Person", 1234})
	items.PushBack(Item{"Joel", "Person", 1234})
	PrintList(items)
	RemoveDupes(items)
	PrintList(items)
}
