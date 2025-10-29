package handlers

import (
	"math/big"
	"net/http"

	"github.com/dev-araujo/golang__sandbox/projects/blockchain/transactions-ethereum/ethereum"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	privateKey string
}

func NewHandler(pk string) *Handler {
	return &Handler{privateKey: pk}
}

type TransactionRequest struct {
	To    string
	Value string
}

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *Handler) SendTransactionHandler(c *gin.Context) {

	request := TransactionRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	bigInt, ok := new(big.Int).SetString(request.Value, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor inválido, deve ser um número"})
		return
	}

	err := ethereum.CreateAndSendTransaction(h.privateKey, request.To, bigInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno ao processar a transação"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Transação enviada com sucesso"})
}
