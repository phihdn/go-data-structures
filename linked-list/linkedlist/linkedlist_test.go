package linkedlist

import (
	"testing"
)

// TestNewLinkedList tests the creation of a new empty linked list
func TestNewLinkedList(t *testing.T) {
	tests := []struct {
		name         string
		expectedHead *Node
		expectedLen  int
	}{
		{
			name:         "New empty linked list",
			expectedHead: nil,
			expectedLen:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := LinkedList{}

			if list.Head != tt.expectedHead {
				t.Errorf("New list head = %v, want %v", list.Head, tt.expectedHead)
			}

			if list.Length != tt.expectedLen {
				t.Errorf("New list length = %d, want %d", list.Length, tt.expectedLen)
			}
		})
	}
}

// TestPrepend tests adding nodes to the beginning of the list
func TestPrepend(t *testing.T) {
	tests := []struct {
		name       string
		operations []int // Values to prepend in sequence
		wantLen    int   // Expected final length
		wantOrder  []int // Expected order of values in list
	}{
		{
			name:       "Prepend to empty list",
			operations: []int{10},
			wantLen:    1,
			wantOrder:  []int{10},
		},
		{
			name:       "Multiple prepends",
			operations: []int{10, 20},
			wantLen:    2,
			wantOrder:  []int{20, 10},
		},
		{
			name:       "Series of prepends",
			operations: []int{5, 10, 15, 20},
			wantLen:    4,
			wantOrder:  []int{20, 15, 10, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := LinkedList{}

			// Perform the operations
			for _, val := range tt.operations {
				list.Prepend(&Node{Data: val})
			}

			// Check the length
			if list.Length != tt.wantLen {
				t.Errorf("List length = %d, want %d", list.Length, tt.wantLen)
			}

			// Check the order of values
			current := list.Head
			for i, expected := range tt.wantOrder {
				if current == nil {
					t.Errorf("List ended prematurely at position %d", i)
					break
				}

				if current.Data != expected {
					t.Errorf("Node at position %d = %d, want %d", i, current.Data, expected)
				}

				current = current.Next
			}

			// Make sure there are no extra nodes
			if current != nil {
				t.Errorf("List has extra nodes beyond expected length")
			}
		})
	}
}

// TestDeleteWithValue tests the deletion of nodes from the list
func TestDeleteWithValue(t *testing.T) {
	tests := []struct {
		name          string
		initialValues []int  // Initial values to build list (in reverse order due to prepend)
		deleteValue   int    // Value to delete
		wantLen       int    // Expected length after deletion
		wantOrder     []int  // Expected order of values after deletion
		description   string // Description of the test case
	}{
		{
			name:          "Delete from empty list",
			initialValues: []int{},
			deleteValue:   10,
			wantLen:       0,
			wantOrder:     []int{},
			description:   "Deleting from an empty list should not change the list",
		},
		{
			name:          "Delete head node",
			initialValues: []int{10, 20, 30}, // This will create 30->20->10
			deleteValue:   30,
			wantLen:       2,
			wantOrder:     []int{20, 10},
			description:   "Deleting the head node should update the head pointer",
		},
		{
			name:          "Delete middle node",
			initialValues: []int{10, 20, 30}, // This will create 30->20->10
			deleteValue:   20,
			wantLen:       2,
			wantOrder:     []int{30, 10},
			description:   "Deleting a middle node should connect the previous and next nodes",
		},
		{
			name:          "Delete tail node",
			initialValues: []int{10, 20, 30}, // This will create 30->20->10
			deleteValue:   10,
			wantLen:       2,
			wantOrder:     []int{30, 20},
			description:   "Deleting the tail node should update the last node's Next pointer",
		},
		{
			name:          "Delete non-existent value",
			initialValues: []int{10, 20}, // This will create 20->10
			deleteValue:   30,
			wantLen:       2,
			wantOrder:     []int{20, 10},
			description:   "Deleting a non-existent value should not change the list",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := LinkedList{}

			// Build the initial list
			for _, val := range tt.initialValues {
				list.Prepend(&Node{Data: val})
			}

			// Perform the deletion
			list.DeleteWithValue(tt.deleteValue)

			// Check the length
			if list.Length != tt.wantLen {
				t.Errorf("Length after deletion = %d, want %d", list.Length, tt.wantLen)
			}

			// Check the values
			current := list.Head
			for i, expected := range tt.wantOrder {
				if i >= tt.wantLen {
					break
				}

				if current == nil {
					t.Errorf("List ended prematurely at position %d", i)
					break
				}

				if current.Data != expected {
					t.Errorf("Node at position %d = %d, want %d", i, current.Data, expected)
				}

				current = current.Next
			}

			// Make sure there are no extra nodes
			if current != nil && len(tt.wantOrder) > 0 {
				t.Errorf("List has extra nodes beyond expected length")
			}
		})
	}
}

