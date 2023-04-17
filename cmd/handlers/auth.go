package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/kuzin57/OnlineShop/cmd/auth"
)

func AddAuthPageHandler(router *http.ServeMux, conf PagesConfig) {
	router.HandleFunc(
		conf.Auth.Path,
		htmlSources(conf.Auth.Templates).authPageHandler,
	)
}

func (s htmlSources) authPageHandler(w http.ResponseWriter, r *http.Request) {
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

	user := auth.User{}
	if r.Method == "POST" {
		body := make([]byte, 1000)
		bytes, err := r.Body.Read(body)
		if err != nil {
			logError(err, w)
		}

		body = body[:bytes]
		fmt.Println("body", string(body))
		if err = json.Unmarshal(body, &user); err != nil {
			logError(err, w)
		}
	}
}
