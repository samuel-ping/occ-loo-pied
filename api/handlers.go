package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/samuel-ping/occ-loo-pied/internal/db"
	"github.com/samuel-ping/occ-loo-pied/web"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var bathroomOccupied bool
var occupiedStartTime *time.Time

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(web.GetSvelteFs())).ServeHTTP(w, r)
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
