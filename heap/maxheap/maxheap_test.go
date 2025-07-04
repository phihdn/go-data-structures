package maxheap

import (
	"reflect"
	"testing"
)

func TestInitMaxHeap(t *testing.T) {
	heap := InitMaxHeap()

	if heap == nil {
		t.Error("InitMaxHeap() returned nil")
	}

	if !heap.IsEmpty() {
		t.Error("New heap should be empty")
	}

	if heap.Size() != 0 {
		t.Errorf("Expected size 0, got %d", heap.Size())
	}
}

func TestInsert(t *testing.T) {
	heap := InitMaxHeap()
	testValues := []int{10, 20, 5, 15, 30}

	// Insert values one by one
	for _, val := range testValues {
		heap.Insert(val)
	}

	// Check size
	if heap.Size() != len(testValues) {
		t.Errorf("Expected size %d, got %d", len(testValues), heap.Size())
	}

	// Check max value
	max, ok := heap.GetMax()
	if !ok || max != 30 {
		t.Errorf("Expected max value 30, got %d", max)
	}
}

func TestExtract(t *testing.T) {
	heap := InitMaxHeap()
	testValues := []int{10, 20, 5, 15, 30}

	// Insert all values
	for _, val := range testValues {
		heap.Insert(val)
	}

	// Expected extraction order (in descending order)
	expected := []int{30, 20, 15, 10, 5}

	// Extract and verify all values
	for i, expectedVal := range expected {
		val, ok := heap.Extract()
		if !ok {
			t.Errorf("Extract() failed on iteration %d", i)
		}
		if val != expectedVal {
			t.Errorf("Expected %d, got %d", expectedVal, val)
		}
	}

	// Heap should be empty after extracting all values
	if !heap.IsEmpty() {
		t.Error("Heap should be empty after extracting all values")
	}

	// Extract from empty heap
	_, ok := heap.Extract()
	if ok {
		t.Error("Extract from empty heap should return false")
	}
}

func TestGetMax(t *testing.T) {
	heap := InitMaxHeap()

	// Empty heap
	_, ok := heap.GetMax()
	if ok {
		t.Error("GetMax on empty heap should return false")
	}

	// Insert values
	heap.Insert(5)
	max, ok := heap.GetMax()
	if !ok || max != 5 {
		t.Errorf("Expected max 5, got %d", max)
	}

	heap.Insert(10)
	max, ok = heap.GetMax()
	if !ok || max != 10 {
		t.Errorf("Expected max 10, got %d", max)
	}

	// Make sure GetMax doesn't modify the heap
	if heap.Size() != 2 {
		t.Errorf("GetMax should not modify heap size, expected 2, got %d", heap.Size())
	}
}

func TestBuildHeap(t *testing.T) {
	heap := InitMaxHeap()
	inputArray := []int{10, 20, 5, 15, 30}

	heap.BuildHeap(inputArray)

	// Check size
	if heap.Size() != len(inputArray) {
		t.Errorf("Expected size %d, got %d", len(inputArray), heap.Size())
	}

	// Check max value
	max, ok := heap.GetMax()
	if !ok || max != 30 {
		t.Errorf("Expected max value 30, got %d", max)
	}

	// Extract all and verify heap property
	var prev int
	for i := 0; i < len(inputArray); i++ {
		current, ok := heap.Extract()
		if !ok {
			t.Errorf("Failed to extract at index %d", i)
			continue
		}
		if i > 0 && current > prev {
			t.Errorf("Heap property violated: %d > %d", current, prev)
		}
		prev = current
	}
}

func TestHeapifyUpAndDown(t *testing.T) {
	// This test indirectly tests heapifyUp and heapifyDown through Insert and Extract
	heap := InitMaxHeap()

	// Insert in ascending order, which would trigger heapifyUp
	for i := 1; i <= 10; i++ {
		heap.Insert(i)
		max, _ := heap.GetMax()
		if max != i {
			t.Errorf("After inserting %d, max should be %d but got %d", i, i, max)
		}
	}

	// Extract all, which would trigger heapifyDown
	for i := 10; i >= 1; i-- {
		val, ok := heap.Extract()
		if !ok || val != i {
			t.Errorf("Expected to extract %d, got %d", i, val)
		}
	}
}

