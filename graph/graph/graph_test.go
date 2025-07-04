package graph

import (
	"reflect"
	"testing"
)

func TestAddVertex(t *testing.T) {
	testCases := []struct {
		name        string
		key         int
		expectError bool
	}{
		{
			name:        "Add unique vertex",
			key:         1,
			expectError: false,
		},
		{
			name:        "Add duplicate vertex",
			key:         1,
			expectError: true,
		},
	}

	g := NewGraph()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := g.AddVertex(tc.key)
			if tc.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestAddEdge(t *testing.T) {
	testCases := []struct {
		name        string
		fromKey     int
		toKey       int
		setupGraph  func() *Graph
		expectError bool
	}{
		{
			name:    "Add edge between existing vertices",
			fromKey: 1,
			toKey:   2,
			setupGraph: func() *Graph {
				g := NewGraph()
				_ = g.AddVertex(1)
				_ = g.AddVertex(2)
				return g
			},
			expectError: false,
		},
		{
			name:    "Add edge from non-existent vertex",
			fromKey: 3,
			toKey:   2,
			setupGraph: func() *Graph {
				g := NewGraph()
				_ = g.AddVertex(1)
				_ = g.AddVertex(2)
				return g
			},
			expectError: true,
		},
		{
			name:    "Add edge to non-existent vertex",
			fromKey: 1,
			toKey:   3,
			setupGraph: func() *Graph {
				g := NewGraph()
				_ = g.AddVertex(1)
				_ = g.AddVertex(2)
				return g
			},
			expectError: true,
		},
		{
			name:    "Add duplicate edge",
			fromKey: 1,
			toKey:   2,
			setupGraph: func() *Graph {
				g := NewGraph()
				_ = g.AddVertex(1)
				_ = g.AddVertex(2)
				_ = g.AddEdge(1, 2)
				return g
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := tc.setupGraph()
			err := g.AddEdge(tc.fromKey, tc.toKey)
			if tc.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestHasVertex(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)

	testCases := []struct {
		name     string
		key      int
		expected bool
	}{
		{
			name:     "Has existing vertex",
			key:      1,
			expected: true,
		},
		{
			name:     "Doesn't have non-existent vertex",
			key:      3,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := g.HasVertex(tc.key); result != tc.expected {
				t.Errorf("HasVertex(%d) = %v, expected %v", tc.key, result, tc.expected)
			}
		})
	}
}

func TestHasEdge(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_ = g.AddVertex(3)
	_ = g.AddEdge(1, 2)
	_ = g.AddEdge(2, 3)

	testCases := []struct {
		name     string
		fromKey  int
		toKey    int
		expected bool
	}{
		{
			name:     "Has existing edge",
			fromKey:  1,
			toKey:    2,
			expected: true,
		},
		{
			name:     "Edge in reverse direction doesn't exist",
			fromKey:  2,
			toKey:    1,
			expected: false,
		},
		{
			name:     "Doesn't have non-existent edge",
			fromKey:  1,
			toKey:    3,
			expected: false,
		},
		{
			name:     "From vertex doesn't exist",
			fromKey:  4,
			toKey:    1,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := g.HasEdge(tc.fromKey, tc.toKey); result != tc.expected {
				t.Errorf("HasEdge(%d, %d) = %v, expected %v", tc.fromKey, tc.toKey, result, tc.expected)
			}
		})
	}
}

func TestGetNeighbors(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_ = g.AddVertex(3)
	_ = g.AddVertex(4)
	_ = g.AddEdge(1, 2)
	_ = g.AddEdge(1, 3)
	_ = g.AddEdge(2, 4)

	testCases := []struct {
		name         string
		key          int
		expectedKeys []int
		expectError  bool
	}{
		{
			name:         "Get neighbors of vertex with multiple edges",
			key:          1,
			expectedKeys: []int{2, 3},
			expectError:  false,
		},
		{
			name:         "Get neighbors of vertex with one edge",
			key:          2,
			expectedKeys: []int{4},
			expectError:  false,
		},
		{
			name:         "Get neighbors of vertex with no edges",
			key:          4,
			expectedKeys: []int{},
			expectError:  false,
		},
		{
			name:         "Get neighbors of non-existent vertex",
			key:          5,
			expectedKeys: nil,
			expectError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			neighbors, err := g.GetNeighbors(tc.key)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if !tc.expectError {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				} else if len(neighbors) == 0 && len(tc.expectedKeys) == 0 {
					// Both are empty, this is fine
				} else if !reflect.DeepEqual(neighbors, tc.expectedKeys) {
					t.Errorf("GetNeighbors(%d) = %v, expected %v", tc.key, neighbors, tc.expectedKeys)
				}
			}
		})
	}
}

func TestBFS(t *testing.T) {
	/*
	   Graph structure:
	   1 --> 2 --> 4
	   |     |
	   v     v
	   3 --> 5
	*/
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_ = g.AddVertex(3)
	_ = g.AddVertex(4)
	_ = g.AddVertex(5)

	_ = g.AddEdge(1, 2)
	_ = g.AddEdge(1, 3)
	_ = g.AddEdge(2, 4)
	_ = g.AddEdge(2, 5)
	_ = g.AddEdge(3, 5)

	testCases := []struct {
		name        string
		startKey    int
		expected    []int
		expectError bool
	}{
		{
			name:        "BFS from vertex 1",
			startKey:    1,
			expected:    []int{1, 2, 3, 4, 5},
			expectError: false,
		},
		{
			name:        "BFS from vertex 3",
			startKey:    3,
			expected:    []int{3, 5},
			expectError: false,
		},
		{
			name:        "BFS from non-existent vertex",
			startKey:    6,
			expected:    nil,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := g.BFS(tc.startKey)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if !tc.expectError {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				} else if !reflect.DeepEqual(result, tc.expected) {
					t.Errorf("BFS(%d) = %v, expected %v", tc.startKey, result, tc.expected)
				}
			}
		})
	}
}

func TestDFS(t *testing.T) {
	/*
	   Graph structure:
	   1 --> 2 --> 4
	   |     |
	   v     v
	   3 --> 5
	*/
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_ = g.AddVertex(3)
	_ = g.AddVertex(4)
	_ = g.AddVertex(5)

	_ = g.AddEdge(1, 2)
	_ = g.AddEdge(1, 3)
	_ = g.AddEdge(2, 4)
	_ = g.AddEdge(2, 5)
	_ = g.AddEdge(3, 5)

	testCases := []struct {
		name        string
		startKey    int
		expected    []int
		expectError bool
	}{
		{
			name:        "DFS from vertex 1",
			startKey:    1,
			expected:    []int{1, 2, 4, 5, 3},
			expectError: false,
		},
		{
			name:        "DFS from vertex 3",
			startKey:    3,
			expected:    []int{3, 5},
			expectError: false,
		},
		{
			name:        "DFS from non-existent vertex",
			startKey:    6,
			expected:    nil,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := g.DFS(tc.startKey)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if !tc.expectError {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				} else if !reflect.DeepEqual(result, tc.expected) {
					t.Errorf("DFS(%d) = %v, expected %v", tc.startKey, result, tc.expected)
				}
			}
		})
	}
}
