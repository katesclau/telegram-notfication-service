package topic

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/routes/not_found"
	"github.com/katesclau/telegramsvc/utils"
)

var DB *db.DBClient

func GetMethods(db *db.DBClient) map[string]func(w http.ResponseWriter, r *http.Request) {
	DB = db
	methods := make(map[string]func(w http.ResponseWriter, r *http.Request))
	methods["GET"] = getHandler
	methods["DELETE"] = deleteHandler
	log.Println("Topic Methods: ", methods)
	return methods
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	topicName := mux.Vars(r)["topicName"]
	if topicName == "" {
		topics := DB.GetTopics()
		utils.BuildResponse(w, r, topics, http.StatusOK)
		return
	}
	topic := DB.GetTopic(topicName)
	if topic == nil {
		not_found.Handler(w, r)
		return
	}
	utils.BuildResponse(w, r, topic, http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// Check if single or multiple
	key := utils.KeyFromPath(r.URL.Path, 2)

	if key != "" {
		topic := DB.DeleteTopic(key)
		utils.BuildResponse(w, r, topic, http.StatusNoContent)
		return
	}
	not_found.Handler(w, r)
}
