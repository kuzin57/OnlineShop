package main

import (
	"log"
	"net/http"

	"github.com/kuzin57/OnlineShop/cmd/handlers"
)

const (
	pathToConf  = "./cmd/config/page_handlers.yaml"
	staticFiles = "./ui/static"
)

func main() {
	mux := http.NewServeMux()
	pagesConfig := handlers.GetHandlersParameters(pathToConf)
	handlers.AddHomePageHandler(mux, pagesConfig)

	fileServer := http.FileServer(http.Dir(staticFiles))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("connecting to http://127.0.0.1:7000")
	err := http.ListenAndServe(":7000", mux)
	log.Fatal(err)
}
