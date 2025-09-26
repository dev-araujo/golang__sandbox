package main

import (
	h "github.com/dev-araujo/golang__sandbox/projects/basic-blockchain-go/balance-wallet/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", h.GetStatus)

	router.Run(":8080")
}
