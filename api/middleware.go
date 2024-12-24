package api

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Handling %s request to %s\n", r.Method, r.URL)
			next.ServeHTTP(w, r)
		},
	)
}
