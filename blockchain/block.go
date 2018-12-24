package blockchain

import (
	"crypto/sha256"
	"bytes"
	"encoding/gob"
	"log"
)

// Block is the struct of the block
type Block struct {
	Hash     	[]byte
	Transactions []*Transaction
	PrevHash 	[]byte
	Nonce    	int
}

// HashTransactions creates hash of transaction
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

// CreateBlock takes string of data and prevHash and outputs pointer to a block
func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis takes Block and returns a create block call
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
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
