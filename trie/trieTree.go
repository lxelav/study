package main

import "fmt"

type Node struct {
	Data        string
	Children    map[rune]*Node
	isEndOfWord bool
}

type TrieTree struct {
	Root *Node
}

func NewTrieTree() *TrieTree {
	return &TrieTree{Root: &Node{Children: make(map[rune]*Node)}}
}

func (tree *TrieTree) Insert(data string) {
	insert(tree.Root, data)
}

func insert(node *Node, data string) {
	currentNode := node
	for _, char := range data {
		if _, ok := currentNode.Children[char]; !ok {
			currentNode.Children[char] = &Node{
				string(char),
				make(map[rune]*Node),
				false,
			}
		}
		currentNode = currentNode.Children[char]
	}
	currentNode.isEndOfWord = true
}

func (tree *TrieTree) Search(data string) bool {
	return search(tree.Root, data)
}

func search(node *Node, data string) bool {
	currentNode := node

	for _, char := range data {
		if _, exists := currentNode.Children[char]; !exists {
			return false
		}
		currentNode = currentNode.Children[char]
	}
	return currentNode.isEndOfWord
}

func (tree *TrieTree) SearchPrefix(data string) bool {
	return searchPrefix(tree.Root, data)
}

func searchPrefix(node *Node, data string) bool {
	currentNode := node

	for _, char := range data {
		if _, ok := currentNode.Children[char]; !ok {
			return false
		}
		currentNode = currentNode.Children[char]
	}
	return true
}

func (tree *TrieTree) Remove(data string) bool {
	return remove(tree.Root, data, 0)
}

func remove(node *Node, data string, index int) bool {
	currentNode := node
	if index == len(data) {
		if !currentNode.isEndOfWord {
			return false
		}
		currentNode.isEndOfWord = false
		return len(currentNode.Children) == 0
	}

	char := rune(data[index])
	if _, ok := currentNode.Children[char]; !ok {
		return false
	}

	currentNode = currentNode.Children[char]
	if remove(currentNode, data, index+1) {
		delete(currentNode.Children, char)
		return len(currentNode.Children) == 0 && !currentNode.isEndOfWord
	}

	return false
}

func main() {
	trie := NewTrieTree()

	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("hell")
	trie.Insert("helium")

	fmt.Println(trie.Search("hello"))       // true
	fmt.Println(trie.Search("hell"))        // true
	fmt.Println(trie.Search("world"))       // true
	fmt.Println(trie.Search("helium"))      // true
	fmt.Println(trie.SearchPrefix("helix")) // false
	fmt.Println(trie.Search("worl"))        // false
	fmt.Println(trie.SearchPrefix("worl"))  // false
	fmt.Println(trie.SearchPrefix("worl"))  // false
	fmt.Println(trie.SearchPrefix("he"))    // false
}
