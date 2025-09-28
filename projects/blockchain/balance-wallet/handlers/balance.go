package handlers

import (
	"context"
	"net/http"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

const RPC_URL_TESTNET = "https://ethereum-sepolia-rpc.publicnode.com"
const RPC_MAINNET = "https://ethereum-rpc.publicnode.com"
const WEI_ETHER = 1e18 // 10^18

func GetStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

var validNetworks = map[string]bool{
	"mainnet": true,
	"testnet": true,
}

func GetBalance(c *gin.Context) {
	network := c.Param("network")

	if !validNetworks[network] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "network inválida. Use 'mainnet' ou 'testnet'",
		})
		return
	}

	rpc := RPC_MAINNET

	if network == "testnet" {
		rpc = RPC_URL_TESTNET
	}

	connection, err := ethclient.Dial(rpc)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	address := c.Param("address")

	balance, err := connection.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	diviser := big.NewInt(WEI_ETHER)

	balanceConversion := big.NewInt(WEI_ETHER) // 10^18
	balanceConversion = balanceConversion.Div(balance, diviser)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"balance":           balance,
		"balanceConversion": balanceConversion,
	})

}
