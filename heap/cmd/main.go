package main

import (
	"fmt"

	"github.com/phihdn/go-data-structures/heap/maxheap"
)

func main() {
	fmt.Println("=== Max Heap Demonstration ===")

	// Create a new max heap
	m := maxheap.InitMaxHeap()
	fmt.Println("Initialized empty heap:", m)

	// Insert elements one by one
	fmt.Println("\n1. Inserting elements one by one:")
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Printf("  After inserting %d: %v\n", v, m)
	}

	// Extract the max element multiple times
	fmt.Println("\n2. Extracting max elements:")
	for i := 0; i < 5; i++ {
		max, success := m.Extract()
		if success {
			fmt.Printf("  Extracted max: %d, Heap after extraction: %v\n", max, m)
		} else {
			fmt.Println("  Failed to extract (heap is empty)")
		}
	}

	// Build a heap from an array directly
	fmt.Println("\n3. Building a heap from an array:")
	newHeap := maxheap.InitMaxHeap()
	newArray := []int{4, 8, 2, 6, 10, 3, 1, 7, 5, 9}
	fmt.Printf("  Original array: %v\n", newArray)

	newHeap.BuildHeap(newArray)
	fmt.Printf("  Resulting heap: %v\n", newHeap)

	// Extract all elements to show they come out in sorted order
	fmt.Println("\n4. Extracting all elements (sorted in descending order):")
	var extracted []int
	for !newHeap.IsEmpty() {
		val, _ := newHeap.Extract()
		extracted = append(extracted, val)
	}
	fmt.Printf("  Extracted values: %v\n", extracted)

	fmt.Println("\n=== Max Heap Demonstration Complete ===")
}
