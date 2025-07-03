# Trie

A Trie (pronounced "try") is an efficient information retrieval data structure. It's commonly used for storing and searching strings in a space-efficient manner. The name comes from the word "re**trie**val", as the main purpose of a trie is to facilitate fast retrieval of information.

Video tutorials:

- <https://www.youtube.com/watch?v=H-6-8_p88r0>

## Structure

- `trie/`: Package implementing the trie data structure
  - `trie.go`: Core implementation of the `Node` and `Trie` types
  - `trie_test.go`: Unit tests for the trie implementation
- `cmd/`: Command-line demo application
  - `main.go`: Demo program showing trie operations

## Features

- Insert words into the trie
- Search for words in the trie
- Check if any word starts with a given prefix
- Delete words from the trie
- Count total number of words in the trie
- List all words stored in the trie

## Time Complexity

| Operation       | Average Case | Worst Case |
|-----------------|--------------|------------|
| Insert          | O(m)         | O(m)       |
| Search          | O(m)         | O(m)       |
| Delete          | O(m)         | O(m)       |
| StartsWith      | O(p)         | O(p)       |

Where:

- m is the length of the word being processed
- p is the length of the prefix being checked

## Space Complexity

Space complexity for a trie is O(n×m), where n is the number of words and m is the average length of the words.

## Usage

Run the demo program to see the trie in action:

```bash
cd cmd
go run main.go
```

## Common Use Cases

- Autocomplete and predictive text
- Spell checking
- IP routing (longest prefix matching)
- Word games (finding all valid words)
- Dictionary implementations
