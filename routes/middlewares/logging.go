package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logging Middleware")
		log.Println(fmt.Printf("Method: %s, ", r.Method))
		log.Println(fmt.Printf("Path: %s, ", r.URL.Path))
		next.ServeHTTP(w, r)
		log.Println("Logging Afterware")
	})
}
