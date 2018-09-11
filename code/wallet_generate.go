package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
)

func main() {
	priv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Fatal(err)
	}

	privBytes := priv.Serialize()
	fmt.Printf("private key [bytes]:\n%v\n\n", privBytes)
	fmt.Printf("private key [hex]:\n%s\n\n", hex.EncodeToString(privBytes))
	fmt.Printf("private key [base58]:\n%s\n\n", base58.Encode(privBytes))

	pub := priv.PubKey()
	uncPubBytes := pub.SerializeCompressed()
	cmpPubBytes := pub.SerializeUncompressed()
	fmt.Printf("public key [bytes] (uncompressed):\n%v\n\n", uncPubBytes)
	fmt.Printf("public key [hex] (uncompressed):\n%s\n\n", hex.EncodeToString(uncPubBytes))
	fmt.Printf("public key [base58] (uncompressed):\n%s\n\n", base58.Encode(uncPubBytes))
	fmt.Printf("public key bytes (compressed):\n%v\n\n", cmpPubBytes)
	fmt.Printf("public key [hex] (compressed):\n%s\n\n", hex.EncodeToString(cmpPubBytes))
	fmt.Printf("public key [base58] (compressed):\n%s\n\n", base58.Encode(cmpPubBytes))

	uncAddr, err := btcutil.NewAddressPubKey(uncPubBytes, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}

	cmpAddr, err := btcutil.NewAddressPubKey(cmpPubBytes, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}

	encUncAddr := uncAddr.EncodeAddress()
	encCmpAddr := cmpAddr.EncodeAddress()
	fmt.Printf("address [base58] (uncompressed):\n%s\n\n", encUncAddr)
	fmt.Printf("address [base58] (compressed):\n%s\n\n", encCmpAddr)
}
