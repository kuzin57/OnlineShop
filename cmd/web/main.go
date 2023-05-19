package main

import (
	"log"
	"net/http"

	"github.com/kuzin57/OnlineShop/pkg/auth"
	"github.com/kuzin57/OnlineShop/pkg/db"
	"github.com/kuzin57/OnlineShop/pkg/handlers"
)

const (
	pathToConf  = "./cmd/config/page_handlers.yaml"
	staticFiles = "./ui/static"
)

func main() {
	mux := http.NewServeMux()
	pagesConfig := handlers.GetHandlersParameters(pathToConf)

	var pageHandlers []handlers.PageHandler

	repo, err := db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	postgres := db.NewAuthPostgresService(repo)
	messageService := auth.InitServiceEmail()

	pageHandlers = append(pageHandlers, handlers.AddHomePageHandler(mux, pagesConfig))
	pageHandlers = append(pageHandlers, handlers.AddAuthPageHandler(mux, pagesConfig, postgres))
	pageHandlers = append(pageHandlers, handlers.AddRegistrationPageHandler(mux, pagesConfig, postgres))
	pageHandlers = append(pageHandlers, handlers.AddCatalogueHandler(mux, pagesConfig, repo))
	pageHandlers = append(pageHandlers, handlers.AddRecoveryPageHandler(
		mux, pagesConfig,
		repo, postgres, messageService))

	pageHandlers = append(pageHandlers, handlers.AddSettingsPageHandler(mux, pagesConfig, repo))
	pageHandlers = append(pageHandlers, handlers.AddOrderPageHandler(mux, pagesConfig, repo, messageService))
	pageHandlers = append(pageHandlers, handlers.AddMyOrdersPageHandler(mux, pagesConfig, repo))

	fileServer := http.FileServer(http.Dir(staticFiles))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("connecting to http://127.0.0.1:7000")
	err = http.ListenAndServe(":7000", mux)
	log.Fatal(err)
}
