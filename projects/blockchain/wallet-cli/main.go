package main

import (
	"fmt"

	"github.com/dev-araujo/golang__sandbox/projects/blockchain/wallet-cli/wallet"
)

func main() {

	wallet := wallet.NewWallet()

	fmt.Printf("Chave privada %v \n", wallet.PrivateKey)
	fmt.Println("====================================")
	fmt.Printf("Chave publica %v \n", wallet.PublicKey)
	fmt.Println("====================================")

	fmt.Printf("Endereço %v", wallet.GetAddress())

}
