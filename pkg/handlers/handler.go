package handlers

import "net/http"

type PageHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
