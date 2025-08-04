package main

import (
	"fmt"
	"log"
	"net/http"
)

// nosso handler (manipulador da rota)
func headersHandler(w http.ResponseWriter, r *http.Request) {
	// 1. LENDO UM HEADER DA REQUISIÇÃO
	// O objeto 'r' (http.Request) contém tudo sobre a requisição que chegou.
	// r.Header é um mapa de todos os headers. Usamos o método Get() para pegar um específico.
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		log.Println("Header 'Authorization' não encontrado.")
	} else {
		log.Printf("Header 'Authorization' recebido: %s\n", authToken)
	}

	// 2. DEFININDO UM HEADER NA RESPOSTA
	// O objeto 'w' (http.ResponseWriter) é usado para construir a resposta.
	// w.Header() nos dá acesso ao mapa de headers da resposta. Usamos Set() para definir um.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Golang") // Um header customizado

	// Antes de escrever o corpo, é importante definir o status code.
	w.WriteHeader(http.StatusOK) // Status 200 OK

	// Escrevendo o corpo da resposta (um JSON simples)
	fmt.Fprintln(w, `{"message": "Olá do servidor Go! Confira os headers da resposta no seu navegador."}`)
}

func main() {
	http.HandleFunc("/headers", headersHandler)

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
