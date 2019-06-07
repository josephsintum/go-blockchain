package wallet

import (
	"log"

	"github.com/mr-tron/base58"
)

func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}

// base58 func does have 0 O 1 I + /
// this reduces the chance of errors while copying str return like in a public key address