package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerPort         int
	CSVFile            string
	SuggestionDistance int
	TemplateFile       string
}

const (
	defaultPort               = 8080
	defaultSuggestionDistance = 1
)

func LoadConfig() *Config {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		port = defaultPort
	}

	suggestionDistance, err := strconv.Atoi(os.Getenv("SUGGESTION_DISTANCE"))
	if err != nil {
		suggestionDistance = defaultSuggestionDistance
	}

	csvFile := os.Getenv("CSV_FILE")
	if csvFile == "" {
		csvFile = "data/us_wo.csv"
	}

	templateFile := os.Getenv("TEMPLATE_FILE")
	if templateFile == "" {
		templateFile = "templates/index.tmpl"
	}

	log.Printf("Loaded configuration: port=%d, csvFile=%s, suggestionDistance=%d", port, csvFile, suggestionDistance)
	return &Config{
		ServerPort:         port,
		CSVFile:            csvFile,
		SuggestionDistance: suggestionDistance,
		TemplateFile:       templateFile,
	}
}
