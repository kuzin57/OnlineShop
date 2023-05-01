package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	tokenHeader = "Token"
)

type homePageHandler struct {
	path        string
	htmlSources []string
}

// This function adds handler function to http router
func AddHomePageHandler(router *http.ServeMux, conf PagesConfig) PageHandler {
	handler := &homePageHandler{
		path:        conf.Home.Path,
		htmlSources: conf.Home.Templates,
	}

	router.HandleFunc(
		conf.Home.Path,
		handler.Handle,
	)

	return handler
}

func (h *homePageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ts, err := template.ParseFiles(h.htmlSources...)

		if err != nil || ts == nil {
			logError(err, w)
		}

		if err = ts.Execute(w, nil); err != nil {
			logError(err, w)
		}

		// tokenValue := r.Header[tokenHeader]
		fmt.Println("headers", r.Header)
		// fmt.Println("token", tokenValue)
		str := []byte("lalalala")
		w.Header().Set("Content-Type", "application/text")
		w.Write(str)
		// if len(tokenValue) == 0 {

		// }
	}
}
