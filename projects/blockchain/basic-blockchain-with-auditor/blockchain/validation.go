package blockchain

import "bytes"

func (bc *Blockchain) IsValid() bool {

	for i := 1; i < len(bc.Blocks); i++ {
		firstHashValidation := bytes.Equal(bc.Blocks[i].PrevBlockHash, bc.Blocks[i-1].Hash)
		secondHashValidation := bytes.Equal(bc.Blocks[i].Hash, bc.Blocks[i].calculateHash())

		if !firstHashValidation || !secondHashValidation {
			return false
		}

	}
	return true

}
