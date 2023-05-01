package handlers

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"encoding/json"
<<<<<<< HEAD
	"net/http"
	"strings"

	"github.com/kuzin57/OnlineShop/cmd/auth"
	"github.com/kuzin57/OnlineShop/cmd/db"
	"github.com/kuzin57/OnlineShop/cmd/services"
=======
=======
	"encoding/json"
>>>>>>> 355ec40 (fix test + fix communication between front and back)
	"fmt"
=======
>>>>>>> 35fe851 (made some changes)
	"html/template"
	"net/http"
<<<<<<< HEAD
	"time"
>>>>>>> 8a6f0f3 (auth_submit fix)
=======

	"github.com/kuzin57/OnlineShop/cmd/auth"
<<<<<<< HEAD
>>>>>>> 355ec40 (fix test + fix communication between front and back)
=======
	"github.com/kuzin57/OnlineShop/cmd/db"
	"github.com/kuzin57/OnlineShop/cmd/services"
>>>>>>> 35fe851 (made some changes)
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
<<<<<<< HEAD
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

<<<<<<< HEAD
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
=======
func (s htmlSources) authPageHandler(w http.ResponseWriter, r *http.Request) {
	logError := func(err error, w http.ResponseWriter) {
		fmt.Println("error")
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
=======
>>>>>>> 35fe851 (made some changes)
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
