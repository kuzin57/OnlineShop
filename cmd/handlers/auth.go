package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/kuzin57/OnlineShop/cmd/auth"
	"github.com/kuzin57/OnlineShop/cmd/db"
	"github.com/kuzin57/OnlineShop/cmd/services"
)

type authPageHandler struct {
	htmlSources []string
	authService services.Authorization
	path        string
}

func AddAuthPageHandler(
	router *http.ServeMux,
	conf PagesConfig,
	postgres *db.AuthPostgres,
) PageHandler {
	handler := &authPageHandler{
		htmlSources: conf.Auth.Templates,
		path:        conf.Auth.Path,
	}

	router.HandleFunc(
		conf.Auth.Path,
		handler.Handle,
	)

	switch conf.Auth.AuthType {
	case "postgres":
		handler.authService = auth.NewAuthService(postgres)
	default:
		panic("Unknown AuthType")
	}

	return handler
}

func (h *authPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ts, err := template.ParseFiles(h.htmlSources...)

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

		response := Response{}

		if err := json.Unmarshal(body, &user); err != nil {
			response.Status = http.StatusBadRequest
			response.Description = err.Error()
			sendResponse(w, response)
			return
		}

		token, err := h.authService.GenerateToken(user.Email, user.Password)
		if err != nil {
			response.Status = http.StatusForbidden
			response.Description = err.Error()
			sendResponse(w, response)
			return
		}

		response.Token = token
		sendResponse(w, response)
	}
}
