package topic

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/katesclau/telegramsvc/routes/not_found"
	"github.com/katesclau/telegramsvc/routes/route"
	"github.com/katesclau/telegramsvc/utils"
)

func GetRoute(ctx *route.Context) route.Route {
	r := route.NewRoute(ctx, "/topic/{topicName}", true)
	methodHandlers := make(map[string]func(ctx *route.Context, w http.ResponseWriter, r *http.Request))
	methodHandlers["GET"] = getHandler
	methodHandlers["DELETE"] = deleteHandler
	r.SetMethodHandlers(methodHandlers)
	return *r
}

func getHandler(ctx *route.Context, w http.ResponseWriter, r *http.Request) {
	topicName := mux.Vars(r)["topicName"]
	if topicName == "" {
		topics := ctx.DB.GetTopics()
		utils.BuildResponse(w, r, topics, http.StatusOK)
		return
	}
	topic := ctx.DB.GetTopic(topicName)
	if topic == nil {
		not_found.Handler(w, r)
		return
	}
	utils.BuildResponse(w, r, topic, http.StatusOK)
}

func deleteHandler(ctx *route.Context, w http.ResponseWriter, r *http.Request) {
	topicName := mux.Vars(r)["topicName"]

	if topicName != "" {
		ctx.DB.DeleteTopic(topicName)
		utils.BuildResponse(w, r, nil, http.StatusNoContent)
		return
	}
	not_found.Handler(w, r)
}
