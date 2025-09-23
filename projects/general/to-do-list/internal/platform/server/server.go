package server

import (
	"net/http"

	"github.com/dev-araujo/golang__sandbox/to-do-list/pkg/middleware"
	"github.com/dev-araujo/golang__sandbox/to-do-list/pkg/task"
)

type Server struct {
	router *http.ServeMux
}

func NewServer(taskService task.Service) *Server {
	server := &Server{
		router: http.NewServeMux(),
	}

	taskRouter := task.Routes(taskService)
	server.router.Handle("/", taskRouter)

	return server
}

func (s *Server) Router() http.Handler {
	return middleware.CorsMiddleware(s.router)
}
