package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

func main() {
	privWif := "cS5LWK2aUKgP9LmvViG3m9HkfwjaEJpGVbrFHuGZKvW2ae3W9aUe"
	txHash := "12e0d25258ec29fadf75a3f569fccaeeb8ca4af5d2d34e9a48ab5a6fdc0efc1e"
	destination := "mrdKfqWEkwferzEQus5NpgK2Dtpq7Qcgif"
	amount := int64(11650795)
	txFee := int64(500000)
	sourceUTXOIndex := uint32(1)
	chainParams := &chaincfg.TestNet3Params // testnet
	// chainParams := &chaincfg.MainNetParams // mainnet

	decodedWif, err := btcutil.DecodeWIF(privWif)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decoded WIF: %v\n", decodedWif) // Decoded WIF: cS5LWK2aUKgP9LmvViG3m9HkfwjaEJpGVbrFHuGZKvW2ae3W9aUe

	addressPubKey, err := btcutil.NewAddressPubKey(decodedWif.PrivKey.PubKey().SerializeUncompressed(), chainParams)
	if err != nil {
		log.Fatal(err)
	}

	sourceUTXOHash, err := chainhash.NewHashFromStr(txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("UTXO hash: %s\n", sourceUTXOHash) // utxo hash: 12e0d25258ec29fadf75a3f569fccaeeb8ca4af5d2d34e9a48ab5a6fdc0efc1e

	sourceUTXO := wire.NewOutPoint(sourceUTXOHash, sourceUTXOIndex)
	sourceTxIn := wire.NewTxIn(sourceUTXO, nil, nil)
	destinationAddress, err := btcutil.DecodeAddress(destination, chainParams)
	if err != nil {
		log.Fatal(err)
	}

	sourceAddress, err := btcutil.DecodeAddress(addressPubKey.EncodeAddress(), chainParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Source Address: %s\n", sourceAddress) // Source Address: mgjHgKi1g6qLFBM1gQwuMjjVBGMJdrs9pP

	destinationPkScript, err := txscript.PayToAddrScript(destinationAddress)
	if err != nil {
		log.Fatal(err)
	}

	sourcePkScript, err := txscript.PayToAddrScript(sourceAddress)
	if err != nil {
		log.Fatal(err)
	}

	sourceTxOut := wire.NewTxOut(amount, sourcePkScript)

	redeemTx := wire.NewMsgTx(wire.TxVersion)
	redeemTx.AddTxIn(sourceTxIn)
	redeemTxOut := wire.NewTxOut((amount - txFee), destinationPkScript)
	redeemTx.AddTxOut(redeemTxOut)

	sigScript, err := txscript.SignatureScript(redeemTx, 0, sourceTxOut.PkScript, txscript.SigHashAll, decodedWif.PrivKey, false)
	if err != nil {
		log.Fatal(err)
	}

	redeemTx.TxIn[0].SignatureScript = sigScript
	fmt.Printf("Signature Script: %v\n", hex.EncodeToString(sigScript)) // Signature Script: 473...b67

	// validate signature
	flags := txscript.StandardVerifyFlags
	vm, err := txscript.NewEngine(sourceTxOut.PkScript, redeemTx, 0, flags, nil, nil, amount)
	if err != nil {
		log.Fatal(err)
	}

	if err := vm.Execute(); err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(make([]byte, 0, redeemTx.SerializeSize()))
	redeemTx.Serialize(buf)

	fmt.Printf("Redeem Tx: %v\n", hex.EncodeToString(buf.Bytes())) // redeem Tx: 01000000011efc...5bb88ac00000000
}
