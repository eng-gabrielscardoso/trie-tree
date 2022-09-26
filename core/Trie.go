/*
Trie Tree implementation
Project: httsp://github.com/eng-gabrielscardoso/trie-tree
Under MIT Licence
Maintainer: Gabriel Santos Cardoso
*/

/*
Package core provides essential abstractions for building a trie tree
*/
package core

/*
Package strings implements simple functions to manipulate UTF-8 encoded strings.
*/
import s "strings"

/*
This struct provides a Trie tree structure
*/
type Trie struct {
	root *Node
}

/*
This method provides a new Trie tree struture

# For each new Trie tree created the root is initialised with the value "\000" that represents the nil value in UNICODE table

@return &Trie: *Trie
*/
func NewTrie() *Trie {
	root := NewNode("\000")

	return &Trie{root: root}
}

/*
This method is reponsible for insert a new element into the Trie

@param word: string - word to be inserted into the Trie

@return _, error: nil, _ - if something went wrong nothing happens, otherwise thats OK
*/
func (t *Trie) Insert(word string) error {
	current := t.root

	strippedWord := s.ToLower(s.ReplaceAll(word, " ", ""))

	for i := 0; i < len(strippedWord); i++ {
		index := strippedWord[i] - 'a'

		if current.children[index] == nil {
			current.children[index] = NewNode(string(strippedWord[i]))
		}

		current = current.children[index]
	}

	return nil
}

/*
This methods is responsible for search a word into the Trie

@param word: string - word to be searched into the Trie

@return bool - if found, return true, otherwise return false
*/
func (t *Trie) SearchWord(word string) bool {
	strippedWord := s.ToLower(s.ReplaceAll(word, " ", ""))

	current := t.root

	for i := 0; i < len(strippedWord); i++ {
		index := strippedWord[i] - 'a'

		if current == nil || current.children[index] == nil {
			return false
		}
	}

	return true
}
