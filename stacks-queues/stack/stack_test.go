package stack

import (
	"testing"
)

// TestNewStack tests the creation of a new empty stack
func TestNewStack(t *testing.T) {
	s := Stack{}

	if !s.IsEmpty() {
		t.Errorf("New stack should be empty")
	}

	if s.Size() != 0 {
		t.Errorf("New stack size = %d, want 0", s.Size())
	}
}

// TestPush tests adding items to the stack
func TestPush(t *testing.T) {
	tests := []struct {
		name       string
		operations []int // Values to push in sequence
		wantSize   int   // Expected final size
	}{
		{
			name:       "Push to empty stack",
			operations: []int{10},
			wantSize:   1,
		},
		{
			name:       "Multiple pushes",
			operations: []int{10, 20, 30},
			wantSize:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{}

			// Perform the operations
			for _, val := range tt.operations {
				s.Push(val)
			}

			// Check the size
			if s.Size() != tt.wantSize {
				t.Errorf("Stack size = %d, want %d", s.Size(), tt.wantSize)
			}
		})
	}
}

// TestPop tests removing items from the stack
func TestPop(t *testing.T) {
	tests := []struct {
		name          string
		setupValues   []int  // Values to push before testing
		popCount      int    // Number of times to pop
		wantValues    []int  // Expected values from pop operations
		wantSuccesses []bool // Expected success indicators from pop operations
		wantFinalSize int    // Expected final size after pops
	}{
		{
			name:          "Pop from empty stack",
			setupValues:   []int{},
			popCount:      1,
			wantValues:    []int{0},
			wantSuccesses: []bool{false},
			wantFinalSize: 0,
		},
		{
			name:          "Push then pop one item",
			setupValues:   []int{42},
			popCount:      1,
			wantValues:    []int{42},
			wantSuccesses: []bool{true},
			wantFinalSize: 0,
		},
		{
			name:          "Multiple pushes and pops",
			setupValues:   []int{10, 20, 30},
			popCount:      2,
			wantValues:    []int{30, 20},
			wantSuccesses: []bool{true, true},
			wantFinalSize: 1,
		},
		{
			name:          "Pop until empty and beyond",
			setupValues:   []int{5, 10},
			popCount:      3,
			wantValues:    []int{10, 5, 0},
			wantSuccesses: []bool{true, true, false},
			wantFinalSize: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{}

			// Setup the stack
			for _, val := range tt.setupValues {
				s.Push(val)
			}

			// Perform pop operations and check results
			for i := 0; i < tt.popCount; i++ {
				val, ok := s.Pop()

				if i < len(tt.wantValues) && val != tt.wantValues[i] {
					t.Errorf("Pop %d value = %d, want %d", i, val, tt.wantValues[i])
				}

				if i < len(tt.wantSuccesses) && ok != tt.wantSuccesses[i] {
					t.Errorf("Pop %d success = %v, want %v", i, ok, tt.wantSuccesses[i])
				}
			}

			// Check final size
			if s.Size() != tt.wantFinalSize {
				t.Errorf("Final stack size = %d, want %d", s.Size(), tt.wantFinalSize)
			}
		})
	}
}

// TestPeek tests viewing the top item without removing it
func TestPeek(t *testing.T) {
	tests := []struct {
		name        string
		setupValues []int // Values to push before testing
		wantValue   int   // Expected value from peek operation
		wantSuccess bool  // Expected success indicator from peek operation
	}{
		{
			name:        "Peek empty stack",
			setupValues: []int{},
			wantValue:   0,
			wantSuccess: false,
		},
		{
			name:        "Peek with one item",
			setupValues: []int{42},
			wantValue:   42,
			wantSuccess: true,
		},
		{
			name:        "Peek with multiple items",
			setupValues: []int{10, 20, 30},
			wantValue:   30,
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{}

			// Setup the stack
			for _, val := range tt.setupValues {
				s.Push(val)
			}

			// Test peek operation
			val, ok := s.Peek()

			if val != tt.wantValue {
				t.Errorf("Peek value = %d, want %d", val, tt.wantValue)
			}

			if ok != tt.wantSuccess {
				t.Errorf("Peek success = %v, want %v", ok, tt.wantSuccess)
			}

			// Verify size hasn't changed after peek
			if len(s.items) != len(tt.setupValues) {
				t.Errorf("Stack size changed after Peek: got %d, want %d", len(s.items), len(tt.setupValues))
			}
		})
	}
}

// TestClear tests removing all items from the stack
func TestClear(t *testing.T) {
	s := Stack{}

	// Add some items
	values := []int{10, 20, 30}
	for _, val := range values {
		s.Push(val)
	}

	// Clear the stack
	s.Clear()

	// Verify stack is empty
	if !s.IsEmpty() {
		t.Errorf("Stack should be empty after Clear()")
	}

	if s.Size() != 0 {
		t.Errorf("Stack size = %d after Clear(), want 0", s.Size())
	}
}
