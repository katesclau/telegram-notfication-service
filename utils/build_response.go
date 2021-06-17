package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func BuildResponse(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal Topic: %+v \n", data)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if _, err := w.Write([]byte(body)); err != nil {
		log.Println("Failed to write Response", err)
	}
}
