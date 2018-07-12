package main

import (
	"bytes"
	"fmt"
)

type Node struct {
	key      string
	value    []string
	children map[rune]Node
}

func main() {
	tags := []string{"abc", "abfde", "dbc"}
	node := rebuild(tags)
	fmt.Println(node)
	fmt.Println(find(node, "abc"))
}

func add_child(n Node, char rune) Node {
	var buffer bytes.Buffer
	buffer.WriteString(n.key)
	buffer.WriteString(string(char))

	children := n.children
	childKey := buffer.String()

	child := Node{
		value:    []string{childKey},
		key:      childKey,
		children: make(map[rune]Node),
	}

	children[char] = child
	n.children = children
	return n
}

func has_child(n Node, char rune) bool {
	children := n.children
	if _, ok := children[char]; ok {
		return true
	}
	return false
}

func rebuild(tags []string) Node {
	trie := Node{
		value:    []string{},
		key:      "",
		children: make(map[rune]Node),
	}

	node := trie
	parent := node
	for _, tag := range tags {
		parent = node
		for _, char := range tag {
			if !has_child(node, char) {
				node = add_child(node, char)
			}
			node = node.children[char]
		}
		node = parent
	}

	return node
}

func find(n Node, word string) []string {
	parent := n
	for _, char := range word {
		if has_child(parent, char) {
			parent = parent.children[char]
		} else {
			return []string{}
		}
	}
	return parent.value
}
