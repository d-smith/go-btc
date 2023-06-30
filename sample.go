package main

import (
	"fmt"
	"log"

	"github.com/wemeetagain/go-hdwallet"
)

func main() {
	// Generate a random 256 bit seed
	seed, err := hdwallet.GenSeed(256)
	if err != nil {
		log.Fatal(err)
	}

	// Create a master private key
	masterprv := hdwallet.MasterKey(seed)
	fmt.Println("private key", masterprv)

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	fmt.Println(masterpub)
	fmt.Println(masterpub.Address())
	/*
	   // Generate new child key based on private or public key
	   childprv, err := masterprv.Child(0)
	   childpub, err := masterpub.Child(0)

	   // Create bitcoin address from public key
	   address := childpub.Address()

	   // Convenience string -> string Child and ToAddress functions
	   walletstring := childpub.String()
	   childstring, err := hdwallet.StringChild(walletstring, 0)
	   childaddress, err := hdwallet.StringAddress(childstring)
	*/
}
