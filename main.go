package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jasonsalas/bc/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()
	chain.AddBlock("Siamese Dream")
	chain.AddBlock("Too Fast For Love")
	chain.AddBlock("The Donnas")
	chain.AddBlock("British Steel")
	chain.AddBlock("Operation: Mindcrime")
	chain.AddBlock("Randy Rhoads: Tribute")
	chain.AddBlock("Pride")

	for _, block := range chain.Blocks {
		time.Sleep(1 * time.Second)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("previous hash: %x\n", block.PrevHash)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("----")
	}
}
