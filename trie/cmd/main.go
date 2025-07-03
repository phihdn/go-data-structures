package main

import (
	"fmt"

	"github.com/phihdn/go-data-structures/trie/trie"
)

func main() {
	// Initialize a new trie
	fmt.Println("Initializing a new trie...")
	myTrie := trie.InitTrie()

	// Insert words
	toInsertWords := []string{"hello", "world", "trie", "data", "structure", "help", "word"}
	fmt.Println("Inserting words:", toInsertWords)
	for _, word := range toInsertWords {
		myTrie.Insert(word)
	}

	// Search for words
	fmt.Println("\n--- Search Operations ---")
	fmt.Println("Searching for 'hello':", myTrie.Search("hello"))
	fmt.Println("Searching for 'notfound':", myTrie.Search("notfound"))

	// Check prefixes
	fmt.Println("\n--- Prefix Operations ---")
	fmt.Println("Starts with 'hel':", myTrie.StartsWith("hel"))
	fmt.Println("Starts with 'dat':", myTrie.StartsWith("dat"))
	fmt.Println("Starts with 'abc':", myTrie.StartsWith("abc"))

	// Count words
	fmt.Println("\n--- Count Operation ---")
	fmt.Println("Number of words in trie:", myTrie.Count())

	// List all words
	fmt.Println("\n--- List Words Operation ---")
	fmt.Println("All words in trie:", myTrie.ListWords())

	// Delete words
	fmt.Println("\n--- Delete Operations ---")
	wordToDelete := "hello"
	fmt.Printf("Deleting '%s': %v\n", wordToDelete, myTrie.Delete(wordToDelete))
	fmt.Printf("Searching for '%s' after deletion: %v\n", wordToDelete, myTrie.Search(wordToDelete))
	fmt.Println("Starts with 'hel' after deletion:", myTrie.StartsWith("hel"))

	// Count after deletion
	fmt.Println("\n--- Final State ---")
	fmt.Println("Number of words after deletion:", myTrie.Count())
	fmt.Println("All words after deletion:", myTrie.ListWords())
}
