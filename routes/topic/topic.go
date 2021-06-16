package topic

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/utils"
)

var DB *db.DBClient

type TopicInput struct {
	Name        string
	Subscribers []string
}

func GetMethods(db *db.DBClient) map[string]func(w http.ResponseWriter, r *http.Request) {
	DB = db
	methods := make(map[string]func(w http.ResponseWriter, r *http.Request))
	methods["GET"] = getHandler
	methods["POST"] = postHandler
	methods["DELETE"] = deleteHandler
	return methods
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Check if single or multiple
	key := utils.KeyFromPath(r.URL.Path, 2)

	if key != "" {
		topic := DB.GetTopic(key)
		buildResponse(w, r, topic, http.StatusOK)
		return
	}
	topics := DB.GetTopics()
	buildResponse(w, r, topics, http.StatusOK)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var topicInput TopicInput
	err := decoder.Decode(&topicInput)
	if err != nil {
		log.Println("Failed to parse body into Topic Input", err)
	}
	topic := DB.CreateTopic(topicInput.Name, []db.Subscriber{})
	// TODO Link Subscribers...
	buildResponse(w, r, topic, http.StatusAccepted)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// Check if single or multiple
	key := utils.KeyFromPath(r.URL.Path, 2)

	if key != "" {
		topic := DB.DeleteTopic(key)
		buildResponse(w, r, topic, http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(http.StatusText(http.StatusNotFound)))
}

func buildResponse(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	body, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Failed to marshal Topic: %+v \n", data)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if _, err := w.Write(body); err != nil {
		fmt.Printf("Failed to write Response\n")
	}
}
