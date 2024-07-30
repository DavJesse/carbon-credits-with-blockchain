package main

import (
	"carbomoney/blockchain"
	"fmt"
)

func main() {
	blockchain := blockchain.InitBlockChain()
	blockchain.AddBlock("Genesis Block")
	blockchain.AddBlock("Block 1")

	for _, block := range blockchain.Chain {
		fmt.Printf("Block: %+v\n", block)
	}
}
