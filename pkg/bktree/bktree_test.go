package bktree

import (
	"testing"
)

func TestSearchInEmptyTree(t *testing.T) {
	tree := NewBKTree()
	results := tree.Search("test", 1)
	if len(results) != 0 {
		t.Errorf("Expected 0 results, got %d", len(results))
	}
}

func TestSearchExactMatch(t *testing.T) {
	tree := NewBKTree()
	tree.Add("test")
	results := tree.Search("test", 1)
	if len(results) != 1 || results[0] != "test" {
		t.Errorf("Expected ['test'], got %v", results)
	}
}

func TestSearchWithinMaxDistance(t *testing.T) {
	tree := NewBKTree()
	tree.Add("test")
	tree.Add("text")
	results := tree.Search("test", 1)
	if len(results) != 2 || results[0] != "test" || results[1] != "text" {
		t.Errorf("Expected ['test', 'text'], got %v", results)
	}
}

func TestSearchBeyondMaxDistance(t *testing.T) {
	tree := NewBKTree()
	tree.Add("test")
	tree.Add("text")
	results := tree.Search("test", 0)
	if len(results) != 1 || results[0] != "test" {
		t.Errorf("Expected ['test'], got %v", results)
	}
}

func TestAddWordToNonEmptyTree(t *testing.T) {
	tree := NewBKTree()
	tree.Add("test")
	tree.Add("text")
	if tree.Root == nil || tree.Root.Word != "test" {
		t.Errorf("Expected 'test', got '%s'", tree.Root.Word)
	}
	if len(tree.Root.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(tree.Root.Children))
	}
}

func TestAddWordCreatesMultipleChildren(t *testing.T) {
	tree := NewBKTree()
	tree.Add("test")
	tree.Add("text")
	tree.Add("team")
	if len(tree.Root.Children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(tree.Root.Children))
	}
}

func TestAddWordUpdatesExistingChild(t *testing.T) {
	tree := NewBKTree()
	tree.Add("test")
	tree.Add("text")
	tree.Add("tent")
	if len(tree.Root.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(tree.Root.Children))
	}
	if len(tree.Root.Children[1].Children) != 1 {
		t.Errorf("Expected 1 grandchild, got %d", len(tree.Root.Children[1].Children))
	}
}
