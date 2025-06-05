# Trie (Prefix Tree) Algorithm

A Trie is a tree-based data structure that provides efficient storage and retrieval of strings. It is particularly useful for implementing dictionaries, auto-complete systems, and prefix-based searches.

## Overview

The Trie structure organizes characters in a way that allows for quick prefix-based operations. Each node in the tree represents a character, and paths from root to leaf represent complete words.

### Time Complexity
- Insert: O(m) where m is the length of the word
- Search: O(m) where m is the length of the word
- Prefix Search: O(p + n) where p is prefix length and n is number of child words

### Space Complexity
- O(ALPHABET_SIZE * m * n) where m is average word length and n is number of words

## Implementation Components

1. **Node Structure**
   - Character storage
   - Child pointers
   - Word completion marker
   - Prefix count tracking

2. **Core Operations**
   - Insert word
   - Search word
   - Prefix search
   - Word deletion

3. **Helper Functions**
   - Node creation
   - Path traversal
   - Word collection
   - Count tracking

## Usage Examples

```go
// Create a new trie
trie := NewTrie()

// Insert words
trie.Insert("cat")
trie.Insert("catch")
trie.Insert("caught")

// Search for words
exists := trie.Search("cat")     // true
exists = trie.Search("cap")      // false

// Get words with prefix
words := trie.GetWordsWithPrefix("cat") // ["cat", "catch", "caught"]
```

## Understanding the Algorithm

### Node Structure:
```
TrieNode {
    children: map[rune]*TrieNode
    isEnd: bool
    count: int
    word: string
}
```

### Word Storage Example:
```
Root
 ↓
 c → a → t (cat*)
         ↓
         c → h (catch*)
         ↓
         u → g → h → t (caught*)
```
(*) marks end of word

## Applications

1. **Auto-complete Systems**
   - Search suggestions
   - Type-ahead features
   - Command completion

2. **Spell Checkers**
   - Word validation
   - Suggestion generation
   - Error correction

3. **Dictionary Implementation**
   - Word lookup
   - Prefix search
   - Word counting

4. **IP Routing Tables**
   - Address lookup
   - Prefix matching
   - Route aggregation

## Edge Cases Handled

- Empty strings
- Non-existent words
- Overlapping prefixes
- Case sensitivity
- Special characters

## Implementation Details

1. **Insert Operation**
   - Character-by-character traversal
   - Node creation as needed
   - End-of-word marking
   - Count increment

2. **Search Operation**
   - Path following
   - End node verification
   - Null checking
   - Count verification

3. **Delete Operation**
   - Path traversal
   - Node removal
   - Reference cleanup
   - Count update

## Common Use Cases

1. **Text Editors**
   - Auto-completion
   - Spell checking
   - Word suggestions

2. **Search Engines**
   - Query suggestions
   - Typeahead search
   - Keyword matching

3. **Games**
   - Word validation
   - Command completion
   - Name lookup

4. **Natural Language Processing**
   - Word tokenization
   - Pattern matching
   - Vocabulary management

## Performance Optimizations

1. **Memory Usage**
   - Compressed nodes
   - Character sharing
   - Path compression

2. **Search Speed**
   - Early termination
   - Path compression
   - Count caching

3. **Insertion Efficiency**
   - Batch processing
   - Memory preallocation
   - Path optimization 