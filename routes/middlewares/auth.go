package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/profclems/go-dotenv"
)

var token = dotenv.GetString("API_TOKEN")

// TODO: https://hackernoon.com/creating-a-middleware-in-golang-for-jwt-based-authentication-cx3f32z8 when Client Model is implemented
func WithAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		if len(authHeader) != 2 {
			message := "Unauthorized: Malformed Token!"
			log.Println(message)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(message))
			return
		}

		// Temporary... later, we can have a single Client endpoint to register JWT and manage tokens by Client identity
		if authHeader[1] != token {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Unauthorized: Invalid Authentication Header"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
