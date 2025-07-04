package main

import (
	"fmt"

	"github.com/phihdn/go-data-structures/graph/graph"
)

func main() {
	// Initialize a new graph
	fmt.Println("Initializing Graph...")
	g := graph.NewGraph()

	// Add vertices
	fmt.Println("\nAdding vertices 1-5...")
	for i := 1; i <= 5; i++ {
		if err := g.AddVertex(i); err != nil {
			fmt.Printf("Error adding vertex %d: %v\n", i, err)
		}
	}

	// Try adding a duplicate vertex
	fmt.Println("\nTrying to add duplicate vertex 1...")
	if err := g.AddVertex(1); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Add edges to create a graph structure
	fmt.Println("\nAdding edges to create a graph structure...")
	fmt.Println("1 --> 2 --> 4")
	fmt.Println("|     |")
	fmt.Println("v     v")
	fmt.Println("3 --> 5")

	edges := [][2]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 5},
	}

	for _, e := range edges {
		if err := g.AddEdge(e[0], e[1]); err != nil {
			fmt.Printf("Error adding edge %d->%d: %v\n", e[0], e[1], err)
		}
	}

	// Try adding a duplicate edge
	fmt.Println("\nTrying to add duplicate edge 1->2...")
	if err := g.AddEdge(1, 2); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Print the graph
	fmt.Println("\nGraph structure:")
	fmt.Println(g)

	// Check if vertices exist
	fmt.Printf("Has vertex 3: %v\n", g.HasVertex(3))
	fmt.Printf("Has vertex 6: %v\n", g.HasVertex(6))

	// Check if edges exist
	fmt.Printf("Has edge 1->2: %v\n", g.HasEdge(1, 2))
	fmt.Printf("Has edge 2->1: %v\n", g.HasEdge(2, 1)) // Should be false (directed graph)

	// Get neighbors
	fmt.Println("\nNeighbors of vertex 1:")
	if neighbors, err := g.GetNeighbors(1); err == nil {
		fmt.Println(neighbors)
	} else {
		fmt.Printf("Error: %v\n", err)
	}

	// BFS traversal
	fmt.Println("\nBreadth-First Search starting from vertex 1:")
	if bfsResult, err := g.BFS(1); err == nil {
		fmt.Println(bfsResult)
	} else {
		fmt.Printf("Error: %v\n", err)
	}

	// DFS traversal
	fmt.Println("\nDepth-First Search starting from vertex 1:")
	if dfsResult, err := g.DFS(1); err == nil {
		fmt.Println(dfsResult)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
}
