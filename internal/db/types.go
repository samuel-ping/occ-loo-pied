package db

import "time"

const (
	START_TIME_FIELD = "startTime"
	END_TIME_FIELD   = "endTime"
)

type Metric struct {
	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
}
