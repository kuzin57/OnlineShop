package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func executeTemplates(w http.ResponseWriter, htmlSources []string) {
	ts, err := template.ParseFiles(htmlSources...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}

	if err = ts.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}
