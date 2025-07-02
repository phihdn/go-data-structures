package linkedlist

import (
	"fmt"
)

// Node represents a node in a linked list
// Each node contains data and a pointer to the next node
type Node struct {
	Data int   // The value stored in the node
	Next *Node // Pointer to the next node in the list
}

// LinkedList represents a linked list data structure
// It maintains a reference to the head of the list and tracks its length
type LinkedList struct {
	Head   *Node // Pointer to the first node in the list
	Length int   // Number of nodes in the list
}

// Prepend adds a new node at the beginning of the list
// Time Complexity: O(1) - constant time operation
// Parameters:
//   - n: The node to be added at the beginning of the list
func (l *LinkedList) Prepend(n *Node) {
	second := l.Head     // Save the current head
	l.Head = n           // Set the new head to the provided node
	l.Head.Next = second // Link the new head to the previous head
	l.Length++           // Increment the list length
}

// PrintListData prints all the values in the linked list
// Time Complexity: O(n) where n is the length of the list
// This method traverses the entire list and prints each node's data
func (l LinkedList) PrintListData() {
	toPrint := l.Head               // Start at the head
	for i := 0; i < l.Length; i++ { // Iterate through the list
		fmt.Printf("%d ", toPrint.Data) // Print the current node's data
		toPrint = toPrint.Next          // Move to the next node
	}
	fmt.Println() // Print a newline at the end
}

// DeleteWithValue removes the first node with the specified value
// Time Complexity:
//   - Best case: O(1) if the value is at the head
//   - Worst case: O(n) if the value is at the end or not found
//
// Parameters:
//   - value: The integer value to search for and delete
func (l *LinkedList) DeleteWithValue(value int) {
	// Handle empty list case
	if l.Length == 0 {
		return // Nothing to delete in an empty list
	}

	// Handle case where the head contains the value
	if l.Head.Data == value {
		l.Head = l.Head.Next // Set head to the second node
		l.Length--           // Decrement the length
		return
	}

	// Search for the node before the one with the target value
	previousToDelete := l.Head
	for previousToDelete.Next.Data != value {
		// Check if we've reached the end of the list
		if previousToDelete.Next.Next == nil {
			return // Value not found in the list
		}
		previousToDelete = previousToDelete.Next
	}

	// Remove the node by updating the Next pointer to skip it
	previousToDelete.Next = previousToDelete.Next.Next
	l.Length-- // Decrement the length
}
