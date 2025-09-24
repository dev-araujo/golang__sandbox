package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PublicKey  ecdsa.PublicKey
	PrivateKey ecdsa.PrivateKey
}

func NewWallet() *Wallet {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey

	return &Wallet{PublicKey: publicKey, PrivateKey: *privateKey}

}

func (w *Wallet) getAddress() string {
	x := w.PublicKey.X.Bytes()
	y := w.PublicKey.Y.Bytes()

	pubKeyBytes := append(x, y...)[1:]

	hash := crypto.Keccak256(pubKeyBytes)[12:]

	return hex.EncodeToString(hash)

}
