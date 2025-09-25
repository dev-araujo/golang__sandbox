package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Index         uint64
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) calculateHash() []byte {
	timestampBytes := strconv.FormatInt(b.Timestamp, 10)
	indexBytes := strconv.FormatUint(b.Index, 10)

	info := [][]byte{[]byte(timestampBytes), []byte(indexBytes), b.Data, b.PrevBlockHash}
	joinedInfo := bytes.Join(info, []byte{})

	hash := sha256.Sum256(joinedInfo)
	return hash[:]
}

func NewBlock(data []byte, prevBlockHash []byte) *Block {

	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
		Timestamp:     time.Now().Unix(),
	}

	block.Hash = block.calculateHash()

	return block

}
