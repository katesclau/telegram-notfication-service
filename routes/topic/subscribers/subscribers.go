package subscribers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/katesclau/telegramsvc/routes/route"
	"github.com/katesclau/telegramsvc/utils"
)

func GetRoute(ctx *route.Context) route.Route {
	r := route.NewRoute(ctx, "/topic/{topicName}/subscribers", true)
	methodHandlers := make(map[string]func(ctx *route.Context, w http.ResponseWriter, r *http.Request))
	methodHandlers["GET"] = getHandler
	r.SetMethodHandlers(methodHandlers)
	return *r
}

func getHandler(ctx *route.Context, w http.ResponseWriter, r *http.Request) {
	topicName := mux.Vars(r)["topicName"]

	subscribers := ctx.DB.GetTopic(topicName).Subscribers
	utils.BuildResponse(w, r, subscribers, http.StatusOK)
}
