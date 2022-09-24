package main

import "strings"

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	root := NewNode("\000")

	return &Trie{root: root}
}

func (t *Trie) Insert(word string) error {
	current := t.root

	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	for i := 0; i < len(strippedWord); i++ {
		index := strippedWord[i] - 'a'

		if current.children[index] == nil {
			current.children[index] = NewNode(string(strippedWord[i]))
		}

		current = current.children[index]
	}

	return nil
}

func (t *Trie) SearchWord(word string) bool {
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	current := t.root

	for i := 0; i < len(strippedWord); i++ {
		index := strippedWord[i] - 'a'

		if current == nil || current.children[index] == nil {
			return false
		}
	}

	return true
}
