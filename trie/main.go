package main

import "fmt"

type Trie struct {
	Root *TrieNode
}
type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

// NewTrie Initilaizes a trie
func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{
			Children: make(map[rune]*TrieNode),
		},
	}
}

// Insert a word to a trie
func (t *Trie) Insert(word string) {
	node := t.Root
	for _, char := range word {
		if _, ok := node.Children[char]; !ok {
			node.Children[char] = &TrieNode{
				Children: make(map[rune]*TrieNode),
			}
		}
		node = node.Children[char]
	}
	node.IsEnd = true
}

// Search returns true if the word was found
func (t *Trie) Search(word string) bool {
	node := t.Root
	for _, char := range word {
		if _, ok := node.Children[char]; !ok {
			return false
		}
		node = node.Children[char]
	}
	return node.IsEnd
}

// StartsWith evaluates if the word start with a prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.Root
	for _, char := range prefix {
		if _, ok := node.Children[char]; !ok {
			return false
		}
		node = node.Children[char]
	}
	return true
}
func main() {
	trie := NewTrie()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // Output: true
	fmt.Println(trie.Search("app"))     // Output: false
	fmt.Println(trie.StartsWith("app")) // Output: true

	trie.Insert("app")
	fmt.Println(trie.Search("app")) // Output: true
}
