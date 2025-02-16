package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/samuel-ping/occ-loo-pied/internal/db"
	"github.com/samuel-ping/occ-loo-pied/internal/ntfy"
	"github.com/samuel-ping/occ-loo-pied/internal/server"
)

const PORT = "PORT"
const NTFY_BASE_URL = "NTFY_BASE_URL"
const TOPIC = "NTFY_TOPIC"
const TOKEN = "NTFY_TOKEN"
const DEFAULT_TOPIC = "bathroom_test"

func main() {
	port, found := os.LookupEnv(PORT)
	if !found {
		port = "3333" // default port
	}

	mongoClient, err := db.NewMongoClient()
	if err != nil {
		log.Fatalf("Error creating Mongo client: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	ntfyBaseUrl, found := os.LookupEnv(NTFY_BASE_URL)
	if !found {
		log.Fatalf("Must set %s environmental variable", NTFY_BASE_URL)
	}

	topic, found := os.LookupEnv(TOPIC)
	if !found {
		log.Printf("%s environmental variable not found; defaulting to %s\n", TOPIC, DEFAULT_TOPIC)
		topic = DEFAULT_TOPIC
	}

	token, found := os.LookupEnv(TOKEN)
	if !found {
		log.Printf("%s environmental variable not found; assuming no auth\n", TOKEN)
	}

	notificationClient := ntfy.NewClient(ntfyBaseUrl, topic, token)

	srv := server.NewServer(mongoClient, notificationClient)

	log.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, srv.Router); err != nil {
		log.Fatal(err)
	}
}