// TestMultipleOperations tests a sequence of operations to ensure they work together correctly
func TestMultipleOperations(t *testing.T) {
	tests := []struct {
		name          string
		initialValues []int // Initial values to build list (in reverse order due to prepend)
		deleteValue   int   // Value to delete
		wantLen       int   // Expected final length
		wantValues    []int // Expected values after operations
	}{
		{
			name:          "Add nodes and delete middle",
			initialValues: []int{10, 20, 30, 40}, // This will create 40->30->20->10
			deleteValue:   30,
			wantLen:       3,
			wantValues:    []int{40, 20, 10},
		},
		{
			name:          "Add nodes and delete first",
			initialValues: []int{10, 20, 30, 40}, // This will create 40->30->20->10
			deleteValue:   40,
			wantLen:       3,
			wantValues:    []int{30, 20, 10},
		},
		{
			name:          "Add nodes and delete last",
			initialValues: []int{10, 20, 30, 40}, // This will create 40->30->20->10
			deleteValue:   10,
			wantLen:       3,
			wantValues:    []int{40, 30, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := LinkedList{}

			// Build the initial list
			for _, val := range tt.initialValues {
				list.Prepend(&Node{Data: val})
			}

			// Delete the specified value
			list.DeleteWithValue(tt.deleteValue)

			// Check the length
			if list.Length != tt.wantLen {
				t.Errorf("Length after operations = %d, want %d", list.Length, tt.wantLen)
			}

			// Check the values
			current := list.Head
			for i, expected := range tt.wantValues {
				if current == nil {
					t.Errorf("List ended prematurely at position %d", i)
					break
				}

				if current.Data != expected {
					t.Errorf("Node at position %d = %d, want %d", i, current.Data, expected)
				}

				current = current.Next
			}

			// Make sure there are no extra nodes
			if current != nil {
				t.Errorf("List has extra nodes beyond expected length")
			}
		})
	}
}

// TestDuplicateValues tests how the list handles duplicate values
func TestDuplicateValues(t *testing.T) {
	tests := []struct {
		name          string
		initialValues []int  // Initial values to build list (in reverse order due to prepend)
		deleteValue   int    // Value to delete
		wantLen       int    // Expected length after deletion
		wantValues    []int  // Expected values after deletion
		description   string // Description of the test case
	}{
		{
			name:          "Delete first occurrence of duplicate value",
			initialValues: []int{10, 20, 10, 20}, // This will create 20->10->20->10
			deleteValue:   10,
			wantLen:       3,
			wantValues:    []int{20, 20, 10},
			description:   "Should delete only the first occurrence of the duplicate value",
		},
		{
			name:          "Delete first occurrence of head duplicate",
			initialValues: []int{10, 20, 30, 30}, // This will create 30->30->20->10
			deleteValue:   30,
			wantLen:       3,
			wantValues:    []int{30, 20, 10},
			description:   "Should delete only the first occurrence when it's the head",
		},
		{
			name:          "Delete value with multiple occurrences",
			initialValues: []int{20, 20, 20}, // This will create 20->20->20
			deleteValue:   20,
			wantLen:       2,
			wantValues:    []int{20, 20},
			description:   "Should delete only the first occurrence when all values are the same",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := LinkedList{}

			// Build the initial list
			for _, val := range tt.initialValues {
				list.Prepend(&Node{Data: val})
			}

			// Delete the value
			list.DeleteWithValue(tt.deleteValue)

			// Check the length
			if list.Length != tt.wantLen {
				t.Errorf("Length after deletion = %d, want %d", list.Length, tt.wantLen)
			}

			// Check the values
			current := list.Head
			for i, expected := range tt.wantValues {
				if current == nil {
					t.Errorf("List ended prematurely at position %d", i)
					break
				}

				if current.Data != expected {
					t.Errorf("Node at position %d = %d, want %d", i, current.Data, expected)
				}

				current = current.Next
			}

			// Make sure there are no extra nodes
			if current != nil {
				t.Errorf("List has extra nodes beyond expected length")
			}
		})
	}
}

// TestEmptyListEdgeCases tests more edge cases with empty lists
func TestEmptyListEdgeCases(t *testing.T) {
	tests := []struct {
		name       string
		operations []struct {
			op    string // "add" or "delete"
			value int    // Value to add or delete
		}
		wantLen     int  // Expected final length
		wantHead    bool // Expected whether head is nil (false) or not (true)
		description string
	}{
		{
			name: "Multiple deletions on empty list",
			operations: []struct {
				op    string
				value int
			}{
				{op: "delete", value: 10},
				{op: "delete", value: 20},
				{op: "delete", value: 30},
			},
			wantLen:     0,
			wantHead:    false,
			description: "Multiple deletions on empty list should not change length or head",
		},
		{
			name: "Add and immediately delete same value",
			operations: []struct {
				op    string
				value int
			}{
				{op: "add", value: 42},
				{op: "delete", value: 42},
			},
			wantLen:     0,
			wantHead:    false,
			description: "Adding and then deleting the same value should result in empty list",
		},
		{
			name: "Add two values, delete one",
			operations: []struct {
				op    string
				value int
			}{
				{op: "add", value: 42},
				{op: "add", value: 99},
				{op: "delete", value: 99},
			},
			wantLen:     1,
			wantHead:    true,
			description: "Adding two values and deleting one should leave one node",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := LinkedList{}

			// Perform operations
			for _, op := range tt.operations {
				switch op.op {
				case "add":
					list.Prepend(&Node{Data: op.value})
				case "delete":
					list.DeleteWithValue(op.value)
				}
			}

			// Check length
			if list.Length != tt.wantLen {
				t.Errorf("List length = %d, want %d after operations", list.Length, tt.wantLen)
			}

			// Check head
			headIsNil := list.Head == nil
			if tt.wantHead && headIsNil {
				t.Errorf("List head is nil, but expected a non-nil head")
			} else if !tt.wantHead && !headIsNil {
				t.Errorf("List head is not nil, but expected a nil head")
			}
		})
	}
}
