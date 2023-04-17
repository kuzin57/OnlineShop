package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// This function adds handler function to http router
func AddHomePageHandler(router *http.ServeMux, conf PagesConfig) {
	router.HandleFunc(
		conf.Home.Path,
		htmlSources(conf.Home.Templates).homePageHandler,
	)
}

func (s htmlSources) homePageHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles(s...)

	logError := func(err error, w http.ResponseWriter) {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	if err != nil || ts == nil {
		logError(err, w)
	}

	if err = ts.Execute(w, nil); err != nil {
		logError(err, w)
	}
}
