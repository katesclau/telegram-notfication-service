package route

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/katesclau/telegramsvc/routes/method_not_allowed"
)

type Route struct {
	Path     string
	Methods  map[string]*MethodHandler
	IsAuthed bool
	context  *Context
}

func (r Route) String() string {
	return fmt.Sprintf("{ Path: %s,IsAuthed: %t, Methods: %v }\n", r.Path, r.IsAuthed, r.Methods)
}

func (route *Route) SetMethodHandlers(routeHandlerMap map[string]func(ctx *Context, w http.ResponseWriter, r *http.Request)) {
	methods := make(map[string]*MethodHandler)
	for k, v := range routeHandlerMap {
		methods[k] = NewMethodHandler(route.context, k, v)
	}
	route.Methods = methods
}

func (route *Route) RouteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Route: ", route)
	// Switch Methods
	method := strings.ToUpper(r.Method)
	log.Println("Methods: ", route.Methods)
	if route.Methods[method] != nil {
		methodHandler := route.Methods[method]
		log.Println("MethodHandler: ")
		for k, v := range r.Header {
			if strings.ToLower(k) == "content-length" {
				continue
			}
			log.Println("header:value", k, v[0])
			w.Header().Add(k, v[0])
		}
		methodHandler.Execute(w, r)
		return
	}
	method_not_allowed.Handler(w, r)
}

func NewRoute(ctx *Context, path string, isAuthed bool) *Route {
	return &Route{
		Path:     path,
		context:  ctx,
		IsAuthed: isAuthed,
	}
}
