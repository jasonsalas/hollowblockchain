package main

import (
	"fmt"
	"strconv"

	"github.com/jasonsalas/bc/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()
	chain.AddBlock("Block 001")
	chain.AddBlock("Block 002")
	chain.AddBlock("Block 003")
	chain.AddBlock("Block 004")

	for _, block := range chain.Blocks {
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("previous hash: %x\n", block.PrevHash)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("----")
	}
}
