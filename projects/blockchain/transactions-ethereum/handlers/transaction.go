package handlers

import (
	"math/big"

	"github.com/gin-gonic/gin"
)

type TransactionRequest struct {
	To    string
	Value string
}

func GetStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SendTransactionHandler(c *gin.Context) {
	request := TransactionRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	bigInt := new(big.Int)
	bigInt, ok := bigInt.SetString(request.Value, 10)

	if !ok {
		c.JSON(400, gin.H{"Bad Request": "Invalid value"})
		return
	}

}
