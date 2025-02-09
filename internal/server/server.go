package server

import (
	"net/http"

	"github.com/samuel-ping/occ-loo-pied/api"
	"github.com/samuel-ping/occ-loo-pied/internal/ntfy"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	Router             *http.ServeMux
	MongoClient        *mongo.Client
	NotificationClient *ntfy.Client
}

func NewServer(mongoClient *mongo.Client, notificationClient *ntfy.Client) *Server {
	server := &Server{
		Router:             http.NewServeMux(),
		MongoClient:        mongoClient,
		NotificationClient: notificationClient,
	}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	api.SetupRoutes(s.Router, s.MongoClient, s.NotificationClient)
}
