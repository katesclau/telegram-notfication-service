package route

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Route struct {
	Path     string
	Methods  map[string]func(w http.ResponseWriter, r *http.Request)
	IsAuthed bool
}

func (r Route) String() string {
	return fmt.Sprintf("{ Path: %s,IsAuthed: %t, Methods: %v }\n", r.Path, r.IsAuthed, r.Methods)
}

func (route *Route) RouteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Route: ", route)
	// Switch Methods
	method := strings.ToUpper(r.Method)
	log.Println("Methods: ", route.Methods)
	if route.Methods[method] != nil {
		for k, v := range r.Header {
			if strings.ToLower(k) == "content-length" {
				continue
			}
			log.Println("header:value", k, v[0])
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
