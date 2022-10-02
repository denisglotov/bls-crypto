package bls

import (
	"math/rand"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

// GenerateRandomKeyForTest creates a random private and its corresponding
// public keys, this is determined version of GenerateRandomKey, that uses
// pseudorandom numbers. For testing purposes only!
func GenerateRandomKeyForTest() (PrivateKey, PublicKey) {
	priv, pub, _ := bn256.RandomG2(mathRandReader{})
	return PrivateKey{p: priv}, PublicKey{p: pub}
}

type mathRandReader struct{}

func (mathRandReader) Read(b []byte) (n int, err error) {
	return rand.Read(b)
}
