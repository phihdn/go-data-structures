# Stacks and Queues

This folder contains implementations of two fundamental data structures: stacks and queues.

## Stack

A stack is a linear data structure that follows the Last-In-First-Out (LIFO) principle. The last element added to the stack is the first one to be removed.

### Stack Operations

- **Push**: Add an element to the top of the stack
- **Pop**: Remove the top element from the stack
- **Peek**: View the top element without removing it
- **IsEmpty**: Check if the stack is empty
- **Size**: Get the number of elements in the stack
- **Clear**: Remove all elements from the stack

### Stack Time Complexity

- Push: O(1)
- Pop: O(1)
- Peek: O(1)
- IsEmpty: O(1)
- Size: O(1)
- Clear: O(1)

### Stack Usage

```go
import "github.com/phihdn/go-data-structures/stacks-queues/stack"

// Create a new stack
s := stack.Stack{}

// Push items
s.Push(1)
s.Push(2)
s.Push(3)

// Pop an item
if val, ok := s.Pop(); ok {
    fmt.Printf("Popped: %d\n", val)
}

// Peek at the top item
if val, ok := s.Peek(); ok {
    fmt.Printf("Top item: %d\n", val)
}

// Check if empty
if s.IsEmpty() {
    fmt.Println("Stack is empty")
} else {
    fmt.Printf("Stack size: %d\n", s.Size())
}
```

## Queue

A queue is a linear data structure that follows the First-In-First-Out (FIFO) principle. The first element added to the queue is the first one to be removed.

### Queue Operations

- **Enqueue**: Add an element to the end of the queue
- **Dequeue**: Remove the front element from the queue
- **Front**: View the front element without removing it
- **IsEmpty**: Check if the queue is empty
- **Size**: Get the number of elements in the queue
- **Clear**: Remove all elements from the queue

### Queue Time Complexity

- Enqueue: O(1) amortized
- Dequeue: O(n) due to slice shift
- Front: O(1)
- IsEmpty: O(1)
- Size: O(1)
- Clear: O(1)

### Queue Usage

```go
import "github.com/phihdn/go-data-structures/stacks-queues/queue"

// Create a new queue
q := queue.Queue{}

// Enqueue items
q.Enqueue(1)
q.Enqueue(2)
q.Enqueue(3)

// Dequeue an item
if val, ok := q.Dequeue(); ok {
    fmt.Printf("Dequeued: %d\n", val)
}

// View the front item
if val, ok := q.Front(); ok {
    fmt.Printf("Front item: %d\n", val)
}

// Check if empty
if q.IsEmpty() {
    fmt.Println("Queue is empty")
} else {
    fmt.Printf("Queue size: %d\n", q.Size())
}
```

## Running the Demo

To run the demo program:

```bash
cd stacks-queues/cmd
go run main.go
```

This will demonstrate the operations of both stacks and queues.
