package utils

import (
	"encoding/json"
	"net/http"
)

type responseS struct {
	Message string `json:"message"`
}

func JsonResponse(message string, isSucced bool, rw http.ResponseWriter) {
	if isSucced {
		sendSuccess(message, rw)
	} else {
		sendError(message, rw)
	}
}

func sendSuccess(message string, rw http.ResponseWriter) {
	var r responseS
	r.Message = message
	json.NewEncoder(rw).Encode(r)
}

func sendError(message string, rw http.ResponseWriter) {
	var r responseS
	r.Message = "Error: " + message
	json.NewEncoder(rw).Encode(r)
}
