package main

import "fmt"

func main() {
	myTrie := NewTrie()

	myTrie.Insert("T")
	myTrie.Insert("R")
	myTrie.Insert("I")
	myTrie.Insert("E")
	myTrie.Insert("F")
	myTrie.Insert("G")
	myTrie.Insert("A")
	myTrie.Insert("R")
	myTrie.Insert("R")

	keyword := "TRIE"

	result := myTrie.SearchWord(keyword)

	if result {
		fmt.Println("I've found", keyword)
	} else {
		fmt.Println(keyword, "is not in the trie")
	}
}
