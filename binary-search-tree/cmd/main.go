package main

import (
	"fmt"

	"github.com/phihdn/go-data-structures/binary-search-tree/bst"
)

func main() {
	// Create a new binary search tree with root key 100
	tree := bst.NewBST(100)

	// Insert some values
	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(25)
	tree.Insert(75)
	tree.Insert(125)
	tree.Insert(175)

	// Search for values
	fmt.Println("Searching for 25:", tree.Search(25))   // Should return true
	fmt.Println("Searching for 100:", tree.Search(100)) // Should return true
	fmt.Println("Searching for 200:", tree.Search(200)) // Should return false

	// Display tree traversals
	fmt.Println("\nTree traversals:")
	fmt.Println("In-order traversal (sorted):", tree.InOrderTraversal())
	fmt.Println("Pre-order traversal:", tree.PreOrderTraversal())
	fmt.Println("Post-order traversal:", tree.PostOrderTraversal())

	// Display min and max values
	fmt.Println("\nMin value:", tree.Min())
	fmt.Println("Max value:", tree.Max())
}
