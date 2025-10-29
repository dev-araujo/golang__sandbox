// No arquivo: main.go

package main

import (
	"github.com/dev-araujo/golang__sandbox/projects/blockchain/transactions-ethereum/config"
	"github.com/dev-araujo/golang__sandbox/projects/blockchain/transactions-ethereum/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	privateKey := config.LoadConfig()

	h := handlers.NewHandler(privateKey)

	router := gin.Default()

	router.GET("/ping", handlers.GetStatus)

	router.POST("/send", h.SendTransactionHandler)

	router.Run(":8080")
}
