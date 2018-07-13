package main

type Node struct {
	value    []int
	children map[rune]Node
}

func Rebuild(tags map[string]int) Node {
	node := Node{
		value:    []int{},
		children: make(map[rune]Node),
	}

	for tag, id := range tags {
    children := node.children
		for _, char := range tag {
      if child, exist := children[char]; !exist {
        children[char] = Node{
          value: []int{id},
          children: make(map[rune]Node),
        }
      } else {
        child.value = append(child.value, id)
        children[char] = child
      }
      children = children[char].children
		}
	}

	return node
}

func Find(n Node, word string) []int {
  for _, char := range word {
    var exist bool
    n, exist = n.children[char]
    if !exist { return []int{} }
  }
  return n.value
}
