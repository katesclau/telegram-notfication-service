package topics

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/routes/route"
	"github.com/katesclau/telegramsvc/utils"
)

func GetRoute(ctx *route.Context) route.Route {
	r := route.NewRoute(ctx, "/topic/", true)
	methodHandlers := make(map[string]func(ctx *route.Context, w http.ResponseWriter, r *http.Request))
	methodHandlers["GET"] = getHandler
	methodHandlers["POST"] = postHandler
	r.SetMethodHandlers(methodHandlers)
	return *r
}

func getHandler(ctx *route.Context, w http.ResponseWriter, r *http.Request) {
	topics := ctx.DB.GetTopics()
	utils.BuildResponse(w, r, topics, http.StatusOK)
}

func postHandler(ctx *route.Context, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var topicInput db.TopicInput
	err := decoder.Decode(&topicInput)
	if err != nil {
		log.Println("Failed to parse body into Topic Input", err)
	}
	log.Println("Topic: ", topicInput)
	topic := ctx.DB.CreateTopic(topicInput.Name, topicInput.Subscribers)
	utils.BuildResponse(w, r, topic, http.StatusAccepted)
}
