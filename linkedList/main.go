package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

// AddToHead method of LinkedList to Add value at the Head
func (linkedList *LinkedList) AddToHead(value int) {
	node := &Node{Value: value}
	if linkedList.Head != nil {
		node.Next = linkedList.Head
	} else {
		linkedList.Tail = node
	}
	linkedList.Head = node
}

// AddToTail method of LinkedList to Add value at the Tail(end)
func (linkedList *LinkedList) AddToTail(value int) {
	node := &Node{Value: value}
	if linkedList.Tail != nil {
		linkedList.Tail.Next = node
	} else {
		linkedList.Head = node
	}
	linkedList.Tail = node
}

// RemoveValue method to Remove Node by value
func (linkedList *LinkedList) RemoveValue(value int) {
	if linkedList.Head == nil {
		return
	}
	if linkedList.Head.Value == value { // Remove head
		linkedList.Head = linkedList.Head.Next
		if linkedList.Head == nil { // Removed the last element
			linkedList.Tail = nil
		}
		return
	}
	for node := linkedList.Head; node.Next != nil; node = node.Next {
		if node.Next.Value == value {
			node.Next = node.Next.Next
			if node.Next == nil { // Removed tail
				linkedList.Tail = node
			}
			return
		}
	}
}

// SearchValue method to find Node by value
func (linkedList *LinkedList) SearchValue(value int) *Node {
	for node := linkedList.Head; node != nil; node = node.Next {
		if node.Value == value {
			return node
		}
	}
	return nil
}

// Iterate method of LinkedList to display linked list
func (linkedList *LinkedList) Iterate() {
	for node := linkedList.Head; node != nil; node = node.Next {
		fmt.Println(node.Value)
	}
}

func main() {
	linked := LinkedList{}
	linked.AddToHead(3)
	linked.AddToHead(5)
	linked.AddToTail(7)
	linked.Iterate()                   // Output: 5 3 7
	fmt.Println(linked.SearchValue(5)) // Output: &{5 0x1040c140}
	linked.RemoveValue(3)
	linked.Iterate() // Output: 5 7
}
