package api

import (
	"net/http"

	"github.com/dev-araujo/golang__sandbox/to-do-list/internal/task"
)

type Server struct {
	router *http.ServeMux
	task   task.Service
}

func NewServer(task task.Service) *Server {
	server := &Server{
		router: http.NewServeMux(),
		task:   task,
	}
	server.registerRoutes()
	return server
}

func (s *Server) registerRoutes() {
	s.router.HandleFunc("/todo/list", s.getListTasksHandler)
	s.router.HandleFunc("/todo/add", s.addTaskHandler)
	s.router.HandleFunc("/todo/delete", s.deleteTaskHandler)
	s.router.HandleFunc("/todo/update", s.updateTaskHandler)
	s.router.HandleFunc("/todo/get", s.getTaskHandler)
}

func (s *Server) Router() http.Handler {
	return CorsMiddleware(s.router)
}
