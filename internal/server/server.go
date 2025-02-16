package server

import (
	"net/http"

	"github.com/samuel-ping/occ-loo-pied/api"
	"github.com/samuel-ping/occ-loo-pied/internal/db"
	"github.com/samuel-ping/occ-loo-pied/internal/ntfy"
)

type Server struct {
	Router             *http.ServeMux
	MongoClient        *db.MongoClient
	NotificationClient *ntfy.Client
}

func NewServer(mongoClient *db.MongoClient, notificationClient *ntfy.Client) *Server {
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
