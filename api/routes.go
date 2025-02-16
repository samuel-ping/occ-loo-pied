package api

import (
	"net/http"

	"github.com/samuel-ping/occ-loo-pied/internal/db"
	"github.com/samuel-ping/occ-loo-pied/internal/ntfy"
)

func SetupRoutes(mux *http.ServeMux, mongoClient *db.MongoClient, notificationClient *ntfy.Client) http.Handler {
	mux.HandleFunc("PUT /api/occupied", func(w http.ResponseWriter, r *http.Request) {
		setOccupiedHandler(w, r, mongoClient, notificationClient)
	})
	mux.HandleFunc("GET /api/occupied", getOccupiedHandler)

	mux.HandleFunc("GET /api/metrics", func(w http.ResponseWriter, r *http.Request) {
		getMetricsHandler(w, r, mongoClient)
	})

	mux.HandleFunc("GET /api/metrics/stats", func(w http.ResponseWriter, r *http.Request) {
		getStatsHandler(w, r, mongoClient)
	})

	mux.HandleFunc("DELETE /api/metrics/{id}", func(w http.ResponseWriter, r *http.Request) {
		deleteMetricHandler(w, r, mongoClient)
	})

	mux.HandleFunc("DELETE /api/metrics/{id}/endTimeAndDuration", func(w http.ResponseWriter, r *http.Request) {
		clearMetricEndTimeAndDurationHandler(w, r, mongoClient)
	})

	mux.HandleFunc("GET /api/metrics/usagesByDay", func(w http.ResponseWriter, r *http.Request) {
		usagesByDayHandler(w, r, mongoClient)
	})

	mux.Handle("GET /", homeHandler())

	return addLogging(addCorsHeaders(mux))
}
