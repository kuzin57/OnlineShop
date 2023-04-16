package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func AddCatalogueHandler(router *http.ServeMux, conf PagesConfig) {
	router.HandleFunc(
		conf.Catalogue.Path,
		htmlSources(conf.Catalogue.Templates).cataloguePageHandler,
	)
}

func (s htmlSources) cataloguePageHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles(s...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}

	if err = ts.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}
