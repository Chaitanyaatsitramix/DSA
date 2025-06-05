# Trie (Prefix Tree) Algorithm

A Trie, also called a prefix tree, is an efficient tree-like data structure for storing and retrieving strings. It excels at tasks involving prefix-based operations like autocomplete, spell checking, and IP routing.

## Overview

The Trie structure provides fast lookups and prefix-based operations by storing characters of strings in nodes, where each path from root to a marked node represents a word in the set.

### Time Complexity
- Insert: O(m) where m is the length of the word
- Search: O(m) where m is the length of the word
- Prefix Search: O(p) where p is the length of the prefix
- Space: O(ALPHABET_SIZE * m * n) where n is number of words

## Implementation Components

1. **Node Structure**
   - Character map for children
   - End of word marker
   - Optional value storage
   - Prefix count tracking

2. **Core Operations**
   - Insert word
   - Search word
   - Prefix search
   - Delete word

3. **Advanced Features**
   - Value association
   - Prefix counting
   - Word enumeration
   - Dynamic updates

## Usage Examples

```go
// Create a new trie
trie := NewTrie()

// Insert words
trie.Insert("cat")
trie.Insert("car")
trie.InsertWithValue("code", "computer instructions")

// Search operations
exists := trie.Search("cat")         // true
hasPrefix := trie.StartsWith("ca")   // true
count := trie.CountPrefix("ca")      // 2
words := trie.FindAllWithPrefix("ca") // ["cat", "car"]
```

## Understanding the Algorithm

### Node Structure:
```
     root
    /    \
   c      d
  /        \
 a          o
/  \         \
t   r         g
```

### Operations:

1. **Insert**
   ```
   "cat" ->  c -> a -> t*
   "car" ->  c -> a -> r*
   (* marks word end)
   ```

2. **Search**
   ```
   "cat" -> follow path -> check end marker
   "ca"  -> follow path -> check end marker
   ```

3. **Prefix Search**
   ```
   "ca" -> follow path -> collect all words below
   ```

## Applications

1. **Autocomplete Systems**
   - Search suggestions
   - Type-ahead features
   - Command completion

2. **Spell Checkers**
   - Word validation
   - Suggestion generation
   - Error correction

3. **IP Routing**
   - Address lookup
   - Prefix matching
   - Route aggregation

4. **Dictionary Implementation**
   - Word storage
   - Definition lookup
   - Prefix-based search

## Edge Cases Handled

- Empty string
- Non-existent words
- Prefix-only matches
- Duplicate insertions
- Case sensitivity

## Implementation Details

1. **Memory Optimization**
   - Compressed nodes
   - Character maps
   - Reference counting

2. **Performance Features**
   - Fast lookups
   - Efficient prefix operations
   - Memory-efficient storage

3. **Flexibility**
   - Generic value storage
   - Dynamic operations
   - Custom character sets

## Common Use Cases

1. **Text Editors**
   - Word completion
   - Syntax highlighting
   - Command suggestions

2. **Search Engines**
   - Query suggestions
   - Partial matching
   - Typeahead search

3. **Games**
   - Word validation
   - Command processing
   - Name matching

4. **Natural Language Processing**
   - Word tokenization
   - Pattern matching
   - Dictionary operations

## Algorithm Variations

1. **Compressed Trie**
   - Merges single-child nodes
   - Reduces memory usage
   - Faster traversal

2. **Ternary Search Tree**
   - Binary tree-like structure
   - More space efficient
   - Slower than standard trie

3. **Radix Tree**
   - Compressed paths
   - Variable-length nodes
   - IP routing optimization 