# Go Data Structures

This repository contains implementations of various data structures in Go. Each implementation includes:

- Clean, well-documented Go code
- Comprehensive test coverage
- Example usage in a demo program
- Documentation with Big O time complexity analysis

## Project Structure

```plaintext
go-data-structures/
├── README.md                 # This file
├── go.mod                    # Module definition
├── binary-search-tree/       # Binary Search Tree implementation
│   ├── README.md             # BST documentation
│   ├── cmd/                  # Command-line demo
│   │   └── main.go           # Demo program for BST
│   └── bst/                  # Package implementation
│       ├── bst.go            # BST code
│       └── bst_test.go       # BST tests
├── hash-table/              # Hash Table implementation
│   ├── README.md            # Hash Table documentation
│   ├── cmd/                 # Command-line demo
│   │   └── main.go          # Demo program for hash table
│   └── hashtable/           # Package implementation
│       ├── hashtable.go     # Hash table code
│       └── hashtable_test.go # Hash table tests
├── linked-list/              # Linked List implementation
│   ├── cmd/                  # Command-line demo
│   │   └── main.go           # Demo program for linked list
│   └── linkedlist/           # Package implementation
│       ├── linkedlist.go     # Linked list code
│       └── linkedlist_test.go # Linked list tests
├── stacks-queues/            # Stack and Queue implementations
│   ├── README.md             # Stack and Queue documentation
│   ├── cmd/                  # Command-line demo
│   │   └── main.go           # Demo program for stack and queue
│   ├── stack/                # Stack package implementation
│   │   ├── stack.go          # Stack code
│   │   └── stack_test.go     # Stack tests
│   └── queue/                # Queue package implementation
│       ├── queue.go          # Queue code
│       └── queue_test.go     # Queue tests
├── trie/                    # Trie implementation
│   ├── README.md            # Trie documentation
│   ├── cmd/                 # Command-line demo
│   │   └── main.go          # Demo program for trie
│   └── trie/                # Package implementation
│       ├── trie.go          # Trie code
│       └── trie_test.go     # Trie tests
└── ... (other data structures)
```

## Running the Examples

To run a specific data structure example:

```bash
# Run the binary search tree example
cd binary-search-tree/cmd
go run main.go

# Run the linked list example
cd linked-list/cmd
go run main.go

# Run the stack and queue examples
cd stacks-queues/cmd
go run main.go

# Run the hash table example
cd hash-table/cmd
go run main.go

# Run the trie example
cd trie/cmd
go run main.go
```

## Available Data Structures

1. **Binary Search Tree** - A tree data structure with the key in each node being greater than all keys in the left subtree and less than all keys in the right subtree
2. **Hash Table** - A data structure that uses a hash function to map keys to values for efficient lookup
3. **Linked List** - A linear collection of elements where each element points to the next
4. **Stack** - A LIFO (Last In First Out) data structure that supports push and pop operations
5. **Queue** - A FIFO (First In First Out) data structure that supports enqueue and dequeue operations
6. **Trie** - A tree-like data structure used for efficient storage and retrieval of strings, commonly used for autocomplete and spell checking

## Running Tests

You can run tests for the implemented data structures:

```bash
# Run tests for binary search tree
cd binary-search-tree
go test ./bst

# Run tests for linked list
cd linked-list
go test ./linkedlist

# Run tests for stack and queue
cd stacks-queues
go test ./stack ./queue

# Run tests for hash table
cd hash-table
go test ./hashtable

# Run tests for trie
cd trie
go test ./trie

# Run tests for all data structures (from root directory)
go test ./...
```

## Adding New Data Structures

To add a new data structure:

1. Create a new directory for it at the root level
2. Add the implementation in a package subdirectory
3. Create a cmd/main.go file to demonstrate its usage
4. Add tests in *_test.go files
5. Update this README.md to include your new data structure
