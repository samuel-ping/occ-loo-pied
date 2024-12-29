package api

import (
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) http.Handler {
	fs := http.FileServer(http.Dir("./web/build"))
	mux.Handle("GET /", fs)

	mux.HandleFunc("PUT /occupied", setOccupiedHandler)
	mux.HandleFunc("GET /occupied", getOccupiedHandler)

	return addLogging(addCorsHeaders(mux))
}
