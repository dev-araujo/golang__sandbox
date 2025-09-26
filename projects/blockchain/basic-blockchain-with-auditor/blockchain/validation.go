package blockchain

import "bytes"

func (bc *Blockchain) Auditor() bool {
	if len(bc.Blocks) == 0 {
		return true
	}

	if !bytes.Equal(bc.Blocks[0].Hash, bc.Blocks[0].calculateHash()) {
		return false
	}

	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if !bytes.Equal(currentBlock.PrevBlockHash, previousBlock.Hash) {
			return false
		}

		if !bytes.Equal(currentBlock.Hash, currentBlock.calculateHash()) {
			return false
		}
	}
	return true
}
