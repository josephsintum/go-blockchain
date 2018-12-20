package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block is the struct of the block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock takes string of data and prevHash and outputs pointer to a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis takes Block and returns a create block call
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Serialize takes block and returns bytes
// doing this cause badgerdb takes only bytes and arrays
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

// Deserialize takes in data []byte and returns pointer to block
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

// Handle err
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
