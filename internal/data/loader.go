package data

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

// LoadWords loads the words from a CSV file into a slice of strings.
func LoadWords(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to open file %s", filename))
	}
	defer f.Close()

	var words []string
	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}
		words = append(words, record[0])
	}
	return words, nil
}
