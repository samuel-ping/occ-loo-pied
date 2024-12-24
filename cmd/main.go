package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/samuel-ping/occ-loo-pied/api"
)

func main() {
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "3333" // default port
	}

	mux := http.NewServeMux()
	configuredMux := api.SetupRoutes(mux)

	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, configuredMux); err != nil {
		log.Fatal(err)
	}
}
