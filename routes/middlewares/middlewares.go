package middlewares

import (
	"log"
	"net/http"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

// ChainMiddleware provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
func ChainMiddleware(mw ...Middleware) Middleware {
	log.Printf("ChainMiddleware:mw %d", len(mw))
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}
