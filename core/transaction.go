package core

import "github.com/firman-alam/go-blockchain/crypto"

type Transaction struct {
	Data []byte

	PublicKey	crypto.PublicKey
	Signature 	*crypto.Signature
}
