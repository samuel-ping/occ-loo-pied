package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/samuel-ping/occ-loo-pied/internal/db"
	"github.com/samuel-ping/occ-loo-pied/web"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var bathroomOccupied bool
var occupiedStartTime *time.Time

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// source: https://www.liip.ch/en/blog/embed-sveltekit-into-a-go-binary
		fs := http.FS(web.GetSvelteFs())

		path := strings.TrimPrefix(r.URL.Path, "/")
		_, err := fs.Open(path)
		if errors.Is(err, os.ErrNotExist) {
			path = fmt.Sprintf("%s.html", path)
		}
		r.URL.Path = path

		http.FileServer(fs).ServeHTTP(w, r)
	})
}

func getOccupiedHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		occupiedResponse{
			Occupied:          bathroomOccupied,
			OccupiedStartTime: occupiedStartTime,
		},
	)
}

func setOccupiedHandler(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	var req setOccupiedRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	bathroomOccupied = req.Occupied
	if bathroomOccupied {
		startTime := time.Now()
		occupiedStartTime = &startTime
	} else {
		endTime := time.Now()
		db.AddOccupiedMetric(client, occupiedStartTime, &endTime)

		occupiedStartTime = nil
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		occupiedResponse{
			Occupied:          bathroomOccupied,
			OccupiedStartTime: occupiedStartTime,
		},
	)
}

func getMetricsHandler(w http.ResponseWriter, _ *http.Request, client *mongo.Client) {
	metrics, err := db.GetAllMetrics(client)
	if err != nil {
		log.Printf("Error getting metrics from db: %v", err)
		http.Error(w, "Error getting metrics", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		getMetricsResponse{Metrics: metrics},
	)
}
