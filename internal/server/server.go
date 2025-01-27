package server

import (
	"net/http"

	"github.com/samuel-ping/occ-loo-pied/api"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	Router *http.ServeMux
	Client *mongo.Client
}

func NewServer(client *mongo.Client) *Server {
	server := &Server{
		Router: http.NewServeMux(),
		Client: client,
	}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	api.SetupRoutes(s.Router, s.Client)
}
