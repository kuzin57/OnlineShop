package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func AddHomePageHandler(router *http.ServeMux, conf PagesConfig) {
	router.HandleFunc(
		conf.Home.Path,
		htmlSources(conf.Home.Templates).homePageHandler,
	)
}

type htmlSources []string

func (s htmlSources) homePageHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles(s...)
	if err != nil || ts == nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
