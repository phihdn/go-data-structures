package main

import (
	"fmt"

	"github.com/phihdn/go-data-structures/linked-list/linkedlist"
)

func main() {
	mylist := linkedlist.LinkedList{}

	// Create nodes
	node1 := &linkedlist.Node{Data: 10}
	node2 := &linkedlist.Node{Data: 20}
	node3 := &linkedlist.Node{Data: 30}
	node4 := &linkedlist.Node{Data: 40}

	// Build the list
	mylist.Prepend(node1)
	mylist.Prepend(node2)
	mylist.Prepend(node3)
	mylist.Prepend(node4)

	// Print the original list
	fmt.Println("Original list:")
	mylist.PrintListData()

	// Delete a value and print the result
	mylist.DeleteWithValue(40)
	fmt.Println("After deleting 40:")
	mylist.PrintListData()
}
