package spellchecker

import (
	"errors"
	"fmt"
	"spellchecker/internal/data"
	"spellchecker/pkg/bktree"
	"strings"
)

type SpellChecker struct {
	Dictionary map[string]struct{}
	BKTree     *bktree.BKTree
	Distance   int
}

func NewSpellChecker(csvFilename string, distance int) (*SpellChecker, error) {
	words, err := data.LoadWords(csvFilename)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to load words from %s", csvFilename))
	}
	checker := &SpellChecker{
		Dictionary: make(map[string]struct{}),
		BKTree:     bktree.NewBKTree(),
		Distance:   distance,
	}
	for _, word := range words {
		checker.Add(word)
	}

	return checker, nil
}

func (s *SpellChecker) Check(word string) bool {
	norm := s.Normalize(word)
	_, found := s.Dictionary[norm]
	return found
}

type Incorrect struct {
	Word        string
	Suggestions []string
}

func (s *SpellChecker) CheckText(text string) (string, []Incorrect) {
	splitText := strings.Split(text, " ")
	incorrect := make([]Incorrect, 0)
	var result strings.Builder

	for _, word := range splitText {
		if isSeparator(word) {
			result.WriteString(word)
		} else if s.Check(word) {
			result.WriteString(word + " ")
		} else {
			suggestions := s.Suggest(word)
			incorrect = append(incorrect, Incorrect{Word: word, Suggestions: suggestions})
			// Add HTML markup to highlight the incorrect word with suggestions as a tooltip
			result.WriteString(
				`<mark title="Did you mean ` + strings.Join(
					suggestions, ", ",
				) + `?">` + word + `</mark> `,
			)
		}
	}

	return result.String(), incorrect
}

func (s *SpellChecker) Add(word string) {
	norm := s.Normalize(word)
	s.Dictionary[norm] = struct{}{}
	s.BKTree.Add(norm)
}

func (s *SpellChecker) Suggest(word string) []string {
	norm := s.Normalize(word)
	return s.BKTree.Search(norm, s.Distance)
}

func (s *SpellChecker) Normalize(word string) string {
	return strings.ToLower(word)
}

func isSeparator(r string) bool {
	return r == " " || r == "\n" || r == "\t"
}
