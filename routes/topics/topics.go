package topics

import (
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
	log.Println("Topics Methods: ", methods)
	return methods
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Topics: ", r.URL.Path)
	topics := DB.GetTopics()
	log.Printf("Topics: %v", topics)
	utils.BuildResponse(w, r, topics, http.StatusOK)
}