// TestInsertTableDriven uses a table-driven approach to test various insertion scenarios
func TestInsertTableDriven(t *testing.T) {
	tests := []struct {
		name          string
		inputSequence []int
		expectedMax   int
		expectedSize  int
	}{
		{
			name:          "Empty to single element",
			inputSequence: []int{5},
			expectedMax:   5,
			expectedSize:  1,
		},
		{
			name:          "Ascending order inserts",
			inputSequence: []int{1, 2, 3, 4, 5},
			expectedMax:   5,
			expectedSize:  5,
		},
		{
			name:          "Descending order inserts",
			inputSequence: []int{5, 4, 3, 2, 1},
			expectedMax:   5,
			expectedSize:  5,
		},
		{
			name:          "Mixed order inserts",
			inputSequence: []int{8, 3, 10, 1, 6, 14, 4, 7, 13},
			expectedMax:   14,
			expectedSize:  9,
		},
		{
			name:          "Duplicate values",
			inputSequence: []int{5, 5, 8, 8, 3, 3},
			expectedMax:   8,
			expectedSize:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := InitMaxHeap()

			// Insert all values in the sequence
			for _, val := range tt.inputSequence {
				heap.Insert(val)
			}

			// Check size
			if got := heap.Size(); got != tt.expectedSize {
				t.Errorf("Size() = %v, want %v", got, tt.expectedSize)
			}

			// Check max value
			max, ok := heap.GetMax()
			if !ok || max != tt.expectedMax {
				t.Errorf("GetMax() = (%v, %v), want (%v, true)", max, ok, tt.expectedMax)
			}

			// Check that the heap property is maintained
			got := heap.GetArray()
			if !isMaxHeapProperty(got) {
				t.Errorf("Heap property violated after inserts: %v", got)
			}
		})
	}
}

// TestExtractTableDriven uses a table-driven approach to test various extraction scenarios
func TestExtractTableDriven(t *testing.T) {
	tests := []struct {
		name               string
		initialValues      []int
		extractCount       int
		expectedSequence   []int
		expectedRemaining  int
		expectedFinalArray []int // expected final state of array after extractions
	}{
		{
			name:              "Extract from empty heap",
			initialValues:     []int{},
			extractCount:      1,
			expectedSequence:  []int{0}, // value doesn't matter, only the success status
			expectedRemaining: 0,
		},
		{
			name:               "Extract single element",
			initialValues:      []int{42},
			extractCount:       1,
			expectedSequence:   []int{42},
			expectedRemaining:  0,
			expectedFinalArray: []int{},
		},
		{
			name:               "Extract all elements in order",
			initialValues:      []int{5, 10, 3, 8, 15},
			extractCount:       5,
			expectedSequence:   []int{15, 10, 8, 5, 3},
			expectedRemaining:  0,
			expectedFinalArray: []int{},
		},
		{
			name:               "Extract some elements",
			initialValues:      []int{5, 10, 3, 8, 15},
			extractCount:       3,
			expectedSequence:   []int{15, 10, 8},
			expectedRemaining:  2,
			expectedFinalArray: []int{5, 3}, // specific layout depends on extract implementation
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := InitMaxHeap()

			// Build the initial heap
			for _, val := range tt.initialValues {
				heap.Insert(val)
			}

			// Extract the specified number of elements
			results := make([]int, 0, tt.extractCount)
			success := make([]bool, 0, tt.extractCount)

			for i := 0; i < tt.extractCount; i++ {
				val, ok := heap.Extract()
				results = append(results, val)
				success = append(success, ok)
			}

			// Check extraction sequence
			for i, expected := range tt.expectedSequence {
				// For empty heap case, we only check success status
				if i == 0 && len(tt.initialValues) == 0 {
					if success[i] {
						t.Errorf("Extract() from empty heap returned success=true, want false")
					}
				} else if i < len(tt.initialValues) {
					if !success[i] || results[i] != expected {
						t.Errorf("Extract()[%d] = (%v, %v), want (%v, true)",
							i, results[i], success[i], expected)
					}
				} else {
					if success[i] {
						t.Errorf("Extract()[%d] after heap emptied returned success=true, want false", i)
					}
				}
			}

			// Check remaining size
			if got := heap.Size(); got != tt.expectedRemaining {
				t.Errorf("Size() after extractions = %v, want %v", got, tt.expectedRemaining)
			}

			// Check final array state if specified
			if tt.expectedFinalArray != nil {
				got := heap.GetArray()
				if !reflect.DeepEqual(got, tt.expectedFinalArray) {
					t.Errorf("Final heap array = %v, want %v", got, tt.expectedFinalArray)
				}
			}
		})
	}
}

