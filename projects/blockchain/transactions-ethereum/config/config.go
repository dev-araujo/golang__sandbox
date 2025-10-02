package config

import (
	"os"
)

func LoadConfig() string {

	privateKey := os.Getenv("PRIVATE_KEY")

	if privateKey == "" {
		panic("Variável de ambiente 'PRIVATE_KEY' não encontrada")
	}

	return privateKey

}
