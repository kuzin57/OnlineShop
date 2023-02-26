package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"./ui/html/home.tmpl",
		"./ui/html/base.tmpl",
		"./ui/html/footer.tmpl",
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func romaPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("romaroma"))
}

func lalalaHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		panic("Invalid argument!")
	}

	fmt.Fprintf(w, "Loading content with id %d...", id)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/roma", romaPageHandler)
	mux.HandleFunc("/roma/lalala", lalalaHandler)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("connecting to http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
