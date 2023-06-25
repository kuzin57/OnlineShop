package handlers

import (
	"log"
	"net/http"
)

func logError(err error, w http.ResponseWriter) {
	log.Println(err.Error())
	http.Error(w, "Internal Server Error", 500)
}
