package blockchain

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data []byte) {
	if len(bc.Blocks) == 0 {
		panic("Não pode adicionar um bloco a uma blockchain vazia")
	}
	prev := bc.Blocks[len(bc.Blocks)-1]

	newBlock := NewBlock(data, prev.Hash, prev.Index+1)

	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewGenesisBlock() *Block {
	data := []byte("Bloco Gênesis")
	PrevBlockHash := []byte{}

	return NewBlock(data, PrevBlockHash, 0)
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewGenesisBlock()
	return &Blockchain{Blocks: []*Block{genesisBlock}}
}
