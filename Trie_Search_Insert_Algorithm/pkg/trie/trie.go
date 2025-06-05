package trie

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode // Map of child nodes
	isEnd    bool               // Marks end of a word
	count    int                // Number of words with this prefix
	word     string             // Complete word (only set for end nodes)
}

// NewTrieNode creates a new TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
		count:    0,
	}
}

// Trie represents the prefix tree
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// Insert adds a word to the trie
func (t *Trie) Insert(word string) {
	node := t.root

	// Traverse/create the path for each character
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
		node.count++
	}

	// Mark the end of the word
	node.isEnd = true
	node.word = word
}

// Search returns true if the word exists in the trie
func (t *Trie) Search(word string) bool {
	node := t.searchNode(word)
	return node != nil && node.isEnd
}

// StartsWith returns true if there is any word that starts with the prefix
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchNode(prefix) != nil
}

// searchNode returns the node for the last character of the word/prefix
func (t *Trie) searchNode(str string) *TrieNode {
	node := t.root

	for _, ch := range str {
		if next, exists := node.children[ch]; exists {
			node = next
		} else {
			return nil
		}
	}

	return node
}

// GetWordsWithPrefix returns all words that start with the given prefix
func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	node := t.searchNode(prefix)
	if node == nil {
		return nil
	}

	words := make([]string, 0)
	t.collectWords(node, &words)
	return words
}

// collectWords recursively collects all words under the given node
func (t *Trie) collectWords(node *TrieNode, words *[]string) {
	if node.isEnd {
		*words = append(*words, node.word)
	}

	for _, child := range node.children {
		t.collectWords(child, words)
	}
}

// Delete removes a word from the trie
func (t *Trie) Delete(word string) bool {
	return t.deleteHelper(t.root, word, 0)
}

// deleteHelper recursively removes a word from the trie
func (t *Trie) deleteHelper(node *TrieNode, word string, depth int) bool {
	// If we've processed all characters
	if depth == len(word) {
		if !node.isEnd {
			return false // Word not found
		}
		node.isEnd = false
		node.word = ""
		return len(node.children) == 0
	}

	ch := rune(word[depth])
	if child, exists := node.children[ch]; exists {
		shouldDeleteChild := t.deleteHelper(child, word, depth+1)
		if shouldDeleteChild {
			delete(node.children, ch)
			return len(node.children) == 0 && !node.isEnd
		}
	}
	return false
}

// GetCount returns the number of words with the given prefix
func (t *Trie) GetCount(prefix string) int {
	node := t.searchNode(prefix)
	if node == nil {
		return 0
	}
	return node.count
}

// GetAllWords returns all words in the trie
func (t *Trie) GetAllWords() []string {
	words := make([]string, 0)
	t.collectWords(t.root, &words)
	return words
}
