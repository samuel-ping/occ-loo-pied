package utils

import (
	"github.com/samuel-ping/occ-loo-pied/internal/db"
)

func IntPtr(v int) *int {
	return &v
}

func FindMostUsagesInADay(usagesByDay []db.UsagesByDayMetric) db.UsagesByDayMetric {
	if len(usagesByDay) == 0 {
		return db.UsagesByDayMetric{}
	}

	dayWithMostUsage := db.UsagesByDayMetric{}
	for _, usageInDay := range usagesByDay {
		if usageInDay.TimesUsed > dayWithMostUsage.TimesUsed {
			dayWithMostUsage = usageInDay
		}
	}

	return dayWithMostUsage
}
