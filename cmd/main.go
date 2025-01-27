package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/samuel-ping/occ-loo-pied/internal/db"
	"github.com/samuel-ping/occ-loo-pied/internal/server"
)

func main() {
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "3333" // default port
	}

	mongoClient, err := db.ConnectMongo()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	srv := server.NewServer(mongoClient)

	log.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, srv.Router); err != nil {
		log.Fatal(err)
	}
}
