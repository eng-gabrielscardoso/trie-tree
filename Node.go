package main

type Node struct {
	char     string
	children [26]*Node
}

func NewNode(char string) *Node {
	node := &Node{char: char}

	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}

	return node
}
