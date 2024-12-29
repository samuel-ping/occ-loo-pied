package api

import (
	"encoding/json"
	"net/http"
	"time"
)

var bathroomOccupied bool
var occupiedStartTime *time.Time

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

func setOccupiedHandler(w http.ResponseWriter, r *http.Request) {
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
