package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dev-araujo/golang__sandbox/to-do-list/internal/platform/server"
	"github.com/dev-araujo/golang__sandbox/to-do-list/pkg/task"
)

func main() {
	taskService := task.NewService()
	server := server.NewServer(taskService)

	// Configura o handler com CORS
	http.Handle("/", server.Router())

	// Inicia o servidor na porta 8080
	fmt.Println("Servidor rodando na porta 8080 com CORS habilitado...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}