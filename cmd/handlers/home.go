package handlers

import (
<<<<<<< HEAD
<<<<<<< HEAD
=======
	"fmt"
	"html/template"
>>>>>>> 35fe851 (made some changes)
	"net/http"

	"github.com/kuzin57/OnlineShop/cmd/auth"
)

<<<<<<< HEAD
=======
const (
	tokenHeader = "Token"
=======
	"net/http"

	"github.com/kuzin57/OnlineShop/cmd/auth"
>>>>>>> 573a019 (finished with authorization, started with password recovery)
)

>>>>>>> 35fe851 (made some changes)
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
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 573a019 (finished with authorization, started with password recovery)
		if err := auth.CheckAuthorized(w, r); err != nil {
			w.Header().Add("Authorized", "false")
			executeTemplates(w, h.htmlSources)
			return
<<<<<<< HEAD
		}

		w.Header().Add("Authorized", "true")
		executeTemplates(w, h.htmlSources)
=======
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
>>>>>>> 35fe851 (made some changes)
=======
		}

		w.Header().Add("Authorized", "true")
		executeTemplates(w, h.htmlSources)
>>>>>>> 573a019 (finished with authorization, started with password recovery)
	}
}
