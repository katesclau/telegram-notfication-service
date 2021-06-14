package event

import (
	"log"
	"net/http"
	"time"
)

type Event struct {
	Value     string
	Timestamp time.Time
	TopicName string
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Switch Methods
	switch r.Method {
	case "POST":
		handlePost(w, r)
	default:
		log.Println("Method not supported!", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}
}
