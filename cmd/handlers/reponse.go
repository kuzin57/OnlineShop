package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Token       string `json:"token"`
}

func sendResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(&response)
	if err != nil {
		logError(err, w)
	}

	w.Write([]byte(js))
}
