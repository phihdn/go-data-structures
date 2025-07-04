package main

import (
	"fmt"

	"github.com/phihdn/go-data-structures/heap/minheap"
)

func main() {
	// Initialize a new min heap
	fmt.Println("Initializing MinHeap...")
	heap := minheap.InitMinHeap()

	// Insert elements
	fmt.Println("Inserting elements: 10, 5, 20, 3, 15")
	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(20)
	heap.Insert(3)
	heap.Insert(15)

	// Print the heap
	fmt.Println("Heap after insertions:", heap)

	// Get the minimum value without extracting
	min, ok := heap.GetMin()
	if ok {
		fmt.Println("Minimum value:", min)
	}

	// Extract the minimum value
	extracted, ok := heap.Extract()
	if ok {
		fmt.Println("Extracted minimum:", extracted)
		fmt.Println("Heap after extraction:", heap)
	}

	// Build a heap from an array
	fmt.Println("\nBuilding a new heap from array: [30, 10, 50, 2, 25]")
	arr := []int{30, 10, 50, 2, 25}
	newHeap := minheap.InitMinHeap()
	newHeap.BuildHeap(arr)
	fmt.Println("New heap:", newHeap)

	// Extract all elements (they will come out in ascending order)
	fmt.Println("\nExtracting all elements (in ascending order):")
	for !newHeap.IsEmpty() {
		val, _ := newHeap.Extract()
		fmt.Print(val, " ")
	}
	fmt.Println()
}
