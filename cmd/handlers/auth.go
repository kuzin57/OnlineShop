package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
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
	ts, err := template.ParseFiles(s...)

	logError := func(err error, w http.ResponseWriter) {
		fmt.Println("error")
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
	if r.Method == http.MethodPost {
		// body := make([]byte, 1000)
		// bytes, _ := r.Body.Read(body)

		// body = body[:bytes]
		// if err = json.Unmarshal(body, &user); err != nil {
		// 	logError(err, w)
		// }

		// fmt.Println("body:", string(body))

		// w.Header().Set("Content-Type", "text")
		// response := Response{Status: http.StatusAccepted, Description: "Success!"}
		// // if err = auth.Login(&user); err != nil {
		// // 	response.Status = http.StatusForbidden
		// // 	response.Description = err.Error()
		// // }

		// js, err := json.Marshal(&response)
		// if err != nil {
		// 	logError(err, w)
		// }

		// // // json.NewEncoder(w).Encode(&response)
		// // // w.Write([]byte("ahahahah"))
		// // // n, err := w.Write([]byte(js))
		// // // fmt.Println("bytes written", n, err)
		// // // fmt.Println(err)
		// // fmt.Fprintf()
		// // fmt.Fprintf(w, "hahahaha")

		// fmt.Println("response", string(js))
		// w.Write([]byte(js))
		r.ParseMultipartForm(0)

		message := r.FormValue("message")
		fmt.Fprintf(w, "Server %s \n", message+"|"+time.Now().Format(time.RFC1123))
	}
}
