package main

import (
	"container/list"
	"fmt"
)

// Item is the basic item we want to push to the linked list
type Item struct {
	number int
}

// RemoveDupes ...
// write an algorithm to remove dupes from an unsorted singly linked list
func RemoveDupes(l *list.List) {
	// create a map of [Item]bool pairs
	// for each item seen
	nodes := make(map[interface{}]bool)
	// loop through the list
	for p := l.Back(); p != nil; p = p.Prev() {
		// otherwise mark as seen
		nodes[p.Value] = true
	}

	for p := l.Back(); p != nil; p = p.Prev() {
		if nodes[p.Value] {
			nodes[p.Value] = false
		} else {
			l.Remove(p)
			fmt.Println("Removing dupe", p.Value)
		}
	}
}

// PrintList is just for printing the items in the list
func PrintList(l *list.List) {
	for p := l.Back(); p != nil; p = p.Prev() {
		fmt.Println(p.Value)
	}
}

func main() {
	items := list.New()
	items.PushBack(Item{1})
	items.PushBack(Item{1})
	items.PushBack(Item{1})
	items.PushBack(Item{4})
	items.PushBack(Item{5})
	items.PushBack(Item{1})
	fmt.Print("Items before removal:\n")
	PrintList(items)
	RemoveDupes(items)
	fmt.Print("After Removal:\n")
	PrintList(items)
}
