package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/profclems/go-dotenv"
)

var token = dotenv.GetString("API_TOKEN")

// TODO: https://hackernoon.com/creating-a-middleware-in-golang-for-jwt-based-authentication-cx3f32z8 when Client Model is implemented
func WithAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authentication check...")
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		}
		log.Printf("Authentication header: %s", authHeader)
		if authHeader[1] != token {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Invalid Authentication Header"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
