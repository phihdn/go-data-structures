# Linked List Implementation in Go

This directory contains a simple singly linked list implementation in Go.

Video tutorials:

- <https://www.youtube.com/watch?v=zgAwMts9bRg>
- <https://www.youtube.com/watch?v=PGcTioRPBhU>

## Structure

- `linkedlist/`: Package implementing the linked list data structure
  - `linkedlist.go`: Core implementation of the `Node` and `LinkedList` types
  - `linkedlist_test.go`: Unit tests for the linked list implementation
- `cmd/`: Command-line demo application
  - `main.go`: Demo program showing linked list operations
  - `main_test.go`: Integration tests for the linked list

## Features

- Create a singly linked list
- Prepend nodes to the list
- Delete nodes by value
- Print list contents
- Track list length

## Running the Demo

```bash
cd cmd
go run main.go
```

## Running the Tests

### Run all tests

```bash
cd linked-list
go test ./...
```

### Run unit tests with coverage

```bash
cd linkedlist
go test -v -cover
```

### Run integration tests

```bash
cd cmd
go test -v
```

## Test Coverage

The tests cover:

1. **Basic operations**
   - Creating a new linked list
   - Prepending nodes
   - Deleting nodes by value
   - Checking list length

2. **Edge cases**
   - Operations on an empty list
   - Deleting the first node
   - Deleting a middle node
   - Deleting the last node
   - Deleting a non-existent value
   - Handling duplicate values

3. **Integration testing**
   - Using the linked list package from a client perspective
   - Performing a sequence of operations and verifying the results
