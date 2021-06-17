package subscribers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/utils"
)

var DB *db.DBClient

func GetMethods(db *db.DBClient) map[string]func(w http.ResponseWriter, r *http.Request) {
	DB = db
	methods := make(map[string]func(w http.ResponseWriter, r *http.Request))
	methods["GET"] = getHandler
	return methods
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	topicName := mux.Vars(r)["topicName"]

	subscribers := DB.GetTopic(topicName).Subscribers
	utils.BuildResponse(w, r, subscribers, http.StatusOK)
}
