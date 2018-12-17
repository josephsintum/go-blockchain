package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// BlockChain contains array of Blocks
type BlockChain struct {
	Blocks []*Block
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
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis takes Block and returns a create block call
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain takes BlockChain and returns an array of Blocks with a call to Genesis()
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
