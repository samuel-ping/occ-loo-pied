package api

import "time"

type occupiedResponse struct {
	Occupied          bool       `json:"occupied"`
	OccupiedStartTime *time.Time `json:"occupiedStartTime,omitempty"`
}

type setOccupiedRequest struct {
	Occupied bool `json:"occupied"`
}
