package bst

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := NewBST(100)

	// Test inserting values
	tree.Insert(50)
	tree.Insert(150)

	// Verify left child
	if tree.Left == nil || tree.Left.Key != 50 {
		t.Errorf("Expected left child with key 50, got %v", tree.Left)
	}

	// Verify right child
	if tree.Right == nil || tree.Right.Key != 150 {
		t.Errorf("Expected right child with key 150, got %v", tree.Right)
	}
}

func TestSearch(t *testing.T) {
	tree := NewBST(100)
	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(25)
	tree.Insert(75)
	tree.Insert(125)
	tree.Insert(175)

	// Test searching for existing values
	testCases := []struct {
		key      int
		expected bool
	}{
		{100, true},  // root
		{50, true},   // left child
		{150, true},  // right child
		{25, true},   // leaf node
		{175, true},  // leaf node
		{0, false},   // non-existent value
		{200, false}, // non-existent value
	}

	for _, tc := range testCases {
		if result := tree.Search(tc.key); result != tc.expected {
			t.Errorf("Search(%d) = %v, expected %v", tc.key, result, tc.expected)
		}
	}
}

func setupTestTree() *Node {
	/*
	   Tree structure:
	       100
	      /   \
	     50   150
	    / \   / \
	   25 75 125 175
	*/
	tree := NewBST(100)
	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(25)
	tree.Insert(75)
	tree.Insert(125)
	tree.Insert(175)
	return tree
}

func TestInOrderTraversal(t *testing.T) {
	tree := setupTestTree()
	expected := []int{25, 50, 75, 100, 125, 150, 175}
	result := tree.InOrderTraversal()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InOrderTraversal() = %v, expected %v", result, expected)
	}
}

func TestPreOrderTraversal(t *testing.T) {
	tree := setupTestTree()
	expected := []int{100, 50, 25, 75, 150, 125, 175}
	result := tree.PreOrderTraversal()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreOrderTraversal() = %v, expected %v", result, expected)
	}
}

func TestPostOrderTraversal(t *testing.T) {
	tree := setupTestTree()
	expected := []int{25, 75, 50, 125, 175, 150, 100}
	result := tree.PostOrderTraversal()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PostOrderTraversal() = %v, expected %v", result, expected)
	}
}

func TestMinMax(t *testing.T) {
	tree := setupTestTree()

	if min := tree.Min(); min != 25 {
		t.Errorf("Min() = %d, expected 25", min)
	}

	if max := tree.Max(); max != 175 {
		t.Errorf("Max() = %d, expected 175", max)
	}
}
