// Derived from https://github.com/MarinX/btc-vanity
// Last edit on 4/19/2022 - Bhaven Patel

package main

import (
	"fmt"

	btcvanity "github.com/MarinX/btc-vanity"
)

func main() {

	// create configuration
	cfg := &btcvanity.Config{
		// buffered channel, more buffer, faster to find matching pattern
		Buffer: 5,
		// if you want to use testnet, set true
		TestNet: true,
	}

	btc := btcvanity.New(cfg)

	// find a patters eg adddress which starts with "a"
	address, err := btc.Find("bhav")
	if err != nil {
		panic(err)
	}

	// print our custom public key
	fmt.Printf("PUBLIC KEY\n%s\n", address.PublicKey())

	// print our private key so it can be imported in most btc wallets
	fmt.Printf("PRIVATE KEY\n%s\n", address.PrivateKey())
}
