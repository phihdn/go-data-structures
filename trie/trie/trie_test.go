package trie

import (
	"reflect"
	"sort"
	"testing"
)

// TestInitTrie tests the initialization of a new Trie
func TestInitTrie(t *testing.T) {
	trie := InitTrie()
	if trie == nil {
		t.Errorf("Expected non-nil trie after initialization")
	}
	if trie.root == nil {
		t.Errorf("Expected non-nil root node after initialization")
	}
}

// TestInsertAndSearch tests the Insert and Search operations
func TestInsertAndSearch(t *testing.T) {
	myTrie := InitTrie()

	// Test insert and successful search
	words := []string{"hello", "world", "trie", "data", "structure"}
	for _, word := range words {
		myTrie.Insert(word)
	}

	for _, word := range words {
		if !myTrie.Search(word) {
			t.Errorf("Search failed for word that should be in trie: %s", word)
		}
	}

	// Test unsuccessful search
	notFoundWords := []string{"notfound", "goodbye", "testing"}
	for _, word := range notFoundWords {
		if myTrie.Search(word) {
			t.Errorf("Search found word that shouldn't be in trie: %s", word)
		}
	}

	// Test empty string
	myTrie.Insert("")
	if !myTrie.Search("") {
		t.Errorf("Search failed for empty string after insertion")
	}
}

// TestStartsWith tests the StartsWith operation
func TestStartsWith(t *testing.T) {
	myTrie := InitTrie()

	// Insert some words
	words := []string{"hello", "help", "world", "wonder", "trie"}
	for _, word := range words {
		myTrie.Insert(word)
	}

	// Test prefixes that should exist
	prefixes := []string{"he", "hel", "wo", "tri"}
	for _, prefix := range prefixes {
		if !myTrie.StartsWith(prefix) {
			t.Errorf("StartsWith failed for prefix that should exist: %s", prefix)
		}
	}

	// Test prefixes that should not exist
	nonPrefixes := []string{"abc", "z", "helm", "tries"}
	for _, prefix := range nonPrefixes {
		if myTrie.StartsWith(prefix) {
			t.Errorf("StartsWith found prefix that shouldn't exist: %s", prefix)
		}
	}

	// Test empty prefix (should always return true)
	if !myTrie.StartsWith("") {
		t.Errorf("StartsWith failed for empty prefix")
	}
}

// TestDelete tests the Delete operation
func TestDelete(t *testing.T) {
	myTrie := InitTrie()

	// Insert some words
	words := []string{"hello", "help", "world", "trie", "try"}
	for _, word := range words {
		myTrie.Insert(word)
	}

	// Delete a word
	if !myTrie.Delete("hello") {
		t.Errorf("Delete failed for word that exists: hello")
	}

	// Word should not be found after deletion
	if myTrie.Search("hello") {
		t.Errorf("Search found word that was deleted: hello")
	}

	// Prefix should still exist due to "help"
	if !myTrie.StartsWith("hel") {
		t.Errorf("StartsWith failed for prefix that should exist after deletion: hel")
	}

	// Delete another word
	if !myTrie.Delete("help") {
		t.Errorf("Delete failed for word that exists: help")
	}

	// Prefix should no longer exist
	if myTrie.StartsWith("hel") {
		t.Errorf("StartsWith found prefix that shouldn't exist after deletion: hel")
	}

	// Attempting to delete a non-existent word should return false
	if myTrie.Delete("notfound") {
		t.Errorf("Delete returned true for word that doesn't exist: notfound")
	}

	// Delete a word with a similar prefix
	if !myTrie.Delete("trie") {
		t.Errorf("Delete failed for word that exists: trie")
	}

	// "try" should still be found
	if !myTrie.Search("try") {
		t.Errorf("Search failed for word that should still exist: try")
	}
}

// TestCount tests the Count operation
func TestCount(t *testing.T) {
	myTrie := InitTrie()

	// Empty trie should have count 0
	if count := myTrie.Count(); count != 0 {
		t.Errorf("Count returned %d for empty trie, expected 0", count)
	}

	// Insert some words
	words := []string{"hello", "world", "trie"}
	for _, word := range words {
		myTrie.Insert(word)
	}

	// Check count after insertions
	if count := myTrie.Count(); count != len(words) {
		t.Errorf("Count returned %d after insertions, expected %d", count, len(words))
	}

	// Delete a word
	myTrie.Delete("hello")

	// Check count after deletion
	if count := myTrie.Count(); count != len(words)-1 {
		t.Errorf("Count returned %d after deletion, expected %d", count, len(words)-1)
	}

	// Insert duplicate word
	myTrie.Insert("world")

	// Count should remain the same
	if count := myTrie.Count(); count != len(words)-1 {
		t.Errorf("Count returned %d after duplicate insertion, expected %d", count, len(words)-1)
	}

	// Insert empty string
	myTrie.Insert("")

	// Count should increase
	if count := myTrie.Count(); count != len(words) {
		t.Errorf("Count returned %d after empty string insertion, expected %d", count, len(words))
	}
}

// TestListWords tests the ListWords operation
func TestListWords(t *testing.T) {
	myTrie := InitTrie()

	// Empty trie should return empty list
	if words := myTrie.ListWords(); len(words) != 0 {
		t.Errorf("ListWords returned %v for empty trie, expected []", words)
	}

	// Insert some words
	expected := []string{"hello", "world", "trie", "test"}
	for _, word := range expected {
		myTrie.Insert(word)
	}

	// Get words and sort for comparison
	actual := myTrie.ListWords()
	sort.Strings(actual)
	sort.Strings(expected)

	// Compare results
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ListWords returned %v, expected %v", actual, expected)
	}

	// Delete a word
	myTrie.Delete("hello")

	// Get updated words
	expected = []string{"world", "trie", "test"}
	actual = myTrie.ListWords()
	sort.Strings(actual)
	sort.Strings(expected)

	// Compare results after deletion
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ListWords returned %v after deletion, expected %v", actual, expected)
	}
}
