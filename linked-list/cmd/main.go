// Package main demonstrates the usage of a singly linked list implementation
package main

import (
	"fmt"

	// Import the linkedlist package from the project
	"github.com/phihdn/go-data-structures/linked-list/linkedlist"
)

// main function serves as an entry point and demonstrates the linked list operations
func main() {
	// Initialize an empty linked list
	mylist := linkedlist.LinkedList{}

	// Create nodes with different integer values
	// Each node contains data and initially has no next pointer
	node1 := &linkedlist.Node{Data: 10}
	node2 := &linkedlist.Node{Data: 20}
	node3 := &linkedlist.Node{Data: 30}
	node4 := &linkedlist.Node{Data: 40}

	// Build the list by prepending nodes
	// Note: Since we're using Prepend, the final list will be in reverse order (40->30->20->10)
	mylist.Prepend(node1) // List: 10
	mylist.Prepend(node2) // List: 20->10
	mylist.Prepend(node3) // List: 30->20->10
	mylist.Prepend(node4) // List: 40->30->20->10

	// Print the original list to console
	fmt.Println("Original list:")
	mylist.PrintListData() // Expected output: 40 30 20 10

	// Demonstrate deletion of a node by value
	// This removes the first occurrence of the value 40 from the list
	mylist.DeleteWithValue(40)
	fmt.Println("After deleting 40:")
	mylist.PrintListData() // Expected output: 30 20 10

	// === Additional demonstration of linked list operations ===
	fmt.Println("\n--- Additional Demo Operations ---")

	// Create a second linked list for more demonstrations
	newList := linkedlist.LinkedList{}

	// Demonstrate building a list by adding elements one by one
	// Note how the list order changes as we prepend each element
	fmt.Println("Building a sorted list:")

	// Add first element to the empty list
	newList.Prepend(&linkedlist.Node{Data: 50})
	fmt.Println("After adding 50:")
	newList.PrintListData() // Expected output: 50

	// Prepend adds to the beginning, so 30 will be first
	newList.Prepend(&linkedlist.Node{Data: 30})
	fmt.Println("After adding 30:")
	newList.PrintListData() // Expected output: 30 50

	// Similarly, 10 will become the first element
	newList.Prepend(&linkedlist.Node{Data: 10})
	fmt.Println("After adding 10:")
	newList.PrintListData() // Expected output: 10 30 50

	// Demonstrate deletion of an element in the middle of the list
	// This tests the linked list's ability to reconnect nodes when a middle node is removed
	fmt.Println("\nDeleting middle element (30):")
	newList.DeleteWithValue(30)
	newList.PrintListData() // Expected output: 10 50

	// Demonstrate adding multiple elements and showing how the list structure evolves
	fmt.Println("\nAdding more elements:")
	newList.Prepend(&linkedlist.Node{Data: 25}) // Add 25 to the beginning
	newList.Prepend(&linkedlist.Node{Data: 35}) // Add 35 to the beginning
	newList.Prepend(&linkedlist.Node{Data: 15}) // Add 15 to the beginning
	newList.PrintListData()                     // Expected output: 15 35 25 10 50

	// Demonstrate the behavior when trying to delete a value that doesn't exist in the list
	// The list should remain unchanged after this operation
	fmt.Println("\nTrying to delete non-existent value (100):")
	newList.DeleteWithValue(100) // 100 is not in the list, so this should have no effect
	newList.PrintListData()      // Expected output: 15 35 25 10 50 (unchanged)

	// Demonstrate accessing the list's length property
	// This shows the current count of nodes in the list
	fmt.Printf("\nList length: %d\n", newList.Length) // Expected output: 5

	// === Demonstrate edge case: operations on an empty list ===
	fmt.Println("\n--- Deleting from Empty List Demo ---")

	// Create a new empty linked list for edge case testing
	emptyList := linkedlist.LinkedList{}

	// Show the initial state of the empty list
	fmt.Println("Empty list:")
	emptyList.PrintListData() // Expected output: (empty line, no elements to print)

	// Demonstrate the behavior when trying to delete from an empty list
	// This tests the linked list's robustness with edge cases
	fmt.Println("\nAttempting to delete from empty list:")
	emptyList.DeleteWithValue(10) // The DeleteWithValue method should handle the empty list case gracefully
	emptyList.PrintListData()     // Expected output: (empty line, still empty)

	// Verify that the length of the empty list is still 0
	fmt.Printf("\nEmpty list length: %d\n", emptyList.Length) // Expected output: 0
}
