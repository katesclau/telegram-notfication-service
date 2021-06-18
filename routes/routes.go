package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/routes/middlewares"
	"github.com/katesclau/telegramsvc/routes/route"
	"github.com/katesclau/telegramsvc/routes/topic"
	"github.com/katesclau/telegramsvc/routes/topic/event"
	"github.com/katesclau/telegramsvc/routes/topic/subscribers"
	"github.com/katesclau/telegramsvc/routes/topics"
	"github.com/katesclau/telegramsvc/routes/webhook"
)

var DB *db.DBClient

type Routes struct {
	routes  []route.Route
	router  *mux.Router
	context *route.Context
}

func (r *Routes) GetRouter() *mux.Router {
	// Authed is not working somehow...
	authedLt := middlewares.ChainMiddleware(middlewares.WithLogging, middlewares.WithTracing, middlewares.WithAuthentication)
	lt := middlewares.ChainMiddleware(middlewares.WithLogging, middlewares.WithTracing)

	// Static content
	fs := http.FileServer(http.Dir("./routes/static"))
	r.router.Handle("/", fs)

	for i := 0; i < len(r.routes); i++ {
		route := r.routes[i]
		log.Println("Route: ", route)
		if route.IsAuthed {
			r.router.HandleFunc(route.Path, authedLt(route.RouteHandler))
		} else {
			r.router.HandleFunc(route.Path, lt(route.RouteHandler))
		}
	}
	return r.router
}

func NewRoutes(ctx *route.Context) *Routes {
	routes := &Routes{}
	routes.context = ctx

	// Routes
	routes.routes = []route.Route{
		event.GetRoute(ctx),
		subscribers.GetRoute(ctx),
		topic.GetRoute(ctx),
		topics.GetRoute(ctx),
		webhook.GetRoute(ctx),
	}

	// Gorilla Mux Router
	router := mux.NewRouter()
	// Use CORS
	router.Use(mux.CORSMethodMiddleware(router))
	routes.router = router

	return routes
}
