package core

import (
	"io"

	"github.com/firman-alam/go-blockchain/crypto"
	"github.com/firman-alam/go-blockchain/types"
)

type Header struct {
	Version   		uint32
	DataHash		types.Hash
	PrevBlockHash 	types.Hash
	Timestamp 		int64
	Height 			uint32
}

type Block struct {
	*Header
	Transcations 	[]Transaction
	Validator		crypto.PublicKey
	Signature		*crypto.Signature

	// cache version of the header hash
	hash 			types.Hash
}

func NewBlock(h *Header, tx []Transaction) *Block {
	return &Block{
		Header: 		h,
		Transcations:	tx,
	}
}

// func (b *Block) Decode

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, dec Encoder[*Block]) error {
	return dec.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}