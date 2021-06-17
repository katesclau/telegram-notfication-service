package topics

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/utils"
)

var DB *db.DBClient

func GetMethods(db *db.DBClient) map[string]func(w http.ResponseWriter, r *http.Request) {
	DB = db
	methods := make(map[string]func(w http.ResponseWriter, r *http.Request))
	methods["GET"] = getHandler
	methods["POST"] = postHandler
	return methods
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	topics := DB.GetTopics()
	utils.BuildResponse(w, r, topics, http.StatusOK)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var topicInput db.TopicInput
	err := decoder.Decode(&topicInput)
	if err != nil {
		log.Println("Failed to parse body into Topic Input", err)
	}
	log.Println("Topic: ", topicInput)
	topic := DB.CreateTopic(topicInput.Name, topicInput.Subscribers)
	// TODO Link Subscribers...
	utils.BuildResponse(w, r, topic, http.StatusAccepted)
}
