# Heap

A heap is a specialized tree-based data structure that satisfies the heap property. In a max heap, for any given node, the value of that node is greater than or equal to the values of its children. In a min heap, the value of any node is less than or equal to the values of its children.

This implementation focuses on a Max Heap, where the maximum value is always at the root of the heap.

Video tutorials:

- <https://youtu.be/5EJKm_u6P1o?si=tyrfSruqiSOP8cIO>

## Structure

- `maxheap/`: Package implementing the max heap data structure
  - `maxheap.go`: Core implementation of the `MaxHeap` type
  - `maxheap_test.go`: Unit tests for the max heap implementation
- `cmd/`: Command-line demo application
  - `main.go`: Demo program showing max heap operations

## Features

- Create an empty max heap
- Insert elements into the heap
- Extract the maximum element
- Get the maximum element without removing it
- Build a heap from an existing array
- Check if the heap is empty
- Get the size of the heap

## Time Complexity

| Operation   | Average Case | Worst Case |
|-------------|--------------|------------|
| Insert      | O(log n)     | O(log n)   |
| Extract Max | O(log n)     | O(log n)   |
| Get Max     | O(1)         | O(1)       |
| Build Heap  | O(n)         | O(n)       |

## Space Complexity

Space complexity for a heap is O(n), where n is the number of elements in the heap.

## Usage

Run the demo program to see the max heap in action:

```bash
cd cmd
go run main.go
```

## Common Use Cases

- Priority queues (where elements with higher priority are served first)
- Heap sort algorithm
- Finding the k largest/smallest elements in a collection
- Graph algorithms like Dijkstra's shortest path
- Median maintenance
