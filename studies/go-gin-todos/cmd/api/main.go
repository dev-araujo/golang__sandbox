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
		log.Println("Aviso: Não foi possível encontrar o ficheiro .env")
	}

	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal("Não foi possível conectar à base de dados: ", err)
	}

	router := handlers.SetupRoutes(store)

	log.Println("Servidor a arrancar na porta 8080...")
	router.Run(":8080")
}
