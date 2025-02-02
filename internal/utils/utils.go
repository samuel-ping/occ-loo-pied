package utils

import (
	"math"

	"github.com/samuel-ping/occ-loo-pied/internal/db"
)

func IntPtr(v int) *int {
	return &v
}

func FindLeastAndMostUsagesInADay(usagesByDay []db.UsagesByDayMetric) (db.UsagesByDayMetric, db.UsagesByDayMetric) {
	if len(usagesByDay) == 0 {
		return db.UsagesByDayMetric{}, db.UsagesByDayMetric{}
	}

	dayWithLeastUsage := db.UsagesByDayMetric{TimesUsed: math.MaxInt}
	dayWithMostUsage := db.UsagesByDayMetric{}
	for _, usageInDay := range usagesByDay {
		if usageInDay.TimesUsed < dayWithLeastUsage.TimesUsed {
			dayWithLeastUsage = usageInDay
		}
		if usageInDay.TimesUsed > dayWithMostUsage.TimesUsed {
			dayWithMostUsage = usageInDay
		}
	}

	return dayWithLeastUsage, dayWithMostUsage
}
