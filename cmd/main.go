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

	api.SetupRoutes()

	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
