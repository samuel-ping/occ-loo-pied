package api

import (
	"net/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupRoutes(mux *http.ServeMux, client *mongo.Client) http.Handler {
	mux.HandleFunc("PUT /api/occupied", func(w http.ResponseWriter, r *http.Request) {
		setOccupiedHandler(w, r, client)
	})
	mux.HandleFunc("GET /api/occupied", getOccupiedHandler)

	mux.HandleFunc("GET /api/metrics", func(w http.ResponseWriter, r *http.Request) {
		getMetricsHandler(w, r, client)
	})

	mux.HandleFunc("GET /api/metrics/stats", func(w http.ResponseWriter, r *http.Request) {
		getStatsHandler(w, r, client)
	})

	mux.HandleFunc("DELETE /api/metrics/{id}", func(w http.ResponseWriter, r *http.Request) {
		deleteMetricHandler(w, r, client)
	})

	mux.HandleFunc("DELETE /api/metrics/{id}/endTimeAndDuration", func(w http.ResponseWriter, r *http.Request) {
		clearMetricEndTimeAndDurationHandler(w, r, client)
	})

	mux.HandleFunc("GET /api/metrics/usagesByDay", func(w http.ResponseWriter, r *http.Request) {
		usagesByDayHandler(w, r, client)
	})

	mux.Handle("GET /", homeHandler())

	return addLogging(addCorsHeaders(mux))
}
