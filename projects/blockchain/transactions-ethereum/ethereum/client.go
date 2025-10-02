package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

const RPC_URL_TESTNET = "https://ethereum-sepolia-rpc.publicnode.com"

func CreateAndSendTransaction(pk string, address string, value *big.Int) {

	connection, err := ethclient.Dial(RPC_URL_TESTNET)
	if err != nil {
		panic("Erro de conexão com a rede Ethereum")
	}

}
