package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/kuzin57/OnlineShop/pkg/auth"
	"github.com/kuzin57/OnlineShop/pkg/db"
	"github.com/kuzin57/OnlineShop/pkg/services"
)

const (
	emailHeader       = "email"
	codeHeader        = "code"
	newPasswordHeader = "New-Password"
	maxCode           = 1000000
)

type recoveryPageHandler struct {
	path            string
	htmlSources     []string
	recoveryService *auth.ServiceEmail
	recoveryCode    int
	repo            *db.Repository
	authService     services.Authorization
}

// This function adds handler function to http router
func AddRecoveryPageHandler(
	router *http.ServeMux,
	conf PagesConfig,
	repo *db.Repository,
	postgres *db.AuthPostgres,
) PageHandler {
	handler := &recoveryPageHandler{
		path:            conf.PasswordRecovery.Path,
		htmlSources:     conf.PasswordRecovery.Templates,
		recoveryService: auth.InitServiceEmail(),
		repo:            repo,
		authService:     auth.NewAuthService(postgres),
	}

	router.HandleFunc(
		conf.PasswordRecovery.Path,
		handler.Handle,
	)

	return handler
}

func (h *recoveryPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		executeTemplates(w, h.htmlSources)

	case http.MethodPost:
		defer r.Body.Close()

		response := Response{}

		email := r.Header[http.CanonicalHeaderKey(emailHeader)]
		code := r.Header[http.CanonicalHeaderKey(codeHeader)]
		newPassword := r.Header[http.CanonicalHeaderKey(newPasswordHeader)]

		if email == nil {
			response.Description = "No email provided"
			response.Status = http.StatusExpectationFailed
			sendResponse(w, response)
			return
		}

		if err := h.repo.CheckEmailExists(email[0]); err != nil {
			response.Description = err.Error()
			response.Status = http.StatusForbidden
			sendResponse(w, response)
			return
		}

		if newPassword != nil {
			if err := h.authService.UpdatePassword(email[0], newPassword[0]); err != nil {
				response.Status = http.StatusInternalServerError
				response.Description = err.Error()
				goto send_response
			}

			response.Status = http.StatusOK
			response.Description = "Password successfully changed!"
			goto send_response
		}

		if code != nil {
			codeNum, err := strconv.Atoi(code[0])
			if err != nil {
				response.Status = http.StatusBadRequest
				response.Description = err.Error()
				goto send_response
			}

			if h.recoveryCode != codeNum {
				response.Description = "Incorrect code"
				response.Status = http.StatusForbidden
				goto send_response
			}

			response.Status = http.StatusOK
			response.Description = "Correct code!"

			goto send_response
		}

		h.recoveryCode = generateCode()
		if err := h.recoveryService.SendCode(email[0], h.recoveryCode); err != nil {
			response.Description = err.Error()
			response.Status = http.StatusInternalServerError
			goto send_response
		}

		response.Status = http.StatusOK
		response.Description = "Code was sent!"

	send_response:
		sendResponse(w, response)
		return
	}
}

func generateCode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxCode) % maxCode
}
