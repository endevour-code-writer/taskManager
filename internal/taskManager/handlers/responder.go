package handlers

import (
	"encoding/json"
	"net/http"
)

type Responder struct {
}

func NewResponder() Responder {
	return Responder{}
}

func (responder Responder) responseError(w http.ResponseWriter, statusCode int, data interface{}) {
	responder.responseJSON(w, statusCode, data)
}

func (responder Responder) responseJSON(w http.ResponseWriter, status int, data interface{}) {
	response, err := json.Marshal(data)

	if err != nil {
		//TODO: add log
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err = w.Write(response); err != nil {
		//TODO: add log
	}
}
