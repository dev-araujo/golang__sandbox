package ethereum

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const RPC_URL_TESTNET = "https://ethereum-sepolia-rpc.publicnode.com"

func CreateAndSendTransaction(pk string, receiverAddress string, value *big.Int) error {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		fmt.Println(err)

		return err
	}

	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	connection, err := ethclient.Dial(RPC_URL_TESTNET)
	if err != nil {
		fmt.Println("Erro de conexão com a rede Ethereum")
		return err
	}

	nonce, err := connection.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		fmt.Println(err)
		return err
	}

	gasPrice, err := connection.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(receiverAddress), value, 21000, gasPrice, nil)

	chainId, err := connection.ChainID(context.Background())

	if err != nil {
		fmt.Println(err)
		return err
	}

	signer := types.NewEIP155Signer(chainId)
	signedTx, err := types.SignTx(tx, signer, privateKey)

	if err != nil {
		fmt.Println(err)
		return err
	}
	err = connection.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
