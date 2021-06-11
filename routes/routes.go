package routes

import (
	"fmt"
	"log"
	"net/http"

	"scada-lts.org/telegramsvc/routes/middlewares"
	"scada-lts.org/telegramsvc/routes/topic"
	"scada-lts.org/telegramsvc/routes/topic/event"
	"scada-lts.org/telegramsvc/routes/topic/subscribers"
	"scada-lts.org/telegramsvc/routes/webhook"
)

type Route struct {
	Path         string
	RouteHandler func(w http.ResponseWriter, _ *http.Request)
	IsAuthed     bool
}

func (r Route) String() string {
	return fmt.Sprintf("{ Path: %s,IsAuthed: %t }", r.Path, r.IsAuthed)
}

func GetRoutes() *http.ServeMux {
	// Dynamic Routes
	routes := [4]Route{
		{
			"/webhook",
			webhook.HandleMessage,
			false,
		},
		{
			"/topic",
			topic.Handler,
			true,
		},
		{
			"/event",
			event.Handler,
			true,
		},
		{
			"/subscribers",
			subscribers.Handler,
			true,
		},
	}

	// Authed is not working somehow...
	authedLt := middlewares.ChainMiddleware(middlewares.WithLogging, middlewares.WithTracing, middlewares.WithAuthentication)
	lt := middlewares.ChainMiddleware(middlewares.WithLogging, middlewares.WithTracing)
	router := http.NewServeMux()

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