// TestBuildHeapTableDriven uses a table-driven approach to test BuildHeap functionality
func TestBuildHeapTableDriven(t *testing.T) {
	tests := []struct {
		name          string
		inputArray    []int
		expectedMax   int
		expectedArray []int // expected state after BuildHeap
	}{
		{
			name:          "Empty array",
			inputArray:    []int{},
			expectedMax:   0, // will return false on GetMax
			expectedArray: []int{},
		},
		{
			name:          "Single element",
			inputArray:    []int{42},
			expectedMax:   42,
			expectedArray: []int{42},
		},
		{
			name:          "Already heapified",
			inputArray:    []int{15, 10, 8, 5, 3},
			expectedMax:   15,
			expectedArray: []int{15, 10, 8, 5, 3},
		},
		{
			name:          "Reverse sorted",
			inputArray:    []int{1, 2, 3, 4, 5},
			expectedMax:   5,
			expectedArray: []int{5, 4, 3, 1, 2}, // specific layout depends on BuildHeap implementation
		},
		{
			name:          "Random order",
			inputArray:    []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7},
			expectedMax:   16,
			expectedArray: []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}, // specific layout depends on BuildHeap implementation
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := InitMaxHeap()

			// Build the heap from the input array
			heap.BuildHeap(tt.inputArray)

			// Check size
			if got := heap.Size(); got != len(tt.inputArray) {
				t.Errorf("Size() = %v, want %v", got, len(tt.inputArray))
			}

			// Check max value (for non-empty heaps)
			if len(tt.inputArray) > 0 {
				max, ok := heap.GetMax()
				if !ok || max != tt.expectedMax {
					t.Errorf("GetMax() = (%v, %v), want (%v, true)", max, ok, tt.expectedMax)
				}
			} else {
				_, ok := heap.GetMax()
				if ok {
					t.Errorf("GetMax() on empty heap returned ok=true, want false")
				}
			}

			// Check if the resulting array maintains heap property
			isHeapPropertyMaintained := true
			heapArray := heap.GetArray()

			for i := 0; i < len(heapArray); i++ {
				left := 2*i + 1
				right := 2*i + 2

				if left < len(heapArray) && heapArray[i] < heapArray[left] {
					isHeapPropertyMaintained = false
					break
				}

				if right < len(heapArray) && heapArray[i] < heapArray[right] {
					isHeapPropertyMaintained = false
					break
				}
			}

			if !isHeapPropertyMaintained {
				t.Errorf("Heap property not maintained after BuildHeap: %v", heapArray)
			}

			// Check resulting array if expected array is specified
			if len(tt.expectedArray) > 0 {
				got := heap.GetArray()
				if !reflect.DeepEqual(got, tt.expectedArray) {
					t.Errorf("Resulting heap array = %v, want %v", got, tt.expectedArray)
				}
			}
		})
	}
}

