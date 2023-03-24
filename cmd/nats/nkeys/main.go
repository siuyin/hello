package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nkeys"
	"github.com/nats-io/nuid"
)

func main() {
	fmt.Println("NATS keys")
	user, err := nkeys.CreateUser()
	//user, err := nkeys.CreateAccount()
	if err != nil {
		log.Fatal("1 ", err)
	}
	pk, _ := user.PublicKey()
	fmt.Printf("public key: %s\n", pk)

	challenge := nuid.Next()
	fmt.Println(challenge)

	signature, _ := user.Sign([]byte(challenge))

	seed, err := user.Seed()
	if err != nil {
		log.Fatal("2 ", err)
	}
	fmt.Printf("seed: %s\n", seed)

	usr2, _ := nkeys.FromSeed(seed)
	seed2, _ := usr2.Seed()
	if string(seed2) != string(seed) {
		log.Fatalf("seed compare failed")
	}

	pubKey, _ := user.PublicKey()
	usr3, _ := nkeys.FromPublicKey(pubKey)
	if err := usr3.Verify([]byte(challenge), signature); err != nil {
		log.Printf("challenge failed: %v")
	}

}
