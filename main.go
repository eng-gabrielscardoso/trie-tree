/*
Trie Tree implementation
Project: httsp://github.com/eng-gabrielscardoso/trie-tree
Under MIT Licence
Maintainer: Gabriel Santos Cardoso
*/

/*
Package main provides the entrypoint for the application
*/
package main

import (
	/*
		Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
	*/
	f "fmt"
	/*
		Package strings implements simple functions to manipulate UTF-8 encoded strings.
	*/
	s "strings"

	/*
		Package core provides essential abstractions for building a trie tree
	*/
	c "github.com/eng-gabrielscardoso/trie-tree/core"
)

/*
This function is the entrypoint of application and is responsible for initialisation and building of application
*/
func main() {
	myTrie := c.NewTrie()

	chars := []string{"R", "T", "R", "I", "E", "F", "G", "A", "R", "B", "R", "I", "E", "L", "R", "R", "E", "L", "R", "N"}

	for i := 0; i < len(chars); i++ {
		myTrie.Insert(chars[i])
	}

	f.Println("# SEARCHING FOR: TRIE")

	keyword := "TRIE"

	result := myTrie.SearchWord(s.ToLower(keyword))

	if result {
		f.Println("I've found", keyword)
	} else {
		f.Println(keyword, "is not in the trie")
	}

	f.Println("# SEARCHING FOR: E")

	keyword = "E"

	result = myTrie.SearchWord(s.ToLower(keyword))

	if result {
		f.Println("I've found", keyword)
	} else {
		f.Println(keyword, "is not in the trie")
	}

	f.Println("# SEARCHING FOR: GABRIEL")

	keyword = "Gabriel"

	result = myTrie.SearchWord(s.ToLower(keyword))

	if result {
		f.Println("I've found", keyword)
	} else {
		f.Println(keyword, "is not in the trie")
	}

	f.Println("# SEARCHING FOR: LEANDRO")

	keyword = "leandro"

	result = myTrie.SearchWord(s.ToLower(keyword))

	if result {
		f.Println("I've found", keyword)
	} else {
		f.Println(keyword, "is not in the trie")
	}

	myTrie.RemoveWord("E")

	f.Println("# SEARCHING FOR: E (AFTER DELETION)")

	keyword = "E"

	result = myTrie.SearchWord(s.ToLower(keyword))

	if result {
		f.Println("I've found", keyword)
	} else {
		f.Println(keyword, "is not in the trie")
	}

	f.Println("# SEARCHING FOR: F (AFTER E DELETION)")

	keyword = "F"

	result = myTrie.SearchWord(s.ToLower(keyword))

	if result {
		f.Println("I've found", keyword)
	} else {
		f.Println(keyword, "is not in the trie")
	}

	f.Println("# SEARCHING FOR: N (MOST INTERNAL NODE AFTER E DELETION)")

	keyword = "N"

	result = myTrie.SearchWord(s.ToLower(keyword))

	if result {
		f.Println("I've found", keyword)
	} else {
		f.Println(keyword, "is not in the trie")
	}
}
