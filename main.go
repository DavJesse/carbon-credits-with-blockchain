package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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
	chain []*Block
}

func newBlock(data, previousHash string, chain []*Block) *Block {
	block := &Block{
		Index:        len(chain) + 1,
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: previousHash,
	}
	block.Hash = calculateHash(block)
	return block
}

func calculateHash(block *Block) string {
	record := strconv.Itoa(block.Index) + strconv.Itoa(int(block.Timestamp)) + block.Data + block.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func AddBlock(blockchain *Blockchain, data string, previousHash string) {
	block := newBlock(data, previousHash, blockchain.chain) // Assuming NewBlock takes chain as an argument
	blockchain.chain = append(blockchain.chain, block)
}

func main() {
	blockchain := &Blockchain{}
	AddBlock(blockchain, "Genesis Block", "0")
	AddBlock(blockchain, "Block 1", blockchain.chain[0].Hash)

	for _, block := range blockchain.chain {
		fmt.Printf("Block: %+v\n", block)
	}
}
