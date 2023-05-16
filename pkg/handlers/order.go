package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kuzin57/OnlineShop/pkg/auth"
	"github.com/kuzin57/OnlineShop/pkg/db"
)

type orderPageHandler struct {
	path           string
	htmlSources    []string
	repo           *db.Repository
	messageService *auth.ServiceEmail
}

// This function adds handler function to http router
func AddOrderPageHandler(
	router *http.ServeMux,
	conf PagesConfig,
	repo *db.Repository,
	messageService *auth.ServiceEmail,
) PageHandler {
	handler := &orderPageHandler{
		path:           conf.Order.Path,
		htmlSources:    conf.Order.Templates,
		repo:           repo,
		messageService: messageService,
	}

	router.HandleFunc(
		conf.Order.Path,
		handler.Handle,
	)

	return handler
}

func (h *orderPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := auth.CheckAuthorized(w, r); err != nil {
			w.Header().Add("Authorized", "false")
			executeTemplates(w, h.htmlSources)
			return
		}

		w.Header().Add("Authorized", "true")
		executeTemplates(w, h.htmlSources)

	case http.MethodPost:
		order := &db.Order{}
		response := Response{}
		email := r.Header[http.CanonicalHeaderKey(emailHeader)]

		body := make([]byte, 1000)
		bytes, _ := r.Body.Read(body)
		body = body[:bytes]

		if err := json.Unmarshal(body, order); err != nil {
			response.Status = http.StatusBadRequest
			response.Description = err.Error()
			sendResponse(w, response)
			return
		}

		orderID, err := h.repo.AddOrder(order)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Description = err.Error()
			goto send_response
		}

		order.Id = uint32(orderID)

		if email == nil {
			response.Status = http.StatusBadRequest
			response.Description = "No email header provided"
			goto send_response
		}

		response.Status = http.StatusOK
		response.Description = "Order was successfully created!"
		h.messageService.SendOrderNotification(email[0], order)

	send_response:
		sendResponse(w, response)
	}
}
