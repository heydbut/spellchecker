package data

import (
	"errors"
	"os"
	"testing"
)

func TestLoadWordsFromValidFile(t *testing.T) {
	words, err := LoadWords("fixtures/valid.csv")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(words) == 0 {
		t.Errorf("Expected non-empty slice, got empty")
	}
}

func TestLoadWordsFromNonexistentFile(t *testing.T) {
	_, err := LoadWords("nonexistent.csv")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("Expected file does not exist error, got %v", err)
	}
}

func TestLoadWordsFromEmptyFile(t *testing.T) {
	words, err := LoadWords("fixtures/empty.csv")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(words) != 0 {
		t.Errorf("Expected empty slice, got %v", words)
	}
}
