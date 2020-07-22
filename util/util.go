package util

import (
	"encoding/json"
	"kiwi/model"
	"net/http"
)

func SendError(w http.ResponseWriter, status int, err model.Error){
	// WriteHeader sends an HTTP response header with the provided
	// status code.
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}){
	json.NewEncoder(w).Encode(data)
}
