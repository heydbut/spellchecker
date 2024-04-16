package spellchecker

import (
	"testing"
)

func TestNewSpellChecker(t *testing.T) {
	_, err := NewSpellChecker("nonexistent.csv", 2)
	if err == nil {
		t.Errorf("Expected error for nonexistent file, got nil")
	}

	checker, err := NewSpellChecker("fixtures/existing.csv", 2)
	if err != nil {
		t.Errorf("Unexpected error for existing file: %v", err)
	}
	if checker == nil {
		t.Errorf("Expected a SpellChecker instance, got nil")
	}
}

func TestCheck(t *testing.T) {
	checker, _ := NewSpellChecker("fixtures/existing.csv", 2)
	if checker.Check("nonexistentword") {
		t.Errorf("Expected false for nonexistent word, got true")
	}
	if !checker.Check("existingword") {
		t.Errorf("Expected true for existing word, got false")
	}
}

func TestCheckText(t *testing.T) {
	checker, _ := NewSpellChecker("fixtures/existing.csv", 2)
	_, incorrect := checker.CheckText("nonexistentword")
	if len(incorrect) != 1 {
		t.Errorf("Expected one incorrect word, got %d", len(incorrect))
	}
	_, incorrect = checker.CheckText("existingword")
	if len(incorrect) != 0 {
		t.Errorf("Expected no incorrect words, got %d", len(incorrect))
	}
}

func TestAdd(t *testing.T) {
	checker, _ := NewSpellChecker("fixtures/existing.csv", 2)
	checker.Add("newword")
	if !checker.Check("newword") {
		t.Errorf("Expected true for newly added word, got false")
	}
}

func TestSuggest(t *testing.T) {
	checker, _ := NewSpellChecker("fixtures/existing.csv", 2)
	suggestions := checker.Suggest("noexistentword")
	if len(suggestions) == 0 {
		t.Errorf("Expected suggestions for nonexistent word, got none")
	}
	suggestions = checker.Suggest("existingword")
	if len(suggestions) != 1 {
		t.Errorf("Expected 1 suggestion for existing word, got %d", len(suggestions))
	}
}

func TestNormalize(t *testing.T) {
	checker, _ := NewSpellChecker("fixtures/existing.csv", 2)
	if checker.Normalize("Word") != "word" {
		t.Errorf("Expected normalized word to be lowercase")
	}
}
