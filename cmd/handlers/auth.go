package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/kuzin57/OnlineShop/cmd/auth"
)

type Response struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
}

func AddAuthPageHandler(router *http.ServeMux, conf PagesConfig) {
	router.HandleFunc(
		conf.Auth.Path,
		htmlSources(conf.Auth.Templates).authPageHandler,
	)
}

func (s htmlSources) authPageHandler(w http.ResponseWriter, r *http.Request) {
	logError := func(err error, w http.ResponseWriter) {
		fmt.Println("error")
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	switch r.Method {
	case http.MethodGet:
		ts, err := template.ParseFiles(s...)

		if err != nil || ts == nil {
			logError(err, w)
		}

		if err = ts.Execute(w, nil); err != nil {
			logError(err, w)
		}

	case http.MethodPost:
		user := auth.User{}
		body := make([]byte, 1000)
		bytes, _ := r.Body.Read(body)

		body = body[:bytes]
		fmt.Println("body:", string(body))

		w.Header().Set("Content-Type", "text")
		response := Response{Status: http.StatusAccepted, Description: "Success!"}
		if err := auth.Login(&user); err != nil {
			response.Status = http.StatusForbidden
			response.Description = err.Error()
		}

		js, err := json.Marshal(&response)
		if err != nil {
			logError(err, w)
		}

		fmt.Println("response", string(js))
		w.Write([]byte(js))
	}
}
