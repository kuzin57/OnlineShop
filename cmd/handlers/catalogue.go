package handlers

import (
	"fmt"
	"net/http"

	"github.com/kuzin57/OnlineShop/cmd/auth"
	"github.com/kuzin57/OnlineShop/cmd/db"
)

type cataloguePageHandler struct {
	path        string
	htmlSources []string
	repo        *db.Repository
}

func AddCatalogueHandler(router *http.ServeMux, conf PagesConfig, repo *db.Repository) PageHandler {
	handler := &cataloguePageHandler{
		htmlSources: conf.Catalogue.Templates,
		path:        conf.Catalogue.Path,
		repo:        repo,
	}

	router.HandleFunc(
		conf.Catalogue.Path,
		handler.Handle,
	)

	return handler
}

func (h *cataloguePageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		executeTemplates(w, h.htmlSources)

	case http.MethodPost:
		response := Response{}

		if err := h.repo.CheckDataBaseAvailable(); err != nil {
			response.Status = http.StatusInternalServerError
			response.Description = err.Error()
			sendResponse(w, response)
			return
		}

		if err := auth.CheckAuthorized(w, r); err != nil {
			w.Header().Add("Authorized", "false")
		} else {
			w.Header().Add("Authorized", "true")
		}

		products, err := h.repo.GetProducts()
		fmt.Println("products", products, "err", err)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Description = err.Error()
			sendResponse(w, response)
		}

		response.Products = make([]db.Product, len(products))
		copy(response.Products, products)
		response.Status = http.StatusOK
		response.Description = "Products successfully got!"
		sendResponse(w, response)
	}
}
