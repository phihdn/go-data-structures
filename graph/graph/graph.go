package graph

import "fmt"

// Vertex represents a node in the graph with a key and list of adjacent vertices
type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

// Graph represents a graph data structure with a collection of vertices
type Graph struct {
	Vertices []*Vertex
}

// NewGraph creates and returns a new empty graph
func NewGraph() *Graph {
	return &Graph{}
}

// AddVertex adds a new vertex with the given key to the graph
// If a vertex with the key already exists, an error is returned
func (g *Graph) AddVertex(key int) error {
	if containsVertex(g.Vertices, key) {
		return fmt.Errorf("vertex with key %d already exists", key)
	}
	v := &Vertex{Key: key}
	g.Vertices = append(g.Vertices, v)
	return nil
}

// AddEdge creates a directed edge from the vertex with fromKey to the vertex with toKey
// Returns an error if either vertex does not exist or if the edge already exists
func (g *Graph) AddEdge(fromKey, toKey int) error {
	fromVertex := g.getVertex(fromKey)
	toVertex := g.getVertex(toKey)

	if fromVertex == nil {
		return fmt.Errorf("from vertex with key %d does not exist", fromKey)
	}

	if toVertex == nil {
		return fmt.Errorf("to vertex with key %d does not exist", toKey)
	}

	if containsVertex(fromVertex.Adjacent, toKey) {
		return fmt.Errorf("edge from %d to %d already exists", fromKey, toKey)
	}

	fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
	return nil
}

// getVertex returns a pointer to the vertex with the given key
// Returns nil if the vertex does not exist
func (g *Graph) getVertex(key int) *Vertex {
	for i, v := range g.Vertices {
		if v.Key == key {
			return g.Vertices[i]
		}
	}
	return nil
}

// HasVertex checks if a vertex with the given key exists in the graph
func (g *Graph) HasVertex(key int) bool {
	return containsVertex(g.Vertices, key)
}

// HasEdge checks if an edge exists from the vertex with fromKey to the vertex with toKey
func (g *Graph) HasEdge(fromKey, toKey int) bool {
	fromVertex := g.getVertex(fromKey)
	if fromVertex == nil {
		return false
	}
	return containsVertex(fromVertex.Adjacent, toKey)
}

// GetNeighbors returns a slice of keys representing all vertices adjacent to the vertex with the given key
// Returns nil if the vertex does not exist
func (g *Graph) GetNeighbors(key int) ([]int, error) {
	vertex := g.getVertex(key)
	if vertex == nil {
		return nil, fmt.Errorf("vertex with key %d does not exist", key)
	}

	var neighbors []int
	for _, v := range vertex.Adjacent {
		neighbors = append(neighbors, v.Key)
	}
	return neighbors, nil
}

// GetVertexCount returns the number of vertices in the graph
func (g *Graph) GetVertexCount() int {
	return len(g.Vertices)
}

// GetAllVertices returns a slice of all vertex keys in the graph
func (g *Graph) GetAllVertices() []int {
	var keys []int
	for _, v := range g.Vertices {
		keys = append(keys, v.Key)
	}
	return keys
}

// containsVertex checks if a vertex with the given key exists in the slice of vertices
func containsVertex(vertices []*Vertex, key int) bool {
	for _, v := range vertices {
		if v.Key == key {
			return true
		}
	}
	return false
}

// String returns a string representation of the graph
func (g *Graph) String() string {
	result := "Graph:\n"
	for _, v := range g.Vertices {
		result += fmt.Sprintf("Vertex %d: ", v.Key)
		if len(v.Adjacent) == 0 {
			result += "No adjacent vertices\n"
		} else {
			for _, adj := range v.Adjacent {
				result += fmt.Sprintf("%d ", adj.Key)
			}
			result += "\n"
		}
	}
	return result
}

// BFS performs a Breadth-First Search starting from the vertex with the given key
// Returns a slice of vertex keys in BFS order or an error if the start vertex doesn't exist
func (g *Graph) BFS(startKey int) ([]int, error) {
	startVertex := g.getVertex(startKey)
	if startVertex == nil {
		return nil, fmt.Errorf("start vertex with key %d does not exist", startKey)
	}

	visited := make(map[int]bool)
	var result []int

	// Create a queue for BFS
	queue := []*Vertex{startVertex}
	visited[startVertex.Key] = true

	for len(queue) > 0 {
		// Dequeue a vertex from queue
		current := queue[0]
		queue = queue[1:]

		// Add to result
		result = append(result, current.Key)

		// Process all adjacent vertices
		for _, adj := range current.Adjacent {
			if !visited[adj.Key] {
				visited[adj.Key] = true
				queue = append(queue, adj)
			}
		}
	}

	return result, nil
}

// DFS performs a Depth-First Search starting from the vertex with the given key
// Returns a slice of vertex keys in DFS order or an error if the start vertex doesn't exist
func (g *Graph) DFS(startKey int) ([]int, error) {
	startVertex := g.getVertex(startKey)
	if startVertex == nil {
		return nil, fmt.Errorf("start vertex with key %d does not exist", startKey)
	}

	visited := make(map[int]bool)
	var result []int

	// Use helper function to perform recursive DFS
	g.dfsHelper(startVertex, visited, &result)

	return result, nil
}

// dfsHelper is a recursive helper function for DFS
func (g *Graph) dfsHelper(vertex *Vertex, visited map[int]bool, result *[]int) {
	// Mark the current node as visited and add to result
	visited[vertex.Key] = true
	*result = append(*result, vertex.Key)

	// Recur for all adjacent vertices
	for _, adj := range vertex.Adjacent {
		if !visited[adj.Key] {
			g.dfsHelper(adj, visited, result)
		}
	}
}
