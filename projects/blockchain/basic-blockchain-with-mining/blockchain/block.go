package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index         uint64
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Difficulty    int
	Nonce         int
}

func (b *Block) calculateHash() []byte {
	timestampBytes := strconv.FormatInt(b.Timestamp, 10)
	indexBytes := strconv.FormatUint(b.Index, 10)
	difficultyBytes := strconv.FormatInt(int64(b.Difficulty), 10)
	nonceBytes := strconv.FormatInt(int64(b.Nonce), 10)

	info := [][]byte{[]byte(timestampBytes), []byte(indexBytes), []byte(difficultyBytes), []byte(nonceBytes), b.Data, b.PrevBlockHash}
	joinedInfo := bytes.Join(info, []byte{})

	hash := sha256.Sum256(joinedInfo)
	return hash[:]
}

func (b *Block) Mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		hash := b.calculateHash()
		hexHash := hex.EncodeToString(hash)

		if strings.HasPrefix(hexHash, target) {
			b.Hash = hash
			break
		}
		b.Nonce++
	}
}

func NewBlock(data []byte, prevBlockHash []byte, index uint64) *Block {

	block := &Block{
		Index:         index,
		Data:          data,
		PrevBlockHash: prevBlockHash,
		Timestamp:     time.Now().Unix(),
		Difficulty:    6,
	}

	block.Mine()

	return block

}
