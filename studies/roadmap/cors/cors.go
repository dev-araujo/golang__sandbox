package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors" // Importa a biblioteca
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "Esta rota está protegida por CORS!"}`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	// Configuração do CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "https://meu-site.com"}, // Suas origens permitidas
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true, // Permite que o navegador envie cookies e headers de autenticação
		Debug:            true, // Habilita logs para depuração
	})

	// Envolve o seu router (mux) com o handler de CORS
	handler := c.Handler(mux)

	fmt.Println("Servidor rodando na porta 8080 com CORS habilitado...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
