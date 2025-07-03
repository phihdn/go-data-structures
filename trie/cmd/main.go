package main

import "fmt"

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{
		root: &Node{},
	}
	return result
}

func (t *Trie) Insert(word string) {
	wordLength := len(word)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	wordLength := len(word)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

func main() {
	myTrie := InitTrie()

	toInsertWords := []string{"hello", "world", "trie", "data", "structure"}

	for _, word := range toInsertWords {
		myTrie.Insert(word)
	}

	fmt.Println("Inserted words into the trie.")
	fmt.Println("Searching for 'hello':", myTrie.Search("hello"))
	fmt.Println("Searching for 'notfound':", myTrie.Search("notfound"))
}
