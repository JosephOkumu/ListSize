package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func listSize(l *List) int {
	count := 0
	current := l.Head

	for current != nil {
		count++
		current = current.Next
	}
	return count
}

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
	link := &List{}

	listPushFront(link, "Hello")
	listPushFront(link, "2")
	listPushFront(link, "you")
	listPushFront(link, "man")

	fmt.Println(listSize(link))
}
