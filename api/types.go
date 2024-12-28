package api

import "time"

type getOccupiedResponse struct {
	Occupied          bool       `json:"occupied"`
	OccupiedStartTime *time.Time `json:"occupiedStartTime,omitempty"`
}

type setOccupiedRequest struct {
	Occupied bool `json:"occupied"`
}
