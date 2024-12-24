package api

import (
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("GET /", homeHandler)
	http.HandleFunc("PUT /occupied", setOccupiedHandler)
	http.HandleFunc("GET /occupied", getOccupiedHandler)
}
