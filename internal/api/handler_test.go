package api_test

import (
	"net/http"
	"net/http/httptest"
	"spellchecker/internal/api"
	"spellchecker/internal/spellchecker"
	"strings"
	"testing"
)

func TestNewHandlerReturnsErrorWhenTemplateParseFails(t *testing.T) {
	_, err := api.NewHandler(nil, "nonexistent.tmpl")
	if err == nil {
		t.Errorf("Expected error when template parse fails, got nil")
	}
}

func TestIndexExecutesTemplateSuccessfully(t *testing.T) {
	speller, err := spellchecker.NewSpellChecker("../../data/us_wo.csv", 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	handler, err := api.NewHandler(speller, "../../templates/index.tmpl")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCheckRedirectsWhenNotPost(t *testing.T) {
	speller, err := spellchecker.NewSpellChecker("../../data/us_wo.csv", 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	handler, err := api.NewHandler(speller, "../../templates/index.tmpl")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	req, _ := http.NewRequest("GET", "/check", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}
}

func TestCheckReturnsErrorWhenFormParseFails(t *testing.T) {
	speller, err := spellchecker.NewSpellChecker("../../data/us_wo.csv", 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	handler, err := api.NewHandler(speller, "../../templates/index.tmpl")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	req, _ := http.NewRequest("POST", "/check", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestCheckExecutesTemplateSuccessfully(t *testing.T) {
	speller, err := spellchecker.NewSpellChecker("../../data/us_wo.csv", 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	handler, err := api.NewHandler(speller, "../../templates/index.tmpl")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	req, _ := http.NewRequest("POST", "/check", strings.NewReader("text=word"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
