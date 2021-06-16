package event

import (
	"net/http"
	"time"

	"github.com/katesclau/telegramsvc/db"
)

type Event struct {
	Value     string
	Timestamp time.Time
	TopicName string
}

var DB *db.DBClient

type EventInput struct {
	Value     string
	TopicName string
}

func GetMethods(db *db.DBClient) map[string]func(w http.ResponseWriter, r *http.Request) {
	DB = db
	methods := make(map[string]func(w http.ResponseWriter, r *http.Request))
	methods["POST"] = postHandler
	return methods
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// TODO, Send message to all subscribers
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
