package handlers

import (
	"net/http"

	"github.com/kuzin57/OnlineShop/pkg/auth"
	"github.com/kuzin57/OnlineShop/pkg/db"
)

type myOrdersPageHandler struct {
	path        string
	htmlSources []string
	repo        *db.Repository
}

func AddMyOrdersPageHandler(
	router *http.ServeMux,
	conf PagesConfig,
	repo *db.Repository,
) PageHandler {
	handler := &myOrdersPageHandler{
		path:        conf.MyOrders.Path,
		htmlSources: conf.MyOrders.Templates,
		repo:        repo,
	}

	router.HandleFunc(
		conf.MyOrders.Path,
		handler.Handle,
	)

	return handler
}

func (h *myOrdersPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := auth.CheckAuthorized(w, r); err != nil {
			w.Header().Add("Authorized", "false")
			executeTemplates(w, h.htmlSources)
			return
		}

		getOrders := r.Header[http.CanonicalHeaderKey(getOrdersHeader)]
		w.Header().Add("Authorized", "true")
		if getOrders[0] != "yes" {
			executeTemplates(w, h.htmlSources)
			return
		}

		response := Response{}

		email := r.Header[http.CanonicalHeaderKey(emailHeader)]
		if email == nil {
			response.Status = http.StatusBadRequest
			response.Description = "No email header provided"
			sendResponse(w, response)
			return
		}

		orders, err := h.repo.GetUserOrders(email[0])
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Description = err.Error()
			sendResponse(w, response)
			return
		}

		response.Orders = make([]db.Order, len(orders))
		for i, order := range orders {
			tmp := *order
			response.Orders[i] = tmp
		}

		sendResponse(w, response)
	}
}
