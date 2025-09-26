package blockchain

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	t.Run("Deve retornar true para uma blockchain válida", func(t *testing.T) {
		bc := NewBlockchain()
		bc.AddBlock([]byte("block 1"))

		if !bc.Auditor() {
			t.Errorf("Esperava uma blockchain válida, mas ela não é válida")
		}
	})

	t.Run("Deve retornar false para uma blockchain com hashes inválidos", func(t *testing.T) {
		bc := NewBlockchain()
		bc.AddBlock([]byte("block 1"))
		bc.Blocks[1].PrevBlockHash = []byte("invalid hash")

		if bc.Auditor() {
			t.Errorf("Esperava uma blockchain com hashes inválidos, mas ela é válida")
		}
	})

	t.Run("Deve retornar false para uma blockchain com dados corrompidos", func(t *testing.T) {
		bc := NewBlockchain()
		bc.AddBlock([]byte("block 1"))
		bc.Blocks[1].Data = []byte("invalid data")

		if bc.Auditor() {
			t.Errorf("Esperava uma blockchain com dados corrompidos, mas ela é válida")
		}
	})

	t.Run("Deve retornar false quando o bloco genesis é corrompido", func(t *testing.T) {
		bc := NewBlockchain()
		bc.AddBlock([]byte("block 1"))
		bc.Blocks[0].Data = []byte("genesis corrompido")

		if bc.Auditor() {
			t.Errorf("Esperava que blockchain com genesis corrompido fosse inválida")
		}
	})
}
