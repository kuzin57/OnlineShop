package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/kuzin57/OnlineShop/pkg/auth"
	"github.com/kuzin57/OnlineShop/pkg/db"
)

type settingsPageHandler struct {
	path        string
	htmlSources []string
	repo        *db.Repository
}

const (
	cookieEmail = "email"
)

// This function adds handler function to http router
func AddSettingsPageHandler(router *http.ServeMux, conf PagesConfig, repo *db.Repository) PageHandler {
	handler := &settingsPageHandler{
		path:        conf.Settings.Path,
		htmlSources: conf.Settings.Templates,
		repo:        repo,
	}

	router.HandleFunc(
		conf.Settings.Path,
		handler.Handle,
	)

	return handler
}

func (h *settingsPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := auth.CheckAuthorized(w, r); err != nil {
			w.Header().Add("Authorized", "false")
		} else {
			w.Header().Add("Authorized", "true")
		}

		cookie, err := r.Cookie(cookieEmail)
		if err != nil {
			w.Header().Add("Status", strconv.Itoa(http.StatusInternalServerError))
			return
		}

		user, err := h.repo.GetUserByEmail(strings.ReplaceAll(cookie.Value, "%40", "@"))
		if err != nil {
			w.Header().Add("Status", strconv.Itoa(http.StatusInternalServerError))
			return
		}

		w.Header().Add("Status", strconv.Itoa(http.StatusOK))
		w.Header().Add("Email", user.Email)
		w.Header().Add("Firstname", user.Firstname)
		w.Header().Add("Surname", user.Surname)
		w.Header().Add("Birthdate", user.Birthday)
		w.Header().Add("Phone-Number", user.PhoneNumber)
		executeTemplates(w, h.htmlSources)

	case http.MethodPost:
		user := &db.User{}
		response := Response{}
		body := make([]byte, 1000)
		bytes, _ := r.Body.Read(body)
		body = body[:bytes]

		if err := json.Unmarshal(body, user); err != nil {
			response.Status = http.StatusInternalServerError
			response.Description = err.Error()
			sendResponse(w, response)
			return
		}

		oldEmail := r.Header[http.CanonicalHeaderKey(emailHeader)]
		if oldEmail == nil {
			response.Status = http.StatusBadRequest
			response.Description = "No email provided"
			sendResponse(w, response)
			return
		}

		if err := h.repo.UpdateUser(user, oldEmail[0]); err != nil {
			response.Status = http.StatusInternalServerError
			response.Description = err.Error()
			sendResponse(w, response)
			return
		}

		response.Status = http.StatusOK
		response.Description = "Data was successfully updated!"
		sendResponse(w, response)
	}
}
