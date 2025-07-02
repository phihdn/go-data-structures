package queue

import (
	"testing"
)

// TestNewQueue tests the creation of a new empty queue
func TestNewQueue(t *testing.T) {
	q := Queue{}

	if !q.IsEmpty() {
		t.Errorf("New queue should be empty")
	}

	if q.Size() != 0 {
		t.Errorf("New queue size = %d, want 0", q.Size())
	}
}

// TestEnqueue tests adding items to the queue
func TestEnqueue(t *testing.T) {
	tests := []struct {
		name       string
		operations []int // Values to enqueue in sequence
		wantSize   int   // Expected final size
	}{
		{
			name:       "Enqueue to empty queue",
			operations: []int{10},
			wantSize:   1,
		},
		{
			name:       "Multiple enqueues",
			operations: []int{10, 20, 30},
			wantSize:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Queue{}

			// Perform the operations
			for _, val := range tt.operations {
				q.Enqueue(val)
			}

			// Check the size
			if q.Size() != tt.wantSize {
				t.Errorf("Queue size = %d, want %d", q.Size(), tt.wantSize)
			}
		})
	}
}

// TestDequeue tests removing items from the queue
func TestDequeue(t *testing.T) {
	tests := []struct {
		name          string
		setupValues   []int  // Values to enqueue before testing
		dequeueCount  int    // Number of times to dequeue
		wantValues    []int  // Expected values from dequeue operations
		wantSuccesses []bool // Expected success indicators from dequeue operations
		wantFinalSize int    // Expected final size after dequeues
	}{
		{
			name:          "Dequeue from empty queue",
			setupValues:   []int{},
			dequeueCount:  1,
			wantValues:    []int{0},
			wantSuccesses: []bool{false},
			wantFinalSize: 0,
		},
		{
			name:          "Enqueue then dequeue one item",
			setupValues:   []int{42},
			dequeueCount:  1,
			wantValues:    []int{42},
			wantSuccesses: []bool{true},
			wantFinalSize: 0,
		},
		{
			name:          "Multiple enqueues and dequeues",
			setupValues:   []int{10, 20, 30},
			dequeueCount:  2,
			wantValues:    []int{10, 20},
			wantSuccesses: []bool{true, true},
			wantFinalSize: 1,
		},
		{
			name:          "Dequeue until empty and beyond",
			setupValues:   []int{5, 10},
			dequeueCount:  3,
			wantValues:    []int{5, 10, 0},
			wantSuccesses: []bool{true, true, false},
			wantFinalSize: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Queue{}

			// Setup the queue
			for _, val := range tt.setupValues {
				q.Enqueue(val)
			}

			// Perform dequeue operations and check results
			for i := 0; i < tt.dequeueCount; i++ {
				val, ok := q.Dequeue()

				if i < len(tt.wantValues) && val != tt.wantValues[i] {
					t.Errorf("Dequeue %d value = %d, want %d", i, val, tt.wantValues[i])
				}

				if i < len(tt.wantSuccesses) && ok != tt.wantSuccesses[i] {
					t.Errorf("Dequeue %d success = %v, want %v", i, ok, tt.wantSuccesses[i])
				}
			}

			// Check final size
			if q.Size() != tt.wantFinalSize {
				t.Errorf("Final queue size = %d, want %d", q.Size(), tt.wantFinalSize)
			}
		})
	}
}

// TestFront tests viewing the front item without removing it
func TestFront(t *testing.T) {
	tests := []struct {
		name        string
		setupValues []int // Values to enqueue before testing
		wantValue   int   // Expected value from front operation
		wantSuccess bool  // Expected success indicator from front operation
	}{
		{
			name:        "Front of empty queue",
			setupValues: []int{},
			wantValue:   0,
			wantSuccess: false,
		},
		{
			name:        "Front with one item",
			setupValues: []int{42},
			wantValue:   42,
			wantSuccess: true,
		},
		{
			name:        "Front with multiple items",
			setupValues: []int{10, 20, 30},
			wantValue:   10,
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Queue{}

			// Setup the queue
			for _, val := range tt.setupValues {
				q.Enqueue(val)
			}

			// Test front operation
			val, ok := q.Front()

			if val != tt.wantValue {
				t.Errorf("Front value = %d, want %d", val, tt.wantValue)
			}

			if ok != tt.wantSuccess {
				t.Errorf("Front success = %v, want %v", ok, tt.wantSuccess)
			}

			// Verify size hasn't changed after Front
			if len(q.items) != len(tt.setupValues) {
				t.Errorf("Queue size changed after Front: got %d, want %d", len(q.items), len(tt.setupValues))
			}
		})
	}
}

// TestClear tests removing all items from the queue
func TestClear(t *testing.T) {
	q := Queue{}

	// Add some items
	values := []int{10, 20, 30}
	for _, val := range values {
		q.Enqueue(val)
	}

	// Clear the queue
	q.Clear()

	// Verify queue is empty
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty after Clear()")
	}

	if q.Size() != 0 {
		t.Errorf("Queue size = %d after Clear(), want 0", q.Size())
	}
}
