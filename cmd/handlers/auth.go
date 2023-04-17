package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func AddAuthPageHandler(router *http.ServeMux, conf PagesConfig) {
	router.HandleFunc(
		conf.Auth.Path,
		htmlSources(conf.Auth.Templates).authPageHandler,
	)
}

func (s htmlSources) authPageHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles(s...)

	if err != nil || ts == nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	if err = ts.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	switch r.Method {
	case "POST":
		fmt.Println("POST query received!")
		body := make([]byte, 100)
		size, _ := r.Body.Read(body)
		fmt.Println("size", size)
		fmt.Println("body", string(body))
	case "GET":
		fmt.Println(r.Body)
	}
}
