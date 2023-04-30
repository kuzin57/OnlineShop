package handlers

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"encoding/json"
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
	"html/template"
	"log"
	"net/http"
<<<<<<< HEAD
	"time"
>>>>>>> 8a6f0f3 (auth_submit fix)
=======

	"github.com/kuzin57/OnlineShop/cmd/auth"
>>>>>>> 355ec40 (fix test + fix communication between front and back)
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
	}

	switch r.Method {
	case http.MethodGet:
		ts, err := template.ParseFiles(s...)

		if err != nil || ts == nil {
			logError(err, w)
		}

		if err = ts.Execute(w, nil); err != nil {
			logError(err, w)
		}

	case http.MethodPost:
		user := auth.User{}
		body := make([]byte, 1000)
		bytes, _ := r.Body.Read(body)

		body = body[:bytes]
		fmt.Println("body:", string(body))

		w.Header().Set("Content-Type", "text")
		response := Response{Status: http.StatusAccepted, Description: "Success!"}
		if err := auth.Login(&user); err != nil {
			response.Status = http.StatusForbidden
			response.Description = err.Error()
		}

		js, err := json.Marshal(&response)
		if err != nil {
			logError(err, w)
		}

		fmt.Println("response", string(js))
		w.Write([]byte(js))
	}
}
