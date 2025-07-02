package main

import (
	"fmt"

	"github.com/phihdn/go-data-structures/stacks-queues/queue"
	"github.com/phihdn/go-data-structures/stacks-queues/stack"
)

func demoStack() {
	fmt.Println("\n=== Stack Demo ===")

	// Create a new stack
	s := stack.Stack{}

	// Push some items
	fmt.Println("Pushing items: 1, 2, 3")
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Check the size
	fmt.Printf("Stack size: %d\n", s.Size())

	// Peek at the top item
	if val, ok := s.Peek(); ok {
		fmt.Printf("Top item (peek): %d\n", val)
	}

	// Pop some items
	if val, ok := s.Pop(); ok {
		fmt.Printf("Popped item: %d\n", val)
	}

	if val, ok := s.Pop(); ok {
		fmt.Printf("Popped item: %d\n", val)
	}

	// Check the size again
	fmt.Printf("Stack size after pops: %d\n", s.Size())

	// Clear the stack
	s.Clear()
	fmt.Printf("Stack size after clear: %d\n", s.Size())
}

func demoQueue() {
	fmt.Println("\n=== Queue Demo ===")

	// Create a new queue
	q := queue.Queue{}

	// Enqueue some items
	fmt.Println("Enqueuing items: 1, 2, 3")
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Check the size
	fmt.Printf("Queue size: %d\n", q.Size())

	// Look at the front item
	if val, ok := q.Front(); ok {
		fmt.Printf("Front item: %d\n", val)
	}

	// Dequeue some items
	if val, ok := q.Dequeue(); ok {
		fmt.Printf("Dequeued item: %d\n", val)
	}

	if val, ok := q.Dequeue(); ok {
		fmt.Printf("Dequeued item: %d\n", val)
	}

	// Check the size again
	fmt.Printf("Queue size after dequeues: %d\n", q.Size())

	// Clear the queue
	q.Clear()
	fmt.Printf("Queue size after clear: %d\n", q.Size())
}

func main() {
	// Demonstrate Stack operations
	demoStack()

	// Demonstrate Queue operations
	demoQueue()
}
