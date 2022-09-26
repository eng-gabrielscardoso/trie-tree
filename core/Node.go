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
This struct provides a node structure

@attr char: string
@attr children: slice
*/
type Node struct {
	char     string
	children [26]*Node
}

/*
This method provides a new node structure

# For each children of node is inserted a nil to allocate memory and complete the tree

@param char: string - The character to be inserted into the node

@return node: *Node
*/
func NewNode(char string) *Node {
	node := &Node{char: char}

	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}

	return node
}
