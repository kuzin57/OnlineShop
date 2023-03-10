package handlers

import (
	"fmt"
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

type htmlSources []string

func (s htmlSources) homePageHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles(s...)
	fmt.Println("ts", ts)
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
