package main

import (
	"github.com/dev-araujo/golang__sandbox/projects/blockchain/transactions-ethereum/config"
	h "github.com/dev-araujo/golang__sandbox/projects/blockchain/transactions-ethereum/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	env := config.LoadConfig()

	router := gin.Default()
	router.GET("/ping", h.GetStatus)

	router.Run(":8080")
}
