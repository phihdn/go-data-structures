package maxheap

import "fmt"

// MaxHeap represents a max heap data structure
// A max heap is a complete binary tree where the value of each node is greater than or equal to
// the values of its children, making the root node the maximum value in the heap.
// The heap is implemented using an array where:
// - The root is at index 0
// - For a node at index i:
//   - Its parent is at index (i-1)/2
//   - Its left child is at index 2*i + 1
//   - Its right child is at index 2*i + 2
type MaxHeap struct {
	array []int // Slice that stores the heap elements
}

// Insert adds a value to the heap and maintains the heap property
// Time complexity: O(log n) where n is the number of elements in the heap
// The operation requires potentially bubbling up the new element from the bottom to its correct position
func (h *MaxHeap) Insert(value int) {
	// Add the new value to the end of the array
	h.array = append(h.array, value)
	// Restore heap property by bubbling up the new value to its correct position
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract removes and returns the maximum value from the heap
// Returns the extracted value and a boolean indicating success
// Time complexity: O(log n) where n is the number of elements in the heap
// This operation removes the root (maximum element) and restores the heap property
func (h *MaxHeap) Extract() (int, bool) {
	if len(h.array) == 0 {
		fmt.Println("Heap is empty")
		return 0, false
	}

	// The maximum value in a max heap is always at the root (index 0)
	extracted := h.array[0]
	lastIndex := len(h.array) - 1

	// Replace the root with the last element in the heap
	h.array[0] = h.array[lastIndex]
	// Remove the last element (which is now duplicated at the root)
	h.array = h.array[:lastIndex]

	// Restore the heap property by bubbling down the new root to its correct position
	// Only need to do this if the heap isn't empty after extraction
	if len(h.array) > 0 {
		h.maxHeapifyDown(0)
	}

	return extracted, true
}

// GetMax returns the maximum value without extracting it
// Returns the max value and a boolean indicating success
// Time complexity: O(1) since the maximum is always at the root
func (h *MaxHeap) GetMax() (int, bool) {
	if len(h.array) == 0 {
		return 0, false // Return zero value and false for empty heap
	}
	return h.array[0], true // Return root element (max value in a max heap)
}

// Size returns the number of elements in the heap
// Time complexity: O(1)
func (h *MaxHeap) Size() int {
	return len(h.array)
}

// IsEmpty returns true if the heap has no elements
// Time complexity: O(1)
func (h *MaxHeap) IsEmpty() bool {
	return len(h.array) == 0
}

// GetArray returns a copy of the underlying array
// This is useful when you need to access the heap elements without modifying the heap
// Returns a new array to prevent modification of the internal heap structure
// Time complexity: O(n) where n is the number of elements
func (h *MaxHeap) GetArray() []int {
	result := make([]int, len(h.array))
	copy(result, h.array) // Make a deep copy to avoid external modifications
	return result
}

// maxHeapifyUp maintains the heap property going upward from a node
// Used during insertion to position a new element correctly
// It compares a node with its parent and swaps them if the heap property is violated
// Time complexity: O(log n) in the worst case, where n is the number of elements
func (h *MaxHeap) maxHeapifyUp(index int) {
	// Continue until we reach the root (index 0) or the heap property is restored
	// The heap property is violated if the current node is greater than its parent
	for h.array[parent(index)] < h.array[index] && index > 0 {
		// Swap the current node with its parent
		h.swap(index, parent(index))
		// Move up to the parent index and continue the process
		index = parent(index)
	}
}

// maxHeapifyDown maintains the heap property going downward from a node
// Used during extraction to position the root element correctly
// It compares a node with its children and swaps it with the larger child if needed
// Time complexity: O(log n) in the worst case, where n is the number of elements
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index) // Get the indices of left and right children
	childToCompare := 0               // Will store the index of the larger child

	// Continue as long as the left child exists (if left doesn't exist, right won't either)
	for l <= lastIndex {
		// Determine which child to compare with the current node
		if l == lastIndex {
			// Case 1: Left child is the only child (right child doesn't exist)
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			// Case 2: Left child is greater than right child
			childToCompare = l
		} else {
			// Case 3: Right child is greater than or equal to left child
			childToCompare = r
		}

		// If the current node is smaller than the larger child, swap them and continue down
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			// Move down to the child position and update indices for the next iteration
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			// The heap property is satisfied, no further adjustments needed
			return
		}
	}
}

// BuildHeap constructs a heap from an array in O(n) time
// This is more efficient than inserting elements one by one (which would be O(n log n))
// The algorithm works by starting from the first non-leaf node and performing maxHeapifyDown
// for each node up to the root
func (h *MaxHeap) BuildHeap(arr []int) {
	// Make a copy of the input array
	h.array = make([]int, len(arr))
	copy(h.array, arr)

	// Start heapify from the first non-leaf node and move up to the root
	// The index of the first non-leaf node is (n/2)-1 where n is the size of the array
	// All nodes after this index are leaf nodes and don't need heapifying down
	for i := (len(h.array) / 2) - 1; i >= 0; i-- {
		h.maxHeapifyDown(i)
	}
}

// parent returns the parent index of a node at index i
// Formula: (i-1)/2
// This formula works because in a binary heap:
// - If i is the index of a node, its parent is at floor((i-1)/2)
// - For a zero-based array representing a binary heap, this gives the correct parent index
func parent(i int) int {
	return (i - 1) / 2
}

// left returns the left child index of a node at index i
// Formula: 2*i + 1
// This formula works because in a binary heap:
// - If i is the index of a node, its left child is at 2*i + 1
// - The +1 accounts for the zero-based indexing of the array
func left(i int) int {
	return 2*i + 1
}

// right returns the right child index of a node at index i
// Formula: 2*i + 2
// This formula works because in a binary heap:
// - If i is the index of a node, its right child is at 2*i + 2
// - This is one position to the right of the left child (which is at 2*i + 1)
func right(i int) int {
	return 2*i + 2
}

// swap exchanges two elements in the heap
// This is a helper function used during heap operations to swap elements at positions i1 and i2
// Uses Go's tuple assignment feature for an efficient swap without a temporary variable
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// InitMaxHeap creates and initializes a new empty MaxHeap
// Returns a pointer to the newly created heap
// This is the recommended way to create a new heap instance
func InitMaxHeap() *MaxHeap {
	return &MaxHeap{array: []int{}}
}

// String returns a string representation of the heap
// This method implements the Stringer interface for better debugging and printing
// Time complexity: O(n) where n is the number of elements
func (h *MaxHeap) String() string {
	return fmt.Sprintf("MaxHeap{array: %v}", h.array)
}
