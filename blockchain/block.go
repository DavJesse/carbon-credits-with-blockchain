package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	Timestamp    int64
	Data         string
	PreviousHash string
	Hash         string
}

type Blockchain struct {
	Chain []*Block
}

func newBlock(data, previousHash string) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: previousHash,
	}
	block.calculateHash()
	return block
}

func (b *Block) calculateHash() {
	record := strconv.Itoa(b.Index) + strconv.Itoa(int(b.Timestamp)) + b.Data + b.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	b.Hash = hex.EncodeToString(hashed)
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Chain[len(chain.Chain)-1]
	block := newBlock(data, prevBlock.Hash) // Assuming NewBlock takes chain as an argument
	chain.Chain = append(chain.Chain, block)
}

func Genesis() *Block {
	return newBlock("Genesis Block", "0")
}

func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
