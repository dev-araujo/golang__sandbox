// Ficheiro: cmd/api/main.go

package main

import (
	"log"
	"todo-api/internal/handlers"
	"todo-api/internal/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o ficheiro .env")
	}

	db := storage.InitDB()
	defer db.Close()

	router := handlers.SetupRoutes(db)

	// 3. Arranca o servidor
	log.Println("Servidor a arrancar na porta 8080...")
	router.Run(":8080")
}
