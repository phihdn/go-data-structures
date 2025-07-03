# Binary Search Tree

A binary search tree (BST) implementation in Go. A binary search tree is a tree data structure with the following properties:

- The left subtree of a node contains only nodes with keys lesser than the node's key
- The right subtree of a node contains only nodes with keys greater than the node's key
- The left and right subtree each must also be a binary search tree

## Features

- Node structure with key and left/right child pointers
- Basic operations:
  - Insert a new key
  - Search for a key
  - Find minimum and maximum values
- Tree traversal algorithms:
  - In-order traversal (returns sorted keys)
  - Pre-order traversal
  - Post-order traversal

## Time Complexity

| Operation       | Average Case | Worst Case |
|-----------------|--------------|------------|
| Insert          | O(log n)     | O(n)       |
| Search          | O(log n)     | O(n)       |
| Min/Max         | O(log n)     | O(n)       |
| Tree Traversal  | O(n)         | O(n)       |

The worst case occurs when the tree becomes unbalanced, approaching a linked list structure.

## Usage Example

```go
// Create a new binary search tree with root key 100
tree := bst.NewBST(100)

// Insert some values
tree.Insert(50)
tree.Insert(150)
tree.Insert(25)
tree.Insert(75)

// Search for values
found := tree.Search(25)     // returns true
notFound := tree.Search(200) // returns false

// Get min and max values
min := tree.Min() // returns 25
max := tree.Max() // returns 150

// Tree traversals
inOrder := tree.InOrderTraversal()   // returns [25, 50, 75, 100, 150]
preOrder := tree.PreOrderTraversal() // returns [100, 50, 25, 75, 150]
postOrder := tree.PostOrderTraversal() // returns [25, 75, 50, 150, 100]
```

See the `cmd/main.go` file for complete usage examples.

## Testing

To run the tests:

```bash
cd binary-search-tree
go test ./bst
```
