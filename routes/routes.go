package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/routes/middlewares"
	"github.com/katesclau/telegramsvc/routes/topic"
	"github.com/katesclau/telegramsvc/routes/topic/event"
	"github.com/katesclau/telegramsvc/routes/topic/subscribers"
	"github.com/katesclau/telegramsvc/routes/topics"
	"github.com/katesclau/telegramsvc/routes/webhook"

	"github.com/gorilla/mux"
)

// Doing this https://stackoverflow.com/questions/35038864/how-to-access-global-variables seems stupid, do I have a better option?
var DB *db.DBClient

type Routes struct {
	routes []Route
	router *mux.Router
	DB     *db.DBClient
}

func (r *Routes) GetRouter() *mux.Router {
	// Authed is not working somehow...
	authedLt := middlewares.ChainMiddleware(middlewares.WithLogging, middlewares.WithTracing, middlewares.WithAuthentication)
	lt := middlewares.ChainMiddleware(middlewares.WithLogging, middlewares.WithTracing)
	r.router = mux.NewRouter()

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
	routes.DB = db
	routes.routes = []Route{
		{
			"/topic/{topicName}/event",
			event.GetMethods(db),
			true,
		},
		{
			"/topic/{topicName}/subscribers",
			subscribers.GetMethods(db),
			true,
		},
		{
			"/topic/{topicName}",
			topic.GetMethods(db),
			true,
		},
		{
			"/topic",
			topics.GetMethods(db),
			true,
		},
		{
			"/webhook",
			webhook.GetMethods(db),
			false,
		},
	}
	return routes
}

type Route struct {
	Path     string
	Methods  map[string]func(w http.ResponseWriter, r *http.Request)
	IsAuthed bool
}

func (r Route) String() string {
	return fmt.Sprintf("{ Path: %s,IsAuthed: %t }", r.Path, r.IsAuthed)
}

func (route *Route) RouteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Route: ", route)
	// Switch Methods
	method := strings.ToUpper(r.Method)
	log.Println("Methods: ", route.Methods)
	if route.Methods[method] != nil {
		for k, v := range r.Header {
			w.Header().Add(k, v[0])
		}
		log.Println("Method: ", method, " - ", route.Methods)
		route.Methods[method](w, r)
		return
	}
	log.Println("Method not supported!", r.Method, r.URL.Path)
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
}
