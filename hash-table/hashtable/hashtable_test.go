package hashtable

import (
	"strings"
	"testing"
)

func TestHashTableInsert(t *testing.T) {
	ht := InitHashTable()

	// Test inserting keys
	ht.Insert("ERIC")
	ht.Insert("KENNY")

	// Check if keys exist
	if !ht.Search("ERIC") {
		t.Errorf("Expected to find key 'ERIC' in hash table")
	}
	if !ht.Search("KENNY") {
		t.Errorf("Expected to find key 'KENNY' in hash table")
	}
}

func TestHashTableDelete(t *testing.T) {
	ht := InitHashTable()

	// Insert keys
	ht.Insert("ERIC")
	ht.Insert("KENNY")

	// Delete a key and verify it returns true for successful deletion
	if deleted := ht.Delete("ERIC"); !deleted {
		t.Errorf("Delete operation should return true when key 'ERIC' is deleted")
	}

	// Check if the key was deleted
	if ht.Search("ERIC") {
		t.Errorf("Expected key 'ERIC' to be deleted from hash table")
	}
	if !ht.Search("KENNY") {
		t.Errorf("Expected to find key 'KENNY' in hash table")
	}

	// Try to delete a non-existent key and verify it returns false
	if deleted := ht.Delete("NONEXISTENT"); deleted {
		t.Errorf("Delete operation should return false when key 'NONEXISTENT' doesn't exist")
	}

	// Deleting an already deleted key should return false
	if deleted := ht.Delete("ERIC"); deleted {
		t.Errorf("Delete operation should return false when key 'ERIC' is already deleted")
	}
}

func TestHashTableSearch(t *testing.T) {
	ht := InitHashTable()

	// Test with empty hash table
	if ht.Search("NONEXISTENT") {
		t.Errorf("Expected to not find key 'NONEXISTENT' in empty hash table")
	}

	// Insert keys
	keys := []string{"ERIC", "KENNY", "KYLE", "STAN", "BUTTERS", "RANDY"}
	for _, key := range keys {
		ht.Insert(key)
	}

	// Check if all keys exist
	for _, key := range keys {
		if !ht.Search(key) {
			t.Errorf("Expected to find key '%s' in hash table", key)
		}
	}

	// Check for a key that was not inserted
	if ht.Search("NONEXISTENT") {
		t.Errorf("Expected to not find key 'NONEXISTENT' in hash table")
	}
}

func TestHashCollisions(t *testing.T) {
	ht := InitHashTable()

	// These two keys should hash to the same value due to our simple hash function
	key1 := "abc"
	key2 := "cba"

	// Verify they have the same hash
	if hash(key1) != hash(key2) {
		t.Skip("This test assumes keys hash to the same value, but they don't with the current hash function")
	}

	// Insert both keys
	ht.Insert(key1)
	ht.Insert(key2)

	// Check if both keys can be found despite the hash collision
	if !ht.Search(key1) {
		t.Errorf("Expected to find key '%s' in hash table despite hash collision", key1)
	}
	if !ht.Search(key2) {
		t.Errorf("Expected to find key '%s' in hash table despite hash collision", key2)
	}

	// Delete the first key and check return value
	if deleted := ht.Delete(key1); !deleted {
		t.Errorf("Delete operation should return true when key '%s' is deleted", key1)
	}

	// Check if only the first key was deleted
	if ht.Search(key1) {
		t.Errorf("Expected key '%s' to be deleted from hash table", key1)
	}
	if !ht.Search(key2) {
		t.Errorf("Expected to still find key '%s' in hash table after deleting '%s'", key2, key1)
	}
}

func TestTableDrivenInsertAndSearch(t *testing.T) {
	// Table-driven test cases for Insert and Search operations
	tests := []struct {
		name          string
		keysToInsert  []string
		keyToSearch   string
		expectedFound bool
	}{
		{
			name:          "Search for inserted key",
			keysToInsert:  []string{"apple", "banana", "cherry"},
			keyToSearch:   "banana",
			expectedFound: true,
		},
		{
			name:          "Search for non-existent key",
			keysToInsert:  []string{"apple", "banana", "cherry"},
			keyToSearch:   "grape",
			expectedFound: false,
		},
		{
			name:          "Search in empty hash table",
			keysToInsert:  []string{},
			keyToSearch:   "anything",
			expectedFound: false,
		},
		{
			name:          "Search with case sensitivity",
			keysToInsert:  []string{"Apple", "Banana", "Cherry"},
			keyToSearch:   "apple", // lowercase should not match uppercase
			expectedFound: false,
		},
		{
			name:          "Search with special characters",
			keysToInsert:  []string{"apple!", "@banana", "#cherry"},
			keyToSearch:   "@banana",
			expectedFound: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Initialize a new hash table for each test case
			ht := InitHashTable()

			// Insert the keys for this test case
			for _, key := range test.keysToInsert {
				ht.Insert(key)
			}

			// Perform the search
			found := ht.Search(test.keyToSearch)

			// Check if the result matches the expected result
			if found != test.expectedFound {
				t.Errorf("Search for '%s' returned %t, want %t",
					test.keyToSearch, found, test.expectedFound)
			}
		})
	}
}

