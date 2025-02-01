package db

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	// mongodb keywords
	GROUP          = "$group"
	SUM            = "$sum"
	DATE_TO_STRING = "$dateToString"
	FORMAT         = "format"
	DATE           = "date"

	// fields
	ID_FIELD         = "_id"
	START_TIME_FIELD = "startTime"
	END_TIME_FIELD   = "endTime"
	DURATION_FIELD   = "duration"
	COUNT_FIELD      = "count"
)

type Metric struct {
	Id        bson.ObjectID  `json:"id" bson:"_id"`
	StartTime *time.Time     `json:"startTime" bson:"startTime"`
	EndTime   *time.Time     `json:"endTime" bson:"endTime,omitempty"`
	Duration  *time.Duration `json:"duration,omitempty" bson:"duration"`
}

type UsagesByDayMetric struct {
	Date      string `json:"date" bson:"_id"`
	TimesUsed int    `json:"timesUsed" bson:"count"`
}
