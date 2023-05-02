package handlers

import (
	"encoding/json"
	"net/http"
<<<<<<< HEAD
<<<<<<< HEAD

	"github.com/kuzin57/OnlineShop/cmd/db"
)

type Response struct {
	Status      int          `json:"status"`
	Description string       `json:"description"`
	Token       string       `json:"token"`
	UserName    string       `json:"userName"`
	Products    []db.Product `json:"products"`
=======
)

type Response struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Token       string `json:"token"`
>>>>>>> 35fe851 (made some changes)
=======

	"github.com/kuzin57/OnlineShop/cmd/db"
)

type Response struct {
	Status      int          `json:"status"`
	Description string       `json:"description"`
	Token       string       `json:"token"`
	UserName    string       `json:"userName"`
	Products    []db.Product `json:"products"`
>>>>>>> 573a019 (finished with authorization, started with password recovery)
}

func sendResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(&response)
	if err != nil {
		logError(err, w)
	}

	w.Write([]byte(js))
}
