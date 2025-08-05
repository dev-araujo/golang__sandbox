package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dev-araujo/golang__sandbox/to-do-list/internal/api"
)

func main() { // Inicializa o servidor
	server := api.NewServer()

	// Configura o handler com CORS
	http.Handle("/", server.Router())

	// Inicia o servidor na porta 8080
	fmt.Println("Servidor rodando na porta 8080 com CORS habilitado...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
