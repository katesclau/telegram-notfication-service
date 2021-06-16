package middlewares

import (
	"log"
	"net/http"
)

func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("=======Logging Middleware=========")
		log.Printf("Method: %s \n", r.Method)
		log.Printf("Path: %s \n", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
