package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kuzin57/OnlineShop/pkg/auth"
	"github.com/kuzin57/OnlineShop/pkg/db"
	"github.com/kuzin57/OnlineShop/pkg/services"
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
		executeTemplates(w, h.htmlSources)

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

		indexDog := strings.Index(user.Email, "@")
		response.UserName = user.Email[:indexDog]

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
