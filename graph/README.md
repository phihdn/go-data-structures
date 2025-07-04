# Graph

A graph data structure implementation in Go. A graph is a non-linear data structure consisting of nodes (vertices) and edges connecting these nodes.

Video tutorials:

- <https://www.youtube.com/watch?v=D3oPCqn_HO8>
- <https://www.youtube.com/watch?v=JQ79-7q75qU>

## Features

- Directed graph implementation with vertices and adjacency lists
- Basic operations:
  - Add vertices
  - Add directed edges between vertices
  - Check if a vertex exists
  - Check if an edge exists
  - Get neighbors of a vertex
- Graph traversal algorithms:
  - Breadth-First Search (BFS)
  - Depth-First Search (DFS)

## Time Complexity

| Operation                | Time Complexity |
|--------------------------|-----------------|
| Add Vertex               | O(n)            |
| Add Edge                 | O(n)            |
| Check if Vertex Exists   | O(n)            |
| Check if Edge Exists     | O(n)            |
| Get Neighbors            | O(1)            |
| BFS Traversal            | O(V + E)        |
| DFS Traversal            | O(V + E)        |

Where:

- V is the number of vertices
- E is the number of edges

## Usage Example

```go
// Create a new graph
g := graph.NewGraph()

// Add vertices
_ = g.AddVertex(1)
_ = g.AddVertex(2)
_ = g.AddVertex(3)

// Add edges
_ = g.AddEdge(1, 2)
_ = g.AddEdge(2, 3)
_ = g.AddEdge(1, 3)

// Check if vertex exists
hasVertex := g.HasVertex(1)     // returns true

// Check if edge exists
hasEdge := g.HasEdge(1, 2)      // returns true
hasEdge = g.HasEdge(3, 1)       // returns false (directed graph)

// Get neighbors
neighbors, _ := g.GetNeighbors(1)  // returns [2, 3]

// Graph traversals
bfs, _ := g.BFS(1)              // Breadth-First Search
dfs, _ := g.DFS(1)              // Depth-First Search
```

See the `cmd/main.go` file for complete usage examples.

## Testing

To run the tests:

```bash
cd graph
go test ./graph
```

## Applications of Graphs

- Social Networks
- Web Page Links
- Maps and Navigation
- Network Routing
- Dependency Resolution
- Recommendation Systems
