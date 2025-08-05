package api

import (
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	server := &Server{
		router: http.NewServeMux(),
	}
	server.registerRoutes()
	return server
}

func (s *Server) registerRoutes() {
	s.router.HandleFunc("/to-do", TodoHandler)
}

func (s *Server) Router() http.Handler {
	return CorsMiddleware(s.router)
}
