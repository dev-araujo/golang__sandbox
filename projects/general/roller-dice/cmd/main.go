package main

import (
	"log"
	"net/http"

	"github.com/dev-araujo/golang__sandbox/roller-dice/internal/api"
)

func main() {

	server := api.NewServer()

	log.Println("Server on http://localhost:8080")

	if err := http.ListenAndServe(":8080", server.Router()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}