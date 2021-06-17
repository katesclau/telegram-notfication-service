package middlewares

import (
	"log"
	"net/http"

	"github.com/katesclau/telegramsvc/utils"
)

func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s \n", r.Method)
		log.Printf("Path: %s \n", r.URL.Path)
		log.Printf("Body: %q \n", utils.DecodeBody(r))
		next.ServeHTTP(w, r)
	})
}
