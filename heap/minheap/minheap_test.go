package minheap

import (
	"testing"
)

func TestInitMinHeap(t *testing.T) {
	h := InitMinHeap()
	if h.Size() != 0 {
		t.Errorf("Expected empty heap, got size %d", h.Size())
	}
	if !h.IsEmpty() {
		t.Error("Expected IsEmpty to return true for new heap")
	}
}

func TestInsert(t *testing.T) {
	h := InitMinHeap()
	h.Insert(10)
	if h.Size() != 1 {
		t.Errorf("Expected size 1, got %d", h.Size())
	}

	min, ok := h.GetMin()
	if !ok || min != 10 {
		t.Errorf("Expected min value 10, got %d, ok: %v", min, ok)
	}

	// Insert more values
	h.Insert(5)
	h.Insert(15)
	h.Insert(2)
	h.Insert(20)

	// Check if min value is correctly maintained
	min, ok = h.GetMin()
	if !ok || min != 2 {
		t.Errorf("Expected min value 2, got %d, ok: %v", min, ok)
	}

	// Check size
	if h.Size() != 5 {
		t.Errorf("Expected size 5, got %d", h.Size())
	}
}

func TestExtract(t *testing.T) {
	h := InitMinHeap()

	// Test extract on empty heap
	_, ok := h.Extract()
	if ok {
		t.Error("Expected Extract to fail on empty heap")
	}

	// Insert values
	values := []int{10, 20, 5, 15, 30, 2}
	for _, v := range values {
		h.Insert(v)
	}

	// Extract values and verify they come out in ascending order
	expected := []int{2, 5, 10, 15, 20, 30}
	for _, exp := range expected {
		val, ok := h.Extract()
		if !ok {
			t.Error("Extract failed unexpectedly")
		}
		if val != exp {
			t.Errorf("Expected to extract %d, got %d", exp, val)
		}
	}

	// Verify heap is empty after all extractions
	if !h.IsEmpty() {
		t.Error("Expected heap to be empty after all extractions")
	}
}

func TestBuildHeap(t *testing.T) {
	h := InitMinHeap()
	input := []int{10, 20, 5, 15, 30, 2}
	h.BuildHeap(input)

	// Verify size
	if h.Size() != len(input) {
		t.Errorf("Expected size %d, got %d", len(input), h.Size())
	}

	// Verify min value
	min, ok := h.GetMin()
	if !ok || min != 2 {
		t.Errorf("Expected min value 2, got %d, ok: %v", min, ok)
	}

	// Extract values and verify they come out in ascending order
	expected := []int{2, 5, 10, 15, 20, 30}
	for _, exp := range expected {
		val, ok := h.Extract()
		if !ok {
			t.Error("Extract failed unexpectedly")
		}
		if val != exp {
			t.Errorf("Expected to extract %d, got %d", exp, val)
		}
	}
}

func TestGetArray(t *testing.T) {
	h := InitMinHeap()
	input := []int{10, 5, 20, 30, 15}
	for _, v := range input {
		h.Insert(v)
	}

	arr := h.GetArray()

	// Check that the returned array does not reference the internal array
	arr[0] = 999
	min, _ := h.GetMin()
	if min == 999 {
		t.Error("GetArray should return a copy, not a reference")
	}

	// Extract all values to ensure heap property is maintained
	values := []int{}
	for !h.IsEmpty() {
		val, _ := h.Extract()
		values = append(values, val)
	}

	// Verify values are in ascending order
	for i := 1; i < len(values); i++ {
		if values[i] < values[i-1] {
			t.Errorf("Heap property violated: %v not in ascending order", values)
		}
	}
}
