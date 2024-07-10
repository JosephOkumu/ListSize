package main

import "fmt"

// NodeL represents a node in the linked list
type NodeL struct {
	Data interface{} // Data stores the value of the node
	Next *NodeL      // Next points to the next node in the list
}

// List represents the linked list
type List struct {
	Head *NodeL // Head points to the first node in the list
	Tail *NodeL // Tail points to the last node in the list
}

// listSize returns the size (number of nodes) of the linked list
func listSize(l *List) int {
	count := 0
	current := l.Head

	for current != nil {
		count++
		current = current.Next
	}
	return count
}

// listPushFront inserts a new node with the given data at the beginning of the linked list
func listPushFront(l *List, data interface{}) {
	newNode := &NodeL{
		Data: data,
		Next: l.Head,
	}
	if l.Head == nil {
		l.Tail = newNode
	}
	l.Head = newNode
}

func main() {
	// Create a new linked list
	link := &List{}

	// Insert nodes at the beginning of the list
	listPushFront(link, "Hello")
	listPushFront(link, "2")
	listPushFront(link, "you")
	listPushFront(link, "man")

	// Print the size of the linked list
	fmt.Println(listSize(link))
}
