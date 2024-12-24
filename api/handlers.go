package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var bathroomOccupied bool

type Bathroom struct {
	Occupied bool `json:"occupied"`
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	var homeText string
	if bathroomOccupied {
		homeText = "Bathroom is currently occupied 🚽"
	} else {
		homeText = "Bathroom is not occupied 😀"
	}
	fmt.Fprintf(w, homeText)
}

func setOccupiedHandler(w http.ResponseWriter, r *http.Request) {
	var occupiedReq Bathroom
	err := json.NewDecoder(r.Body).Decode(&occupiedReq)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	bathroomOccupied = occupiedReq.Occupied

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/json")
}

func getOccupiedHandler(w http.ResponseWriter, _ *http.Request) {
	body, err := json.Marshal(map[string]bool{"occupied": bathroomOccupied})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/json")
	w.Write(body)
}
