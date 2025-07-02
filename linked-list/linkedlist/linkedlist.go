package linkedlist

import (
	"fmt"
)

// Node represents a node in a linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents a linked list data structure
type LinkedList struct {
	Head   *Node
	Length int
}

// Prepend adds a new node at the beginning of the list
func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Length++
}

// PrintListData prints all the values in the linked list
func (l LinkedList) PrintListData() {
	toPrint := l.Head
	for i := 0; i < l.Length; i++ {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
	}
	fmt.Println()
}

// DeleteWithValue removes the first node with the specified value
func (l *LinkedList) DeleteWithValue(value int) {
	if l.Length == 0 {
		return
	}

	if l.Head.Data == value {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	previousToDelete := l.Head
	for previousToDelete.Next.Data != value {
		if previousToDelete.Next.Next == nil {
			return
		}
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = previousToDelete.Next.Next
	l.Length--
}
