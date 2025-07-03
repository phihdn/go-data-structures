package trie

// AlphabetSize is the size of the lowercase English alphabet
const AlphabetSize = 26

// Node represents a node in the Trie
type Node struct {
	children [AlphabetSize]*Node // Child nodes for each letter
	isEnd    bool                // Flag indicating if this node represents the end of a word
}

// Trie is a data structure for efficient retrieval of words
type Trie struct {
	root *Node // Root node of the Trie
}

// InitTrie creates and initializes a new Trie
func InitTrie() *Trie {
	result := &Trie{
		root: &Node{},
	}
	return result
}

// Insert adds a new word to the Trie
func (t *Trie) Insert(word string) {
	currentNode := t.root

	for _, char := range word {
		// Convert the character to an array index (0-25)
		// For example:
		// 'a' - 'a' = 0
		// 'b' - 'a' = 1
		// 'z' - 'a' = 25
		// This maps each lowercase letter to an index in the children array,
		// where 'a' is at index 0, 'b' is at index 1, and so on
		charIndex := char - 'a'

		// Create a new node if the path doesn't exist
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	// Mark the end of the word
	currentNode.isEnd = true
}

// Search checks if a word exists in the Trie
func (t *Trie) Search(word string) bool {
	currentNode := t.root

	for _, char := range word {
		// Convert the character to an array index (0-25)
		// This is the same mapping used in Insert:
		// 'a' (ASCII 97) - 'a' (ASCII 97) = 0
		// 'b' (ASCII 98) - 'a' (ASCII 97) = 1
		// And so on...
		charIndex := char - 'a'

		// Return false if the path doesn't exist
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	// Return true only if this is the end of a word
	return currentNode.isEnd
}

// StartsWith checks if any word in the Trie starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	currentNode := t.root

	for _, char := range prefix {
		// Convert the character to an array index (0-25)
		// Using the same character-to-index mapping:
		// The ASCII value of the current character minus the ASCII value of 'a'
		// This gives us the relative position in the alphabet (0-25)
		charIndex := char - 'a'

		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return true
}

// Delete removes a word from the Trie if it exists
func (t *Trie) Delete(word string) bool {
	if word == "" {
		if t.root.isEnd {
			t.root.isEnd = false
			return true
		}
		return false
	}

	// Use a different approach that tracks whether the word was found
	found := false
	t.deleteHelper(t.root, word, 0, &found)
	return found
}

// deleteHelper is a helper function for Delete
// It modifies found to indicate if the word was found and deleted
func (t *Trie) deleteHelper(node *Node, word string, depth int, found *bool) bool {
	// Base case: end of the word
	if depth == len(word) {
		// Word not found if this isn't marked as end of word
		if !node.isEnd {
			return false
		}

		// Mark as not end of word and indicate word was found
		node.isEnd = false
		*found = true

		// Return true if this node has no children and can be deleted
		for _, child := range node.children {
			if child != nil {
				return false
			}
		}
		return true
	}

	// Get current character index
	// Convert the character to its corresponding index in the children array
	// This is the same mapping used throughout the Trie:
	// Subtract the ASCII value of 'a' from the current character
	// to get a zero-based index for the children array (0-25)
	charIndex := word[depth] - 'a'
	if node.children[charIndex] == nil {
		// Path doesn't exist, word not found
		return false
	}

	// Recursively delete in child node
	shouldDeleteChild := t.deleteHelper(node.children[charIndex], word, depth+1, found)

	// Delete the child node if needed
	if shouldDeleteChild {
		node.children[charIndex] = nil

		// Check if this node can be deleted too
		if !node.isEnd {
			for _, child := range node.children {
				if child != nil {
					return false
				}
			}
			return true
		}
	}

	return false
}

// Count returns the number of words stored in the Trie
func (t *Trie) Count() int {
	return countWords(t.root)
}

// countWords is a helper function that counts words recursively
func countWords(node *Node) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.isEnd {
		count = 1
	}

	for _, child := range node.children {
		if child != nil {
			count += countWords(child)
		}
	}

	return count
}

// ListWords returns all words stored in the Trie
func (t *Trie) ListWords() []string {
	result := []string{}
	collectWords(t.root, "", &result)
	return result
}

// collectWords is a helper function that collects words recursively
func collectWords(node *Node, prefix string, result *[]string) {
	if node.isEnd {
		*result = append(*result, prefix)
	}

	for i, child := range node.children {
		if child != nil {
			// Convert the index back to the corresponding character
			// This is the inverse of the charIndex calculation:
			// Add the index (0-25) to the ASCII value of 'a'
			// to get the actual character ('a' to 'z')
			char := rune('a' + i)
			collectWords(child, prefix+string(char), result)
		}
	}
}
