package queue

// Queue represents a queue data structure that follows FIFO (First In First Out) principle
type Queue struct {
	items []int
}

// Enqueue adds an item to the end of the queue
// Time Complexity: O(1) - constant time operation (amortized)
// Parameters:
//   - item: The integer to be added to the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front item from the queue
// Time Complexity: O(n) - linear time operation due to slice shift
// Returns:
//   - int: The front item from the queue
//   - bool: True if the queue was not empty, false otherwise
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false // Return zero value and false for empty queue
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Front returns the front item without removing it
// Time Complexity: O(1) - constant time operation
// Returns:
//   - int: The front item from the queue
//   - bool: True if the queue was not empty, false otherwise
func (q *Queue) Front() (int, bool) {
	if len(q.items) == 0 {
		return 0, false // Return zero value and false for empty queue
	}

	return q.items[0], true
}

// IsEmpty checks if the queue is empty
// Time Complexity: O(1) - constant time operation
// Returns:
//   - bool: True if the queue is empty, false otherwise
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
// Time Complexity: O(1) - constant time operation
// Returns:
//   - int: The number of items in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

// Clear removes all items from the queue
// Time Complexity: O(1) - constant time operation
func (q *Queue) Clear() {
	q.items = []int{}
}
