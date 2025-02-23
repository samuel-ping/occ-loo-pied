package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/samuel-ping/occ-loo-pied/internal/db"
	"github.com/samuel-ping/occ-loo-pied/internal/ntfy"
	"github.com/samuel-ping/occ-loo-pied/internal/utils"
	"github.com/samuel-ping/occ-loo-pied/web"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var bathroomOccupied bool
var occupiedStartTime *time.Time

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// source: https://www.liip.ch/en/blog/embed-sveltekit-into-a-go-binary
		fs := http.FS(web.GetSvelteFs())

		path := strings.TrimPrefix(r.URL.Path, "/")
		_, err := fs.Open(path)
		if errors.Is(err, os.ErrNotExist) {
			path = fmt.Sprintf("%s.html", path)
		}
		r.URL.Path = path

		http.FileServer(fs).ServeHTTP(w, r)
	})
}

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

func setOccupiedHandler(w http.ResponseWriter, r *http.Request, mongoClient *db.MongoClient, ntfyClient *ntfy.Client) {
	var req setOccupiedRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding request: %v\n", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	bathroomOccupied = req.Occupied
	if bathroomOccupied {
		startTime := time.Now()
		occupiedStartTime = &startTime
	} else {
		endTime := time.Now()
		mongoClient.AddOccupiedMetric(occupiedStartTime, &endTime)
		occupiedStartTime = nil
	}

	// send notification
	ntfyClient.SendOccupationNotification(bathroomOccupied)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		occupiedResponse{
			Occupied:          bathroomOccupied,
			OccupiedStartTime: occupiedStartTime,
		},
	)
}

func getMetricsHandler(w http.ResponseWriter, r *http.Request, client *db.MongoClient) {
	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		log.Printf("error parsing page number: %v\n", err)
		http.Error(w, "error parsing page number", http.StatusBadRequest)
		return
	}
	itemsPerPageParam := r.URL.Query().Get("itemsPerPage")
	itemsPerPage, err := strconv.Atoi(itemsPerPageParam)
	if err != nil {
		log.Printf("error parsing items per page number: %v\n", err)
		http.Error(w, "error parsing page number", http.StatusBadRequest)
		return
	}

	skip := (page - 1) * itemsPerPage

	totalDocuments, err := client.DocumentCount()
	if err != nil {
		log.Printf("error getting document count: %v\n", err)
		http.Error(w, "Error getting document count", http.StatusInternalServerError)
		return
	}

	totalPages := int(math.Ceil(float64(totalDocuments) / float64(itemsPerPage)))
	if page > totalPages {
		log.Println("requested page number exceeds page count")
		http.Error(w, "Requested page number exceeds page count", http.StatusBadRequest)
		return
	}

	metrics, err := client.GetMetrics(int64(skip), int64(itemsPerPage))
	if err != nil {
		log.Printf("Error getting metrics from db: %v\n", err)
		http.Error(w, "Error getting metrics", http.StatusInternalServerError)
		return
	}

	var nextPage *int
	if page+1 <= totalPages {
		nextPage = utils.IntPtr(page + 1)
	}

	var prevPage *int
	if page-1 >= 1 {
		prevPage = utils.IntPtr(page - 1)
	}

	w.Header().Set("Content-Type", "text/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		getMetricsResponse{
			Metrics: metrics,
			Pagination: Pagination{
				TotalItems: totalDocuments,
				Page:       page,
				TotalPages: totalPages,
				NextPage:   nextPage,
				PrevPage:   prevPage,
			},
		},
	)
}

func clearMetricEndTimeAndDurationHandler(w http.ResponseWriter, r *http.Request, client *db.MongoClient) {
	rawMetricId := r.PathValue("id")
	if len(rawMetricId) == 0 {
		log.Println("No id in path param")
		http.Error(w, "No id in path param", http.StatusBadRequest)
		return
	}

	metricId, err := bson.ObjectIDFromHex(rawMetricId)
	if err != nil {
		log.Printf("error converting id %s to objectId: %v\n", rawMetricId, err)
		http.Error(w, "error converting id to objectId", http.StatusBadRequest)
		return
	}

	if err := client.ClearMetricEndTimeAndDuration(metricId); err != nil {
		log.Printf("Error clearing endTime for metric %s: %v\n", metricId, err)
		http.Error(w, "Error clearing endTime for metric", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteMetricHandler(w http.ResponseWriter, r *http.Request, client *db.MongoClient) {
	idToDelete := r.PathValue("id")
	if len(idToDelete) == 0 {
		log.Println("No id in path param")
		http.Error(w, "No id in path param", http.StatusBadRequest)
		return
	}

	_, err := client.DeleteMetric(idToDelete)
	if err != nil {
		log.Printf("Error deleting metric %s: %v\n", idToDelete, err)
		http.Error(w, "Error deleting metric", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func usagesByDayHandler(w http.ResponseWriter, _ *http.Request, client *db.MongoClient) {
	usagesByDay, err := client.UsagesByDay()
	if err != nil {
		log.Printf("Error getting usages by day: %v\n", err)
		http.Error(w, "Error getting usages by day", http.StatusInternalServerError)
		return
	}

	dayWithLeastUsage, dayWithMostUsage := utils.FindLeastAndMostUsagesInADay(usagesByDay)

	w.Header().Set("Content-Type", "text/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		usagesByDayResponse{
			UsagesByDay:       usagesByDay,
			LeastUsagesInADay: dayWithLeastUsage.TimesUsed,
			MostUsagesInADay:  dayWithMostUsage.TimesUsed,
		},
	)
}

func getStatsHandler(w http.ResponseWriter, _ *http.Request, client *db.MongoClient) {
	generalMetrics, err := client.CalcStats()
	if err != nil {
		log.Printf("error calculating stats: %v\n", err)
		http.Error(w, "error calculating stats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		statsResponse{
			Stats: generalMetrics,
		},
	)
}
