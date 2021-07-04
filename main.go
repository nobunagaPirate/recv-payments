package main

import (
	"fmt"
	"github.com/btcsuite/btcd/rpcclient"
	"log"
)

const walletName = "gowallet"

func main() {
	config := &rpcclient.ConnConfig{
		Host:         "localhost:18332",
		User:         "test",
		Pass:         "test",
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(config, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Shutdown()

	res, err := client.CreateWallet(walletName)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Wallet created: ", res.Name)
}
