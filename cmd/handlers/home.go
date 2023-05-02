package handlers

import (
	"net/http"

	"github.com/kuzin57/OnlineShop/cmd/auth"
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
		if err := auth.CheckAuthorized(w, r); err != nil {
			w.Header().Add("Authorized", "false")
			executeTemplates(w, h.htmlSources)
			return
		}

		w.Header().Add("Authorized", "true")
		executeTemplates(w, h.htmlSources)
	}
}
