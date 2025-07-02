# Go Data Structures

This repository contains implementations of various data structures in Go.

## Project Structure

```plaintext
go-data-structures/
├── README.md                 # This file
├── go.mod                    # Module definition
├── linked-list/              # Linked List implementation
│   ├── cmd/                  # Command-line demo
│   │   └── main.go           # Demo program for linked list
│   └── linkedlist/           # Package implementation
│       └── linkedlist.go     # Linked list code
└── ... (other data structures)
```

## Running the Examples

To run a specific data structure example:

```bash
# Run the linked list example
cd linked-list/cmd
go run main.go
```

## Available Data Structures

1. **Linked List** - A linear collection of elements where each element points to the next

## Adding New Data Structures

To add a new data structure:

1. Create a new directory for it at the root level
2. Add the implementation in a package subdirectory
3. Create a cmd/main.go file to demonstrate its usage
