package main

import (
	"fmt"
	"bytes"
	"crypto/sha256"
)

// BlockChain contains array of blocks
type BlockChain struct {
	blocks []*Block
}

// Block is the struct of the block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash produces Hash from Data and PrevHash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock takes string of data and prevHash and outputs pointer to a block
func CreateBlock(data string, PrevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), PrevHash}
	block.DeriveHash()
	return block
}

// AddBlock takes BlockChain and data string 
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Genesis takes Block and returns a create block call
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain takes BlockChain and returns an array of blocks with a call to Genesis()
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}

}
