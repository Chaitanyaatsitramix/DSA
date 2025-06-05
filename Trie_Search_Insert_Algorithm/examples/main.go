package main

import (
	"fmt"
	"trie/pkg/trie"
)

func main() {
	fmt.Println("=== Trie (Prefix Tree) Algorithm Demonstrations ===\n")

	// Example 1: Basic Word Dictionary
	fmt.Println("Example 1: Basic Word Dictionary")
	fmt.Println("Demonstrating basic word storage and retrieval")

	dictionary := trie.NewTrie()

	// Insert words
	words := []string{"cat", "car", "cart", "dog", "done"}
	fmt.Println("\nInserting words:", words)
	for _, word := range words {
		dictionary.Insert(word)
	}

	// Search for words
	fmt.Println("\nSearching for words:")
	searchWords := []string{"cat", "car", "cart", "dog", "do", "done", "cats"}
	for _, word := range searchWords {
		exists := dictionary.Search(word)
		fmt.Printf("'%s' exists: %v\n", word, exists)
	}

	// Check prefixes
	fmt.Println("\nChecking prefixes:")
	prefixes := []string{"ca", "do", "tr"}
	for _, prefix := range prefixes {
		hasPrefix := dictionary.StartsWith(prefix)
		count := dictionary.CountPrefix(prefix)
		fmt.Printf("Prefix '%s': exists=%v, word count=%d\n", prefix, hasPrefix, count)

		if hasPrefix {
			words := dictionary.FindAllWithPrefix(prefix)
			fmt.Printf("Words with prefix '%s': %v\n", prefix, words)
		}
	}

	// Example 2: Dictionary with Word Definitions
	fmt.Println("\nExample 2: Dictionary with Word Definitions")
	fmt.Println("Demonstrating storage of words with associated values")

	definitionDict := trie.NewTrie()

	// Insert words with definitions
	definitions := map[string]string{
		"code":      "Instructions for a computer",
		"coding":    "The process of writing computer programs",
		"coder":     "Someone who writes code",
		"algorithm": "A step-by-step procedure for calculations",
	}

	fmt.Println("\nInserting words with definitions:")
	for word, def := range definitions {
		definitionDict.InsertWithValue(word, def)
		fmt.Printf("Added: %s\n", word)
	}

	// Look up definitions
	fmt.Println("\nLooking up definitions:")
	lookupWords := []string{"code", "coding", "coder", "algorithm", "programmer"}
	for _, word := range lookupWords {
		if def, found := definitionDict.SearchWithValue(word); found {
			fmt.Printf("%s: %s\n", word, def)
		} else {
			fmt.Printf("%s: Not found in dictionary\n", word)
		}
	}

	// Example 3: Dynamic Dictionary Operations
	fmt.Println("\nExample 3: Dynamic Dictionary Operations")
	fmt.Println("Demonstrating insertion, deletion, and prefix operations")

	dynamicDict := trie.NewTrie()

	// Step 1: Insert words
	fmt.Println("\nStep 1: Inserting words")
	initialWords := []string{"apple", "app", "apricot", "banana", "band"}
	for _, word := range initialWords {
		dynamicDict.Insert(word)
		fmt.Printf("Added: %s\n", word)
	}

	// Step 2: Show all words with prefix
	fmt.Println("\nStep 2: Words with prefix 'ap':")
	apWords := dynamicDict.FindAllWithPrefix("ap")
	fmt.Printf("Found %d words: %v\n", len(apWords), apWords)

	// Step 3: Delete some words
	fmt.Println("\nStep 3: Deleting words")
	deleteWords := []string{"app", "banana", "notexist"}
	for _, word := range deleteWords {
		deleted := dynamicDict.Delete(word)
		fmt.Printf("Deleting '%s': %v\n", word, deleted)
	}

	// Step 4: Show final dictionary state
	fmt.Println("\nStep 4: Final dictionary state")
	fmt.Printf("Dictionary size: %d\n", dynamicDict.Size())
	fmt.Println("Words with prefix 'ap':", dynamicDict.FindAllWithPrefix("ap"))
	fmt.Println("Words with prefix 'ban':", dynamicDict.FindAllWithPrefix("ban"))
}
