package levenshtein

import "testing"

func TestDistanceSameStrings(t *testing.T) {
	result := Distance("test", "test")
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestDistanceDifferentStrings(t *testing.T) {
	result := Distance("test", "text")
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestDistanceEmptyStrings(t *testing.T) {
	result := Distance("", "")
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestDistanceOneEmptyString(t *testing.T) {
	result := Distance("test", "")
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

func TestDistanceCaseSensitive(t *testing.T) {
	result := Distance("Test", "test")
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
