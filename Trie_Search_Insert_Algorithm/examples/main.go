package main

import (
	"fmt"
	"trie/pkg/trie"
)

func main() {
	fmt.Println("=== Trie (Prefix Tree) Algorithm Demonstrations ===\n")

	// Example 1: Dictionary Implementation
	fmt.Println("Example 1: Dictionary Implementation")
	fmt.Println("Building a dictionary with word lookup and prefix search")

	// Create a new trie
	dictionary := trie.NewTrie()

	// Insert some words
	words := []string{
		"cat",
		"catch",
		"caught",
		"dog",
		"dogma",
		"doing",
		"done",
		"zebra",
	}

	fmt.Println("\nInserting words:")
	for _, word := range words {
		dictionary.Insert(word)
		fmt.Printf("Inserted: %s\n", word)
	}

	// Search for words
	fmt.Println("\nSearching for words:")
	searchWords := []string{"cat", "catch", "dog", "zebra", "catchy", "do"}
	for _, word := range searchWords {
		exists := dictionary.Search(word)
		fmt.Printf("Word '%s' exists: %v\n", word, exists)
	}

	// Prefix search
	fmt.Println("\nPrefix searches:")
	prefixes := []string{"cat", "do", "z"}
	for _, prefix := range prefixes {
		words := dictionary.GetWordsWithPrefix(prefix)
		fmt.Printf("Words with prefix '%s': %v\n", prefix, words)
		fmt.Printf("Count of words with prefix '%s': %d\n", prefix, dictionary.GetCount(prefix))
	}

	// Example 2: Auto-complete System
	fmt.Println("\nExample 2: Auto-complete System")
	fmt.Println("Implementing an auto-complete feature")

	// Create a new trie for auto-complete
	autocomplete := trie.NewTrie()

	// Insert programming language names
	languages := []string{
		"python",
		"javascript",
		"java",
		"golang",
		"ruby",
		"rust",
		"php",
		"perl",
		"scala",
		"swift",
	}

	fmt.Println("\nBuilding auto-complete database:")
	for _, lang := range languages {
		autocomplete.Insert(lang)
		fmt.Printf("Added: %s\n", lang)
	}

	// Demonstrate auto-complete
	prefixes = []string{"p", "ja", "go", "r"}
	fmt.Println("\nAuto-complete suggestions:")
	for _, prefix := range prefixes {
		suggestions := autocomplete.GetWordsWithPrefix(prefix)
		fmt.Printf("Typing '%s': %v\n", prefix, suggestions)
	}

	// Example 3: Dynamic Dictionary
	fmt.Println("\nExample 3: Dynamic Dictionary")
	fmt.Println("Demonstrating insertion and deletion")

	// Create a new trie
	dynamic := trie.NewTrie()

	// Insert words
	fmt.Println("\nInserting words:")
	dynamicWords := []string{"hello", "help", "helper", "helping"}
	for _, word := range dynamicWords {
		dynamic.Insert(word)
		fmt.Printf("Inserted: %s\n", word)
	}

	// Show all words
	fmt.Println("\nAll words in dictionary:", dynamic.GetAllWords())

	// Delete some words
	fmt.Println("\nDeleting words:")
	deleteWords := []string{"helper", "hello"}
	for _, word := range deleteWords {
		success := dynamic.Delete(word)
		fmt.Printf("Deleted '%s': %v\n", word, success)
	}

	// Show remaining words
	fmt.Println("\nRemaining words:", dynamic.GetAllWords())

	// Try to delete non-existent word
	success := dynamic.Delete("nothere")
	fmt.Printf("\nTried to delete non-existent word 'nothere': %v\n", success)
}
