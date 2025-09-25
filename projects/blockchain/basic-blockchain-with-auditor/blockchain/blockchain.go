package blockchain

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data []byte) {
    if len(bc.Blocks) == 0 {
        panic("cannot add block to an empty chain")
    }
	prev := bc.Blocks[len(bc.Blocks)-1]

	newBlock := NewBlock(data, prev.Hash)
	newBlock.Index = prev.Index + 1

	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewGenesisBlock() *Block {
	data := []byte("Genesis Block")
	PrevBlockHash := []byte{}

	return NewBlock(data, PrevBlockHash)
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewGenesisBlock()
	return &Blockchain{Blocks: []*Block{genesisBlock}}
}
