package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec"
)

func main() {
	privHex := "12d6913912cedcd1859778902bde0f737740ffb532cd1335b08aff159c474038"
	privBytes, err := hex.DecodeString(privHex)
	if err != nil {
		log.Fatal(err)
	}

	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), privBytes)
	fmt.Printf("private key [hex]:\n%x\n\n", priv.Serialize())

	/*
		privUncWifB58 := "2GY6yKFr8FRX25zPtrAzLRko1Uryz7QWPy94Hw7i6Vaw"
		privWif, err := btcutil.DecodeWIF(privUncWifB58)
		if err != nil {
			log.Fatal(err)
		}
		priv2 := privWif.PrivKey
		fmt.Printf("private key2 [hex]:\n%x\n\n", priv2.Serialize())
	*/

	uncPubHex := "048be27052fff64179cdb83d5e360606e6c696cf05445815cdb8ed2f47f8bb0a8e11af9ab997ef643262df572defb1af55ea876b48830ca99585e613cd7ac04ab0"
	uncPubBytes, err := hex.DecodeString(uncPubHex)
	if err != nil {
		log.Fatal(err)
	}

	pub, err := btcec.ParsePubKey(uncPubBytes, btcec.S256())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("public key [hex] (uncompressed):\n%x\n\n", pub.SerializeUncompressed())
}