func TestTableDrivenDelete(t *testing.T) {
	// Table-driven test cases for Delete operation
	tests := []struct {
		name            string
		keysToInsert    []string
		keyToDelete     string
		expectedDeleted bool
		expectedExists  map[string]bool // Map of keys and whether they should exist after deletion
	}{
		{
			name:            "Delete existing key",
			keysToInsert:    []string{"apple", "banana", "cherry"},
			keyToDelete:     "banana",
			expectedDeleted: true,
			expectedExists: map[string]bool{
				"apple":  true,
				"banana": false,
				"cherry": true,
			},
		},
		{
			name:            "Delete non-existent key",
			keysToInsert:    []string{"apple", "banana", "cherry"},
			keyToDelete:     "grape",
			expectedDeleted: false,
			expectedExists: map[string]bool{
				"apple":  true,
				"banana": true,
				"cherry": true,
			},
		},
		{
			name:            "Delete from empty hash table",
			keysToInsert:    []string{},
			keyToDelete:     "anything",
			expectedDeleted: false,
			expectedExists:  map[string]bool{},
		},
		{
			name:            "Delete key with collision",
			keysToInsert:    []string{"abc", "cba"}, // These should collide with our hash function
			keyToDelete:     "abc",
			expectedDeleted: true,
			expectedExists: map[string]bool{
				"abc": false,
				"cba": true,
			},
		},
		{
			name:            "Delete already deleted key",
			keysToInsert:    []string{"apple", "banana", "cherry"},
			keyToDelete:     "banana", // We'll delete this twice
			expectedDeleted: true,     // First deletion should succeed
			expectedExists: map[string]bool{
				"apple":  true,
				"banana": false,
				"cherry": true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Initialize a new hash table for each test case
			ht := InitHashTable()

			// Insert the keys for this test case
			for _, key := range test.keysToInsert {
				ht.Insert(key)
			}

			// Perform the delete operation
			deleted := ht.Delete(test.keyToDelete)

			// Check if the delete result matches the expected result
			if deleted != test.expectedDeleted {
				t.Errorf("Delete for '%s' returned %t, want %t",
					test.keyToDelete, deleted, test.expectedDeleted)
			}

			// For "Delete already deleted key" test, try deleting again and check it fails
			if test.name == "Delete already deleted key" {
				deleted = ht.Delete(test.keyToDelete)
				if deleted {
					t.Errorf("Second delete for '%s' should return false, got true",
						test.keyToDelete)
				}
			}

			// Check if each key exists or not as expected after deletion
			for key, shouldExist := range test.expectedExists {
				exists := ht.Search(key)
				if exists != shouldExist {
					t.Errorf("After deletion, Search for '%s' returned %t, want %t",
						key, exists, shouldExist)
				}
			}
		})
	}
}

func TestHashCollisionHandling(t *testing.T) {
	// This test verifies that the hash table correctly handles collisions
	// by testing various collision scenarios in a table-driven approach

	// Helper function to find keys that hash to the same value
	findCollidingKeys := func() (string, string, bool) {
		// Common patterns that might collide in simple hash functions
		candidates := []string{
			"abc", "cba",
			"Aa", "BB", // ASCII values: 'A'=65, 'a'=97, 'B'=66, 'B'=66, so 65+97=66+66=132
			"AaA", "BBB",
			"ab", "ba",
			"ac", "ca",
		}

		for i := 0; i < len(candidates); i++ {
			for j := i + 1; j < len(candidates); j++ {
				if hash(candidates[i]) == hash(candidates[j]) {
					return candidates[i], candidates[j], true
				}
			}
		}

		return "", "", false
	}

	key1, key2, found := findCollidingKeys()
	if !found {
		t.Skip("Could not find any colliding keys for testing")
	}

	hashVal := hash(key1)
	t.Logf("Found colliding keys '%s' and '%s' that hash to %d", key1, key2, hashVal)

	tests := []struct {
		name          string
		setup         func(*HashTable)
		action        func(*HashTable) bool
		expectedFound map[string]bool
	}{
		{
			name: "Insert two colliding keys",
			setup: func(ht *HashTable) {
				ht.Insert(key1)
				ht.Insert(key2)
			},
			action: func(ht *HashTable) bool {
				return true // No action, just setup
			},
			expectedFound: map[string]bool{
				key1: true,
				key2: true,
			},
		},
		{
			name: "Delete first colliding key",
			setup: func(ht *HashTable) {
				ht.Insert(key1)
				ht.Insert(key2)
			},
			action: func(ht *HashTable) bool {
				return ht.Delete(key1)
			},
			expectedFound: map[string]bool{
				key1: false,
				key2: true,
			},
		},
		{
			name: "Delete second colliding key",
			setup: func(ht *HashTable) {
				ht.Insert(key1)
				ht.Insert(key2)
			},
			action: func(ht *HashTable) bool {
				return ht.Delete(key2)
			},
			expectedFound: map[string]bool{
				key1: true,
				key2: false,
			},
		}, {
			name: "Insert two colliding keys and verify independence",
			setup: func(ht *HashTable) {
				// Just insert the two keys we already know collide
				ht.Insert(key1)
				ht.Insert(key2)
				t.Logf("Inserted colliding keys '%s' and '%s'", key1, key2)
			},
			action: func(ht *HashTable) bool {
				// Delete key2 and return the result
				return ht.Delete(key2)
			},
			expectedFound: map[string]bool{
				key1: true,  // key1 should still exist
				key2: false, // key2 should be deleted
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Initialize a new hash table for each test case
			ht := InitHashTable()

			// Run the setup
			test.setup(ht)

			// Perform the action
			result := test.action(ht)

			// For delete tests, check the result was as expected
			if strings.Contains(test.name, "Delete") {
				expected := true
				if result != expected {
					t.Errorf("Delete operation returned %t, want %t", result, expected)
				}
			}

			// Check if each key exists or not as expected after the action
			for key, shouldExist := range test.expectedFound {
				exists := ht.Search(key)
				if exists != shouldExist {
					t.Errorf("Search for '%s' returned %t, want %t", key, exists, shouldExist)
				}
			}
		})
	}
}

// We need to modify the approach since Go doesn't allow adding fields to structs at runtime
// Let's use a different method for tracking the third key
