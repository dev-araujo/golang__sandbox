package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const RPC_URL_TESTNET = "https://ethereum-sepolia-rpc.publicnode.com"

func CreateAndSendTransaction(pk string, receiverAddress string, value *big.Int) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	connection, err := ethclient.Dial(RPC_URL_TESTNET)
	if err != nil {
		panic("Erro de conexão com a rede Ethereum")
	}

	nonce, err := connection.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		panic(err)
	}

	gasPrice, err := connection.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(receiverAddress), value, 21000, gasPrice, nil)

	chainId, err := connection.ChainID(context.Background())

	if err != nil {
		panic(err)
	}

	signer := types.NewEIP155Signer(chainId)
	signedTx, err := types.SignTx(tx, signer, privateKey)

	if err != nil {
		panic(err)
	}
	err = connection.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}

}
