package stack

// Stack represents a stack data structure that follows LIFO (Last In First Out) principle
type Stack struct {
	items []int
}

// Push adds an item to the top of the stack
// Time Complexity: O(1) - constant time operation
// Parameters:
//   - item: The integer to be added to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
// Time Complexity: O(1) - constant time operation
// Returns:
//   - int: The top item from the stack
//   - bool: True if the stack was not empty, false otherwise
func (s *Stack) Pop() (int, bool) {
	len := len(s.items)
	if len == 0 {
		return 0, false // Return zero value and false for empty stack
	}

	item := s.items[len-1]
	s.items = s.items[:len-1]
	return item, true
}

// Peek returns the top item without removing it
// Time Complexity: O(1) - constant time operation
// Returns:
//   - int: The top item from the stack
//   - bool: True if the stack was not empty, false otherwise
func (s *Stack) Peek() (int, bool) {
	len := len(s.items)
	if len == 0 {
		return 0, false // Return zero value and false for empty stack
	}

	return s.items[len-1], true
}

// IsEmpty checks if the stack is empty
// Time Complexity: O(1) - constant time operation
// Returns:
//   - bool: True if the stack is empty, false otherwise
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
// Time Complexity: O(1) - constant time operation
// Returns:
//   - int: The number of items in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// Clear removes all items from the stack
// Time Complexity: O(1) - constant time operation
func (s *Stack) Clear() {
	s.items = []int{}
}
