package api

import (
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) http.Handler {
	mux.Handle("GET /", homeHandler())

	mux.HandleFunc("PUT /occupied", setOccupiedHandler)
	mux.HandleFunc("GET /occupied", getOccupiedHandler)

	return addLogging(addCorsHeaders(mux))
}
