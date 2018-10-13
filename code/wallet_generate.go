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
	//chain := &chaincfg.MainNetParams // mainnet
	chain := &chaincfg.TestNet3Params // testnet

	priv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Fatal(err)
	}

	privBytes := priv.Serialize()
	fmt.Printf("private key [bytes]:\n%v\n\n", privBytes)                   // [18 214 ... 64 56]
	fmt.Printf("private key [hex]:\n%s\n\n", hex.EncodeToString(privBytes)) // 12d6913912cedcd1859778902bde0f737740ffb532cd1335b08aff159c474038
	fmt.Printf("private key [base58]:\n%s\n\n", base58.Encode(privBytes))   // 2GY6yKFr8FRX25zPtrAzLRko1Uryz7QWPy94Hw7i6Vaw

	uncWif, err := btcutil.NewWIF(priv, chain, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("private key [wif] (uncompressed):\n%s\n\n", uncWif.String()) // 5Jim1MwMAu5WY8puAKL4gLE7tTKijSqoa9rqXhPWeT38Jd1AfsD

	cmpWif, err := btcutil.NewWIF(priv, chain, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("private key [wif] (compressed):\n%s\n\n", cmpWif.String()) // 2GY6yKFr8FRX25zPtrAzLRko1Uryz7QWPy94Hw7i6Vaw

	pub := priv.PubKey()
	uncPubBytes := pub.SerializeUncompressed()
	cmpPubBytes := pub.SerializeCompressed()
	fmt.Printf("public key [bytes] (uncompressed):\n%v\n\n", uncPubBytes)                   // [4 210 ... 154 207]
	fmt.Printf("public key [hex] (uncompressed):\n%s\n\n", hex.EncodeToString(uncPubBytes)) // 04d28f502980c5e874c3dd2e4aff019b18e3bef83b5828cf974ffc87c8b0f94576611afbf8780fbff9e6a31c7e3b5385b3d24a0777a8b8f37cd6355ed43d219acf
	fmt.Printf("public key [base58] (uncompressed):\n%s\n\n", base58.Encode(uncPubBytes))   // RgbxSrecyPCc3jsEcDmLh5ERueFyrz7m1QEg3U4SUQAZhoPABbik2GvS9adSRHHTV3f2ourctb4qPjuYiyiLdH3k
	fmt.Printf("public key bytes (compressed):\n%v\n\n", cmpPubBytes)                       // [3 210 ... 69 118]
	fmt.Printf("public key [hex] (compressed):\n%s\n\n", hex.EncodeToString(cmpPubBytes))   // 03d28f502980c5e874c3dd2e4aff019b18e3bef83b5828cf974ffc87c8b0f94576
	fmt.Printf("public key [base58] (compressed):\n%s\n\n", base58.Encode(cmpPubBytes))     // 28rtUZpHgFeEKjkBzTqRxGwohCF8KmSaMS9o38VGzoA3X

	uncAddr, err := btcutil.NewAddressPubKey(uncPubBytes, chain)
	if err != nil {
		log.Fatal(err)
	}

	cmpAddr, err := btcutil.NewAddressPubKey(cmpPubBytes, chain)
	if err != nil {
		log.Fatal(err)
	}

	encUncAddr := uncAddr.EncodeAddress()
	encCmpAddr := cmpAddr.EncodeAddress()
	fmt.Printf("address [base58] (uncompressed):\n%s\n\n", encUncAddr) // 16385kYLPqkczsyhJirzjunz27bTpqJrNm
	fmt.Printf("address [base58] (compressed):\n%s\n\n", encCmpAddr)   // 15xQjUYRuk59ijmbCkSFTiP7zYWD4NVN1G
}
