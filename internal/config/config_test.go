package config

import (
	"os"
	"testing"
)

func TestLoadConfigReturnsDefaultValuesWhenEnvVarsNotSet(t *testing.T) {
	cfg := LoadConfig()
	if cfg.ServerPort != defaultPort {
		t.Errorf("Expected default port %d, got %d", defaultPort, cfg.ServerPort)
	}
	if cfg.SuggestionDistance != defaultSuggestionDistance {
		t.Errorf("Expected default suggestion distance %d, got %d", defaultSuggestionDistance, cfg.SuggestionDistance)
	}
	if cfg.CSVFile != "data/us_wo.csv" {
		t.Errorf("Expected default CSV file 'data/us_wo.csv', got '%s'", cfg.CSVFile)
	}
	if cfg.TemplateFile != "templates/index.tmpl" {
		t.Errorf("Expected default template file 'templates/index.tmpl', got '%s'", cfg.TemplateFile)
	}
}

func TestLoadConfigReturnsEnvVarValuesWhenSet(t *testing.T) {
	_ = os.Setenv("SERVER_PORT", "9000")
	_ = os.Setenv("SUGGESTION_DISTANCE", "2")
	_ = os.Setenv("CSV_FILE", "data/test.csv")
	_ = os.Setenv("TEMPLATE_FILE", "templates/test.tmpl")

	cfg := LoadConfig()
	if cfg.ServerPort != 9000 {
		t.Errorf("Expected port 9000 from env var, got %d", cfg.ServerPort)
	}
	if cfg.SuggestionDistance != 2 {
		t.Errorf("Expected suggestion distance 2 from env var, got %d", cfg.SuggestionDistance)
	}
	if cfg.CSVFile != "data/test.csv" {
		t.Errorf("Expected CSV file 'data/test.csv' from env var, got '%s'", cfg.CSVFile)
	}
	if cfg.TemplateFile != "templates/test.tmpl" {
		t.Errorf("Expected template file 'templates/test.tmpl' from env var, got '%s'", cfg.TemplateFile)
	}

	_ = os.Unsetenv("SERVER_PORT")
	_ = os.Unsetenv("SUGGESTION_DISTANCE")
	_ = os.Unsetenv("CSV_FILE")
	_ = os.Unsetenv("TEMPLATE_FILE")
}
