---
description: Tutorial on how to read account balances from the bitcoin blockchain with Go.
---

# Account Balances

---

### Full code

[account_balance.go](https://github.com/miguelmota/bitcoin-development-with-go-book/blob/master/code/account_balance.go)

```go
package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
)

func main() {
	address := "3BMEXpRXTdAHagbFtjtuSS3S8ZXfQQqiTw"

	req := struct {
		ID     int      `json:"id"`
		Method string   `json:"method"`
		Params []string `json:"params"`
	}{
		ID:     1,
		Method: "blockchain.address.get_balance",
		Params: []string{address},
	}

	res := struct {
		JSONRPC string `json:"jsonrpc,omitempty"`
		ID      int    `json:"id"`
		Result  struct {
			Confirmed   *big.Int `json:"confirmed"`
			Uncomfirmed *big.Int `json:"unconfirmed"`
		} `json:"result"`
	}{}

	serverAddr := "electrum.qtornado.com:50002" // mainnet

	certBytes, err := ioutil.ReadFile("certs/example.com.cert")
	if err != nil {
		log.Fatal(err)
	}
	certKeyBytes, err := ioutil.ReadFile("certs/example.com.key")
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.X509KeyPair(certBytes, certKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("dialing to server: %s\n", serverAddr)
	conn, err := tls.Dial("tcp", serverAddr, &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	fmt.Printf("client connected to: %s\n", conn.RemoteAddr())

	reqMsgBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	reqMsg := fmt.Sprintf("%s\n", string(reqMsgBytes))
	fmt.Printf("writing message: %s", reqMsg)
	_, err = io.WriteString(conn, reqMsg)
	if err != nil {
		log.Fatal(err)
	}

	var (
		i        int
		readSize int = 1024
		respData []byte
	)

	for {
		fmt.Println("reading response...")
		respBytes := make([]byte, readSize)
		n, err := conn.Read(respBytes)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
		}

		fmt.Printf("reading: %q (%d bytes)\n", string(respBytes[:n]), n)

		respData = append(respData, respBytes[:n]...)
		i += n

		if n < readSize {
			break
		}
	}

	json.Unmarshal(respData[:i], &res)

	fmt.Printf("unconfirmed: %s\n", res.Result.Uncomfirmed.String()) // unconfirmed: 0
	fmt.Printf("confirmed: %s\n", res.Result.Confirmed.String())     // confirmed: 500000000
}
```
