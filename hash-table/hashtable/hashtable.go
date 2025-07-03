package hashtable

import "fmt"

// ArraySize is the size of the hash table array
const ArraySize = 7

// HashTable represents the hash table data structure
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket represents a linked list in each array position
type bucket struct {
	head *bucketNode
}

// bucketNode represents a linked list node that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert adds a key to the hash table
func (ht *HashTable) Insert(key string) {
	index := hash(key)
	ht.array[index].insert(key)
}

// Search returns true if the key exists in the hash table
func (ht *HashTable) Search(key string) bool {
	index := hash(key)
	return ht.array[index].search(key)
}

// Delete removes a key from the hash table and returns true if the key was found and deleted
func (ht *HashTable) Delete(key string) bool {
	index := hash(key)
	return ht.array[index].delete(key)
}

// insert adds a key to a bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Key already exists:", k)
	}
}

// search looks for a key in a bucket and returns true if found
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete removes a key from a bucket and returns true if the key was found and deleted
func (b *bucket) delete(k string) bool {
	// If bucket is empty
	if b.head == nil {
		return false
	}

	// If head is the key to delete
	if b.head.key == k {
		b.head = b.head.next
		return true
	}

	// Find the key in the bucket
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// Delete by updating the pointer
			previousNode.next = previousNode.next.next
			return true
		}
		previousNode = previousNode.next
	}

	// Key was not found
	return false
}

// hash creates a hash value from the key
func hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % ArraySize
}

// InitHashTable creates and initializes a new HashTable
func InitHashTable() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// GetHashValue returns the hash value for a given key without applying modulo
// This is useful for demonstration purposes to understand how the hash function works
func GetHashValue(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % ArraySize
}
