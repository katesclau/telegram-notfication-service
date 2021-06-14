package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/routes/middlewares"
	"github.com/katesclau/telegramsvc/routes/topic"
	"github.com/katesclau/telegramsvc/routes/topic/event"
	"github.com/katesclau/telegramsvc/routes/topic/subscribers"
	"github.com/katesclau/telegramsvc/routes/webhook"

	"github.com/gorilla/mux"
)

// Doing this https://stackoverflow.com/questions/35038864/how-to-access-global-variables seems stupid, do I have a better option?
var DB *db.DBClient

type Route struct {
	Path         string
	RouteHandler func(w http.ResponseWriter, _ *http.Request)
	IsAuthed     bool
}

func (r Route) String() string {
	return fmt.Sprintf("{ Path: %s,IsAuthed: %t }", r.Path, r.IsAuthed)
}

func GetRoutes() *mux.Router {
	// Dynamic Routes
	routes := [5]Route{
		{
			"/webhook",
			webhook.HandleMessage,
			false,
		},
		{
			"/topic/",
			topic.Handler,
			true,
		},
		{
			"/topic/{topicName}",
			topic.Handler,
			true,
		},
		{
			"/topic/{topicName}/event",
			event.Handler,
			true,
		},
		{
			"/topic/{topicName}/subscribers",
			subscribers.Handler,
			true,
		},
	}

	// Authed is not working somehow...
	authedLt := middlewares.ChainMiddleware(middlewares.WithLogging, middlewares.WithTracing, middlewares.WithAuthentication)
	lt := middlewares.ChainMiddleware(middlewares.WithLogging, middlewares.WithTracing)
	router := mux.NewRouter()

	// Static content
	fs := http.FileServer(http.Dir("./routes/static"))
	router.Handle("/", fs)

	for _, route := range routes {
		log.Printf("Route: %s", route)
		if route.IsAuthed {
			router.HandleFunc(route.Path, authedLt(route.RouteHandler))
			log.Println("Using authedLt")
		} else {
			log.Println("Using lt")
			router.HandleFunc(route.Path, lt(route.RouteHandler))
		}
	}
	return router
}
