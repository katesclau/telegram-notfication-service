package event

import (
	"net/http"
	"time"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/routes/route"
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

func GetRoute(ctx *route.Context) route.Route {
	r := route.NewRoute(ctx, "/topic/{topicName}/event", true)
	methodHandlers := make(map[string]func(ctx *route.Context, w http.ResponseWriter, r *http.Request))
	methodHandlers["POST"] = postHandler
	r.SetMethodHandlers(methodHandlers)
	return *r
}

func postHandler(ctx *route.Context, w http.ResponseWriter, r *http.Request) {
	// TODO, Send message to all subscribers
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
