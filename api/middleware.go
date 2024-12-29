package api

import (
	"log"
	"net/http"
	"strings"
)

func addLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Handling %s request to %s\n", r.Method, r.URL)
			next.ServeHTTP(w, r)
		},
	)
}

func addCorsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Check the origin is valid.
			origin := r.Header.Get("Origin")
			validOrigin := validateOrigin(origin)

			// If it is, allow CORS.
			if validOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET,PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
			}

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}

func validateOrigin(origin string) bool {
	// Client running locally
	if strings.Index(origin, "http://localhost:") == 0 {
		return true
	}

	// Client running in prod
	return origin == "http://192.168.0.12:3333"
}
