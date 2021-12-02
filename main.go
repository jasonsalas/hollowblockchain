package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
	}
	block.DeriveHash()
	return block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{
		blocks: []*Block{
			Genesis(),
		},
	}
}

func main() {
	chain := InitBlockchain()
	chain.AddBlock("Block 001")
	chain.AddBlock("Block 002")
	chain.AddBlock("Block 003")
	chain.AddBlock("Block 004")

	for _, block := range chain.blocks {
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("previous hash: %x\n", block.PrevHash)
		fmt.Printf("hash: %x\n\n", block.Hash)
	}
}
