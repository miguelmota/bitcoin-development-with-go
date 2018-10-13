package util

import (
	"crypto/sha256"

	"github.com/btcsuite/btcutil/base58"
)

func P2PKHToAddress(pkscript []byte, isTestnet bool) (string, error) {
	p := make([]byte, 1)
	p[0] = 0x00 // prefix with 00 if it's mainnet
	if isTestnet {
		p[0] = 0x6F // prefix with 0F if it's testnet
	}
	pub := pkscript[3 : len(pkscript)-2] // get pkhash
	pf := append(p[:], pub[:]...)        // add prefix
	h1 := sha256.Sum256(pf)              // hash it
	h2 := sha256.Sum256(h1[:])           // hash it again
	b := append(pf[:], h2[0:4]...)       // prepend the prefix to the first 5 bytes
	address := base58.Encode(b)          // encode to base58
	if !isTestnet {
		address = "1" + address // prefix with 1 if it's mainnet
	}

	return address, nil
}
