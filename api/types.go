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
	Metrics    []db.Metric `json:"metrics"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	TotalItems int  `json:"totalItems"`
	Page       int  `json:"page"`
	TotalPages int  `json:"totalPages"`
	NextPage   *int `json:"nextPage"`
	PrevPage   *int `json:"prevPage"`
}
