package bktree

import "spellchecker/pkg/levenshtein"

type Node struct {
	Word     string
	Children map[int]*Node
}

type BKTree struct {
	Root *Node
}

// NewBKTree A BK-tree is a tree data structure specialized for searching in a metric space,
// and it's particularly well-suited for tasks like spell checking with edit distance algorithms (like Levenshtein).
// BK-trees can efficiently find all words "close" to a given word (within a specified Levenshtein distance).
func NewBKTree() *BKTree {
	return &BKTree{}
}

func (tree *BKTree) Add(word string) {
	if tree.Root == nil {
		tree.Root = &Node{Word: word, Children: make(map[int]*Node)}
		return
	}
	currentNode := tree.Root
	for {
		distance := levenshtein.Distance(currentNode.Word, word)
		if distance == 0 {
			return // Duplicate word
		}
		nextNode, exists := currentNode.Children[distance]
		if !exists {
			currentNode.Children[distance] = &Node{Word: word, Children: make(map[int]*Node)}
			return
		}
		currentNode = nextNode
	}
}

func (tree *BKTree) Search(word string, maxDistance int) []string {
	var results []string
	var search func(node *Node)
	search = func(node *Node) {
		if node == nil {
			return
		}
		distance := levenshtein.Distance(word, node.Word)
		if distance <= maxDistance {
			results = append(results, node.Word)
		}
		lower := distance - maxDistance
		upper := distance + maxDistance
		for d := lower; d <= upper; d++ {
			search(node.Children[d])
		}
	}
	search(tree.Root)
	return results
}
