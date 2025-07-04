package main

import (
	"fmt"
	"strings"

	"github.com/phihdn/go-data-structures/hash-table/hashtable"
)

func main() {
	fmt.Println("=== HASH TABLE DEMONSTRATION ===")

	// Initialize a new hash table
	fmt.Println("\n1. Creating a new hash table")
	myHashTable := hashtable.InitHashTable()

	// Sample data for demonstration
	fmt.Println("\n2. Preparing sample data")
	listOfWords := []string{"ERIC", "KENNY", "KYLE", "STAN", "BUTTERS", "RANDY", "WENDY", "TOMMY"}
	fmt.Println("   Sample words:", strings.Join(listOfWords, ", "))

	// Insert operation demonstration
	fmt.Println("\n3. Inserting words into the hash table")
	for _, word := range listOfWords {
		fmt.Printf("   Inserting '%s'... ", word)
		myHashTable.Insert(word)
		fmt.Println("Done")
	}

	// Attempt to insert duplicate
	fmt.Println("\n4. Attempting to insert a duplicate")
	fmt.Print("   Inserting 'ERIC' again... ")
	myHashTable.Insert("ERIC") // This should print "Key already exists: ERIC"

	// Search operation demonstration
	fmt.Println("\n5. Searching for words in the hash table")
	fmt.Printf("   Search for 'ERIC': %t\n", myHashTable.Search("ERIC"))
	fmt.Printf("   Search for 'KENNY': %t\n", myHashTable.Search("KENNY"))
	fmt.Printf("   Search for 'CARTMAN' (not inserted): %t\n", myHashTable.Search("CARTMAN"))

	// Delete operation demonstration
	fmt.Println("\n6. Deleting words from the hash table")

	// Delete an existing word
	deleted := myHashTable.Delete("ERIC")
	fmt.Printf("   Deleted 'ERIC': %t\n", deleted)
	fmt.Printf("   Search for 'ERIC' after deletion: %t\n", myHashTable.Search("ERIC"))

	// Try to delete a non-existent word
	deleted = myHashTable.Delete("NONEXISTENT")
	fmt.Printf("   Deleted 'NONEXISTENT' (not inserted): %t\n", deleted)

	// Try to delete an already deleted word
	deleted = myHashTable.Delete("ERIC")
	fmt.Printf("   Deleted 'ERIC' again (already deleted): %t\n", deleted)

	// Demonstrating hash collisions
	fmt.Println("\n7. Demonstrating hash collisions")
	// These words are chosen to potentially cause hash collisions
	// with our simple hash function
	collisionDemo := hashtable.InitHashTable()

	word1 := "abc"
	word2 := "cba"

	// Check if they hash to the same value
	hash1 := hashtable.GetHashValue(word1) // Note: We'll need to add this function
	hash2 := hashtable.GetHashValue(word2)

	fmt.Printf("   Hash value of '%s': %d\n", word1, hash1)
	fmt.Printf("   Hash value of '%s': %d\n", word2, hash2)

	if hash1 == hash2 {
		fmt.Printf("   Collision detected: '%s' and '%s' hash to the same value!\n", word1, word2)

		// Insert both words
		collisionDemo.Insert(word1)
		collisionDemo.Insert(word2)

		// Search for both words
		fmt.Printf("   Search for '%s': %t\n", word1, collisionDemo.Search(word1))
		fmt.Printf("   Search for '%s': %t\n", word2, collisionDemo.Search(word2))

		// Delete one word and check if the other still exists
		collisionDemo.Delete(word1)
		fmt.Printf("   Deleted '%s'\n", word1)
		fmt.Printf("   Search for '%s' after deletion: %t\n", word1, collisionDemo.Search(word1))
		fmt.Printf("   Search for '%s' after deletion of '%s': %t\n", word2, word1, collisionDemo.Search(word2))
	} else {
		fmt.Printf("   No collision: '%s' hashes to %d, '%s' hashes to %d\n", word1, hash1, word2, hash2)

		// Find two words that do collide by checking some combinations
		foundCollision := false
		var collidingWord1, collidingWord2 string

		testWords := []string{"ab", "ba", "ac", "ca", "xy", "yx", "abc", "cba"}

		fmt.Println("   Looking for collisions among test words...")

		// Check all pairs for collisions
		for i := 0; i < len(testWords); i++ {
			for j := i + 1; j < len(testWords); j++ {
				hash1 = hashtable.GetHashValue(testWords[i])
				hash2 = hashtable.GetHashValue(testWords[j])

				if hash1 == hash2 {
					collidingWord1 = testWords[i]
					collidingWord2 = testWords[j]
					foundCollision = true
					break
				}
			}
			if foundCollision {
				break
			}
		}

		if foundCollision {
			fmt.Printf("   Found collision: '%s' and '%s' both hash to %d\n",
				collidingWord1, collidingWord2, hash1)

			// Demonstrate the collision handling
			collisionDemo.Insert(collidingWord1)
			collisionDemo.Insert(collidingWord2)

			fmt.Printf("   Search for '%s': %t\n", collidingWord1, collisionDemo.Search(collidingWord1))
			fmt.Printf("   Search for '%s': %t\n", collidingWord2, collisionDemo.Search(collidingWord2))
		} else {
			fmt.Println("   No collisions found among test words.")
		}
	}

	fmt.Println("\n=== HASH TABLE DEMONSTRATION COMPLETE ===")
}
