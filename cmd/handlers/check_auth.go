package handlers

import "net/http"

func checkAuth(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["token"] != nil {

		}
	})
}
