package api

import (
	"net/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupRoutes(mux *http.ServeMux, client *mongo.Client) http.Handler {
	mux.Handle("GET /", homeHandler())

	mux.HandleFunc("PUT /api/occupied", func(w http.ResponseWriter, r *http.Request) {
		setOccupiedHandler(w, r, client)
	})
	mux.HandleFunc("GET /api/occupied", getOccupiedHandler)

	mux.HandleFunc("GET /api/metrics", func(w http.ResponseWriter, r *http.Request) {
		getMetricsHandler(w, r, client)
	})

	return addLogging(addCorsHeaders(mux))
}
