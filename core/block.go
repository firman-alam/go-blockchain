package core

import (
	"github.com/firman-alam/go-blockchain/crypto"
	"github.com/firman-alam/go-blockchain/types"
)

type Header struct {
	Version   	uint32
	DataHash	types.Hash
	PrevBlock 	types.Hash
	Timestamp 	int64
	Height 		uint32
}

type Block struct {
	Header
	Transcations 	[]Transaction
	Validator		crypto.PublicKey
	Signature		*crypto.Signature

	// cache version of the header hash
	hash 			types.Hash
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}