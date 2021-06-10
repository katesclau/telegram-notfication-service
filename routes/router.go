package routes

import (
	"net/http"

	"scada-lts.org/telegramsvc/webhook"
)

func NewRouter() *http.ServeMux {
	// Endpoints
	router := http.NewServeMux()
	router.HandleFunc("/webhook", webhook.HandleMessage)

	// Static content
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/", fs)

	return router
}
