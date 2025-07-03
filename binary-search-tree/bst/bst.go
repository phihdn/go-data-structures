package bst

// Node represents a node in the binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert adds a new node with the given key to the BST
func (n *Node) Insert(key int) {
	if n.Key < key {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else if n.Key > key {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	}
}

// Search checks if a node with the given key exists in the BST
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

// InOrderTraversal performs an in-order traversal of the tree and returns the keys
func (n *Node) InOrderTraversal() []int {
	if n == nil {
		return []int{}
	}

	// In-order traversal: left -> root -> right
	result := []int{}

	// Traverse left subtree
	if n.Left != nil {
		result = append(result, n.Left.InOrderTraversal()...)
	}

	// Visit root
	result = append(result, n.Key)

	// Traverse right subtree
	if n.Right != nil {
		result = append(result, n.Right.InOrderTraversal()...)
	}

	return result
}

// PreOrderTraversal performs a pre-order traversal of the tree and returns the keys
func (n *Node) PreOrderTraversal() []int {
	if n == nil {
		return []int{}
	}

	// Pre-order traversal: root -> left -> right
	result := []int{n.Key} // Visit root first

	// Traverse left subtree
	if n.Left != nil {
		result = append(result, n.Left.PreOrderTraversal()...)
	}

	// Traverse right subtree
	if n.Right != nil {
		result = append(result, n.Right.PreOrderTraversal()...)
	}

	return result
}

// PostOrderTraversal performs a post-order traversal of the tree and returns the keys
func (n *Node) PostOrderTraversal() []int {
	if n == nil {
		return []int{}
	}

	// Post-order traversal: left -> right -> root
	result := []int{}

	// Traverse left subtree
	if n.Left != nil {
		result = append(result, n.Left.PostOrderTraversal()...)
	}

	// Traverse right subtree
	if n.Right != nil {
		result = append(result, n.Right.PostOrderTraversal()...)
	}

	// Visit root
	result = append(result, n.Key)

	return result
}

// Min returns the minimum value in the BST
func (n *Node) Min() int {
	if n.Left == nil {
		return n.Key
	}
	return n.Left.Min()
}

// Max returns the maximum value in the BST
func (n *Node) Max() int {
	if n.Right == nil {
		return n.Key
	}
	return n.Right.Max()
}

// NewBST creates a new binary search tree with the given root key
func NewBST(key int) *Node {
	return &Node{Key: key}
}
