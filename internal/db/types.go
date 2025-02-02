package db

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	// mongodb keywords
	GROUP          = "$group"
	FACET          = "$facet"
	PROJECT        = "$project"
	LIMIT          = "$limit"
	SORT           = "$sort"
	SUM            = "$sum"
	AVG            = "$avg"
	ARRAY_ELEM_AT  = "$arrayElemAt"
	MAX            = "$max"
	DATE_TO_STRING = "$dateToString"
	FORMAT         = "format"
	DATE           = "date"

	// fields
	ID_FIELD                        = "_id"
	START_TIME_FIELD                = "startTime"
	END_TIME_FIELD                  = "endTime"
	DURATION_FIELD                  = "duration"
	COUNT_FIELD                     = "count"
	TOTAL_DURATION_FIELD            = "totalDuration"
	AVERAGE_DURATION_FIELD          = "averageDuration"
	LONGEST_DURATION_AND_DATE_FIELD = "longestDurationAndDate"
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

type Stats struct {
	TotalUsages   int64         `json:"totalUsages"`
	DurationStats DurationStats `json:"duration"`
	// DayMostUsed          *time.Time     `json:"dayMostUsed"`
	// PopularTimeOfDay     *time.Time     `json:"popularTimeOfDay"`
	// PopularDayOfWeek     *time.Weekday  `json:"popularDayOfWeek"`
}

type DurationStats struct {
	Total   int64   `json:"total" bson:"totalDuration"`
	Longest Metric  `json:"longest" bson:"longestDurationAndDate"`
	Average float64 `json:"average" bson:"averageDuration"`
}