// TestHeapOperationsSequence tests a sequence of heap operations
func TestHeapOperationsSequence(t *testing.T) {
	tests := []struct {
		name           string
		operations     []string // sequence of operations to perform
		operationArgs  []int    // arguments for operations (0 for operations that don't need args)
		expectedValues []int    // expected return values from operations
		expectedOk     []bool   // expected success status from operations
		finalSize      int      // expected final size of heap
		finalArray     []int    // expected final array state
	}{
		{
			name:           "Insert then extract all",
			operations:     []string{"insert", "insert", "insert", "getmax", "extract", "extract", "extract", "extract"},
			operationArgs:  []int{5, 10, 3, 0, 0, 0, 0, 0},
			expectedValues: []int{0, 0, 0, 10, 10, 5, 3, 0},
			expectedOk:     []bool{true, true, true, true, true, true, true, false},
			finalSize:      0,
			finalArray:     []int{},
		},
		{
			name:           "Mixed operations",
			operations:     []string{"insert", "getmax", "insert", "getmax", "extract", "insert", "buildheap", "getmax", "extract"},
			operationArgs:  []int{7, 0, 15, 0, 0, 3, 0, 0, 0}, // buildheap arg is ignored, uses [10, 20, 5]
			expectedValues: []int{0, 7, 0, 15, 15, 0, 0, 20, 20},
			expectedOk:     []bool{true, true, true, true, true, true, true, true, true},
			finalSize:      2,
			finalArray:     []int{10, 5}, // after building [10, 20, 5] and extracting 20
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := InitMaxHeap()
			buildHeapValues := []int{10, 20, 5} // values for buildheap operation

			for i, op := range tt.operations {
				switch op {
				case "insert":
					heap.Insert(tt.operationArgs[i])
					if tt.expectedOk[i] == false {
						t.Errorf("Operation %d (%s): expected to fail but succeeded", i, op)
					}
				case "extract":
					val, ok := heap.Extract()
					if ok != tt.expectedOk[i] || (ok && val != tt.expectedValues[i]) {
						t.Errorf("Operation %d (%s): got (%v, %v), want (%v, %v)",
							i, op, val, ok, tt.expectedValues[i], tt.expectedOk[i])
					}
				case "getmax":
					val, ok := heap.GetMax()
					if ok != tt.expectedOk[i] || (ok && val != tt.expectedValues[i]) {
						t.Errorf("Operation %d (%s): got (%v, %v), want (%v, %v)",
							i, op, val, ok, tt.expectedValues[i], tt.expectedOk[i])
					}
				case "buildheap":
					heap.BuildHeap(buildHeapValues)
					if tt.expectedOk[i] == false {
						t.Errorf("Operation %d (%s): expected to fail but succeeded", i, op)
					}
				}
			}

			// Check final size
			if got := heap.Size(); got != tt.finalSize {
				t.Errorf("Final heap size = %v, want %v", got, tt.finalSize)
			}

			// Check final array state
			if tt.finalArray != nil {
				got := heap.GetArray()
				if !reflect.DeepEqual(got, tt.finalArray) {
					t.Errorf("Final heap array = %v, want %v", got, tt.finalArray)
				}
			}
		})
	}
}

// TestHeapPropertyAfterOperations tests that the heap property is maintained after sequences of operations
func TestHeapPropertyAfterOperations(t *testing.T) {
	tests := []struct {
		name      string
		insertSeq []int
		extractN  int
	}{
		{
			name:      "Small sequence",
			insertSeq: []int{10, 20, 5, 15, 30},
			extractN:  2,
		},
		{
			name:      "Large sequence with duplicates",
			insertSeq: []int{8, 3, 10, 1, 6, 14, 4, 7, 13, 8, 10, 4, 6},
			extractN:  5,
		},
		{
			name:      "Reverse ordered sequence",
			insertSeq: []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10},
			extractN:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := InitMaxHeap()

			// Insert all values in the sequence
			for _, val := range tt.insertSeq {
				heap.Insert(val)

				// Verify heap property after each insert
				if !isMaxHeapProperty(heap.GetArray()) {
					t.Errorf("Heap property violated after inserting %d: %v", val, heap.GetArray())
				}
			}

			// Extract values and check heap property
			for i := 0; i < tt.extractN && !heap.IsEmpty(); i++ {
				_, _ = heap.Extract()

				// Verify heap property after each extract
				if !isMaxHeapProperty(heap.GetArray()) {
					t.Errorf("Heap property violated after extract #%d: %v", i+1, heap.GetArray())
				}
			}
		})
	}
}

// Helper function to check if array maintains the max heap property
func isMaxHeapProperty(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		left := 2*i + 1
		right := 2*i + 2

		if left < len(arr) && arr[i] < arr[left] {
			return false
		}

		if right < len(arr) && arr[i] < arr[right] {
			return false
		}
	}
	return true
}
