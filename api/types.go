package api

import (
	"time"

	"github.com/samuel-ping/occ-loo-pied/internal/db"
)

type occupiedResponse struct {
	Occupied          bool       `json:"occupied"`
	OccupiedStartTime *time.Time `json:"occupiedStartTime,omitempty"`
}

type setOccupiedRequest struct {
	Occupied bool `json:"occupied"`
}

type getMetricsResponse struct {
	Metrics []db.Metric `json:"metrics"`
}
