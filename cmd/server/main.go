package main

import (
	"fmt"
	"log"
	"net/http"
	"spellchecker/internal/api"
	"spellchecker/internal/config"
	"spellchecker/internal/spellchecker"
)

func main() {
	cfg := config.LoadConfig()
	speller, err := spellchecker.NewSpellChecker(cfg.CSVFile, cfg.SuggestionDistance)
	if err != nil {
		log.Fatalf("Failed to create spell checker: %s", err)
	}

	handler, err := api.NewHandler(speller, cfg.TemplateFile)
	if err != nil {
		log.Fatalf("Failed to create handler: %s", err)
	}

	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Printf("Starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
