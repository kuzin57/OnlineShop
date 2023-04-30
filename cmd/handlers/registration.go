package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func AddRegistrationPageHandler(router *http.ServeMux, conf PagesConfig) {
	router.HandleFunc(
		conf.Registration.Path,
		htmlSources(conf.Registration.Templates).registrationPageHandler,
	)
}

func (s htmlSources) registrationPageHandler(w http.ResponseWriter, r *http.Request) {
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

	// user := auth.User{}
	if r.Method == "POST" {
		fmt.Println("POST")
	}
}
