package routes

import (
	"net/http"

	"scada-lts.org/telegramsvc/routes/event"
	"scada-lts.org/telegramsvc/routes/middlewares"
	"scada-lts.org/telegramsvc/routes/topic"
	"scada-lts.org/telegramsvc/routes/webhook"
)

type Route struct {
	Path         string
	RouteHandler func(w http.ResponseWriter, _ *http.Request)
}

func GetRoutes() *http.ServeMux {
	// Dynamic Routes
	routes := [3]Route{
		{
			"/webhook",
			webhook.HandleMessage,
		},
		{
			"/topic",
			topic.Handler,
		},
		{
			"/event",
			event.Handler,
		},
	}

	lt := middlewares.ChainMiddleware(middlewares.WithLogging, middlewares.WithTracing)
	router := http.NewServeMux()

	// Static content
	fs := http.FileServer(http.Dir("./routes/static"))
	router.Handle("/", fs)

	for _, route := range routes {
		router.HandleFunc(route.Path, lt(route.RouteHandler))
	}
	return router
}
