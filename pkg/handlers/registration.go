package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/kuzin57/OnlineShop/pkg/auth"
	"github.com/kuzin57/OnlineShop/pkg/db"
	"github.com/kuzin57/OnlineShop/pkg/services"
)

const (
	successRegistration = "Successful registration!"
)

type registrationHandler struct {
	path          string
	htmlTemplates []string
	authService   services.Authorization
}

func AddRegistrationPageHandler(
	router *http.ServeMux,
	conf PagesConfig,
	postgres *db.AuthPostgres,
) PageHandler {
	handler := &registrationHandler{
		path:          conf.Registration.Path,
		htmlTemplates: conf.Registration.Templates,
	}

	router.HandleFunc(
		conf.Registration.Path,
		handler.Handle,
	)

	handler.authService = auth.NewAuthService(postgres)
	return handler
}

func (h *registrationHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ts, err := template.ParseFiles(h.htmlTemplates...)
		if err != nil || ts == nil {
			logError(err, w)
		}

		if err = ts.Execute(w, nil); err != nil {
			logError(err, w)
		}
	case http.MethodPost:
		user := db.User{}

		body := make([]byte, 1000)
		bytes, _ := r.Body.Read(body)
		body = body[:bytes]

		if err := json.Unmarshal(body, &user); err != nil {
			logError(err, w)
		}

		w.Header().Set("Content-Type", "application/json")
		response := Response{}

		_, err := h.authService.CreateUser(&user)
		if err != nil {
			response.Status = http.StatusForbidden
			response.Description = err.Error()
			sendResponse(w, response)
			return
		}

		response.Status = http.StatusOK
		response.Description = successRegistration

		sendResponse(w, response)
	}
}
