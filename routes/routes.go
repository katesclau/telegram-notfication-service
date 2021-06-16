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
	routes []route.Route
	router *mux.Router
	db     *db.DBClient
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

func NewRoutes(db *db.DBClient) *Routes {
	routes := &Routes{}
	routes.db = db

	// Routes
	routes.routes = []route.Route{
		{
			Path:     "/topic/{topicName}/event",
			Methods:  event.GetMethods(db),
			IsAuthed: true,
		},
		{
			Path:     "/topic/{topicName}/subscribers",
			Methods:  subscribers.GetMethods(db),
			IsAuthed: true,
		},
		{
			Path:     "/topic/{topicName}",
			Methods:  topic.GetMethods(db),
			IsAuthed: true,
		},
		{
			Path:     "/topic/", // TODO Support for multiple paths ('/topic', 'topics', '/topic/')
			Methods:  topics.GetMethods(db),
			IsAuthed: true,
		},
		{
			Path:     "/webhook",
			Methods:  webhook.GetMethods(db),
			IsAuthed: false,
		},
	}

	// Gorilla Mux Router
	router := mux.NewRouter()
	// Use CORS
	router.Use(mux.CORSMethodMiddleware(router))
	routes.router = router

	return routes
}
