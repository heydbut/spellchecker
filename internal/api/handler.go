package api

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"spellchecker/internal/spellchecker"
)

type handler struct {
	speller *spellchecker.SpellChecker
	tmpl    *template.Template
}

// NewHandler configures the HTTP routes for the API
func NewHandler(speller *spellchecker.SpellChecker, templateFile string) (http.Handler, error) {
	funcMap := template.FuncMap{
		"safeHTML": func(html string) template.HTML {
			return template.HTML(html)
		},
	}

	tmpl, err := template.New("index.tmpl").Funcs(funcMap).ParseFiles(templateFile)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to parse template file"))
	}

	h := &handler{
		speller: speller,
		tmpl:    tmpl,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Index)
	mux.HandleFunc("/check", h.Check)

	return mux, nil
}

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	err := h.tmpl.Execute(w, sampleCheckTextData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Failed to execute template in Index: %v", err)
		return
	}
	log.Printf("Served index page")
}

func (h *handler) Check(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusFound)
		log.Printf("Redirected to index page")
		return
	}

	parseErr := r.ParseForm()
	if parseErr != nil {
		http.Error(w, parseErr.Error(), http.StatusInternalServerError)
		log.Printf("Failed to parse form in Check: %v", parseErr)
		return
	}
	text := r.FormValue("text")

	processedText, incorrectWords := h.speller.CheckText(text)
	data := templateData{
		Text:           text,
		ProcessedText:  processedText,
		IncorrectWords: incorrectWords,
	}

	tmplErr := h.tmpl.Execute(w, data)
	if tmplErr != nil {
		http.Error(w, tmplErr.Error(), http.StatusInternalServerError)
		log.Printf("Failed to execute template in Check: %v", tmplErr)
		return
	}
	log.Printf("Checked text: %s", text)
}
