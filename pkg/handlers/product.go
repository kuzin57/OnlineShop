package handlers

import (
	"net/http"
	"strconv"

	"github.com/kuzin57/OnlineShop/pkg/auth"
	"github.com/kuzin57/OnlineShop/pkg/db"
)

type productPageHandler struct {
	path        string
	htmlSources []string
	repo        *db.Repository
}

func AddProductPageHandler(
	router *http.ServeMux,
	conf PagesConfig,
	repo *db.Repository,
) PageHandler {
	handler := &productPageHandler{
		path:        conf.Product.Path,
		htmlSources: conf.Product.Templates,
		repo:        repo,
	}

	router.HandleFunc(
		conf.Product.Path,
		handler.Handle,
	)

	return handler
}

func (h *productPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := auth.CheckAuthorized(w, r); err != nil {
			w.Header().Add("Authorized", "false")
			executeTemplates(w, h.htmlSources)
			return
		}

		w.Header().Add("Authorized", "true")
		executeTemplates(w, h.htmlSources)

		response := Response{
			Status: http.StatusOK,
		}

		sendResponse(w, response)

	case http.MethodPost:
		response := Response{}

		if err := auth.CheckAuthorized(w, r); err != nil {
			w.Header().Add("Authorized", "false")
		} else {
			w.Header().Add("Authorized", "true")
		}

		productID := r.Header.Get("Product-ID")
		id, err := strconv.Atoi(productID)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Description = err.Error()
			sendResponse(w, response)
			return
		}

		product, err := h.repo.GetProductDetailedInfo(id)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Description = err.Error()
			sendResponse(w, response)
			return
		}

		response.Products = append(response.Products, *product)
		sendResponse(w, response)
	}
}