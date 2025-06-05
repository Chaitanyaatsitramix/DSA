package trie

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode // Map of child nodes
	isEnd    bool               // Marks if this node represents end of a word
	value    interface{}        // Optional value associated with the word
	count    int                // Number of words that pass through this node
}

// Trie represents a prefix tree data structure
type Trie struct {
	root *TrieNode
	size int // Total number of words in the trie
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// Insert adds a word to the trie
func (t *Trie) Insert(word string) {
	node := t.root
	node.count++

	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[ch]
		node.count++
	}

	if !node.isEnd {
		node.isEnd = true
		t.size++
	}
}

// InsertWithValue adds a word with an associated value
func (t *Trie) InsertWithValue(word string, value interface{}) {
	node := t.root
	node.count++

	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[ch]
		node.count++
	}

	if !node.isEnd {
		t.size++
	}
	node.isEnd = true
	node.value = value
}

// Search returns true if the word exists in the trie
func (t *Trie) Search(word string) bool {
	node := t.searchNode(word)
	return node != nil && node.isEnd
}

// SearchWithValue returns the value associated with the word and true if found
func (t *Trie) SearchWithValue(word string) (interface{}, bool) {
	node := t.searchNode(word)
	if node != nil && node.isEnd {
		return node.value, true
	}
	return nil, false
}

// searchNode returns the node for the last character of the word
func (t *Trie) searchNode(word string) *TrieNode {
	node := t.root

	for _, ch := range word {
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchNode(prefix) != nil
}

// CountPrefix returns the number of words that start with the given prefix
func (t *Trie) CountPrefix(prefix string) int {
	node := t.searchNode(prefix)
	if node == nil {
		return 0
	}
	return node.count
}

// Size returns the total number of words in the trie
func (t *Trie) Size() int {
	return t.size
}

// FindAllWithPrefix returns all words that start with the given prefix
func (t *Trie) FindAllWithPrefix(prefix string) []string {
	node := t.searchNode(prefix)
	if node == nil {
		return nil
	}

	words := make([]string, 0)
	t.findAllWithPrefixUtil(node, prefix, &words)
	return words
}

// findAllWithPrefixUtil is a recursive helper function for FindAllWithPrefix
func (t *Trie) findAllWithPrefixUtil(node *TrieNode, prefix string, words *[]string) {
	if node.isEnd {
		*words = append(*words, prefix)
	}

	for ch, child := range node.children {
		t.findAllWithPrefixUtil(child, prefix+string(ch), words)
	}
}

// Delete removes a word from the trie
// Returns true if the word was found and deleted
func (t *Trie) Delete(word string) bool {
	if len(word) == 0 {
		return false
	}
	return t.deleteUtil(t.root, word, 0)
}

// deleteUtil is a recursive helper function for Delete
func (t *Trie) deleteUtil(node *TrieNode, word string, depth int) bool {
	if node == nil {
		return false
	}

	// If we've reached the end of the word
	if depth == len(word) {
		if !node.isEnd {
			return false
		}
		node.isEnd = false
		t.size--
		return len(node.children) == 0
	}

	ch := rune(word[depth])
	child := node.children[ch]

	shouldDeleteChild := t.deleteUtil(child, word, depth+1)

	if shouldDeleteChild {
		delete(node.children, ch)
		node.count--
		return len(node.children) == 0 && !node.isEnd
	}

	node.count--
	return false
}
