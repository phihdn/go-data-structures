# Hash Table

A hash table is a data structure that implements an associative array abstract data type, a structure that can map keys to values. It uses a hash function to compute an index into an array of buckets or slots, from which the desired value can be found.

Video tutorials:

- <https://www.youtube.com/watch?v=JCXLwfKMWOU>

## Structure

- `hashtable/`: Package implementing the hash table data structure
  - `hashtable.go`: Core implementation of the `HashTable`, `bucket`, and `bucketNode` types
  - `hashtable_test.go`: Unit tests for the hash table implementation
- `cmd/`: Command-line demo application
  - `main.go`: Demo program showing hash table operations

## Features

- Efficient key insertion, deletion, and lookup
- Constant-time average case complexity for operations
- Collision resolution using chaining with linked lists
- Simple hash function for string keys
- Return values that indicate operation success/failure

## Time Complexity

| Operation | Average Case | Worst Case |
|-----------|--------------|------------|
| Insert    | O(1)         | O(n)       |
| Search    | O(1)         | O(n)       |
| Delete    | O(1)         | O(n)       |

Where:

- n is the number of elements in the hash table
- The worst-case occurs when all elements hash to the same index

## Space Complexity

Space complexity is O(n), where n is the number of key-value pairs stored.

## Implementation Details

The hash table uses an array of buckets, where each bucket is a linked list that stores keys that hash to the same index. This approach, known as chaining, handles hash collisions by allowing multiple keys to be stored at the same index.

### Key Components

- **HashTable**: The main data structure with an array of buckets
- **Bucket**: A linked list to handle collisions
- **BucketNode**: A node in the linked list that stores a key

### Hash Function

The hash function sums the ASCII values of the characters in the key and then performs a modulo operation to get an index within the array size:

```go
func hash(key string) int {
    sum := 0
    for _, char := range key {
        sum += int(char)
    }
    return sum % ArraySize
}
```

## Usage Example

```go
// Initialize a new hash table
ht := hashtable.InitHashTable()

// Insert keys
ht.Insert("apple")
ht.Insert("banana")
ht.Insert("cherry")

// Search for keys
fmt.Println(ht.Search("apple"))  // true
fmt.Println(ht.Search("grape"))  // false

// Delete a key and check the result
deleted := ht.Delete("apple")
fmt.Println("Deleted:", deleted)  // true
fmt.Println(ht.Search("apple"))   // false
```

## Running the Demo

```bash
cd cmd
go run main.go
```

## Running the Tests

```bash
cd hashtable
go test -v
```

## Test Coverage

The tests cover:

1. **Basic operations**
   - Creating a new hash table
   - Inserting keys
   - Searching for keys
   - Deleting keys

2. **Edge cases**
   - Operations on an empty hash table
   - Deleting non-existent keys
   - Deleting already deleted keys
   - Case sensitivity in keys

3. **Collision handling**
   - Inserting keys that hash to the same index
   - Searching for keys that have collisions
   - Deleting keys that have collisions

## Common Use Cases

- Implementing dictionaries and maps
- Database indexing
- Caching (memoization)
- Symbol tables in compilers
- Authentication (storing passwords)
- Counting frequency of items
- De-duplicating data
