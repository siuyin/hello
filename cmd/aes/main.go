package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/siuyin/dflt"
)

func main() {
	const aes128KeyLen = 16
	const aes256KeyLen = 32
	if gk := dflt.EnvString("GENKEY", ""); gk != "" {
		genKey(aes256KeyLen)
		genNonce()
	}
	key:=readBin("key")
	nonce:=readBin("nonce")
	aescgm:=cgm(key,nonce)

	plaintext:=[]byte("The quick brown fox jumps over the lazy dog.")
	ciphertext:=aescgm.Seal(nil,nonce,plaintext,nil)
	fmt.Printf("ciphertext: %x\n",ciphertext)

	decrypted,err:=aescgm.Open(nil,nonce,ciphertext,nil)
	if err!=nil {log.Fatal(err)}
	fmt.Printf("decrypted: %s\n",decrypted)
}

func genKey(n int) {
	b := make([]byte, n)
	i, err := rand.Read(b)
	if err != nil || i != n {
		log.Fatal(err)
	}
	bs := hex.EncodeToString(b)
	fmt.Println(bs)
	f, err := os.Create("key")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(b)
	fmt.Println("key written")
}

func genNonce() {
	const nonceSize = 12
	b:=make([]byte,nonceSize)
	i, err := rand.Read(b)
	if err != nil || i != nonceSize {
		log.Fatal(err)
	}

	bs := hex.EncodeToString(b)
	fmt.Println(bs)

	f, err := os.Create("nonce")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.Write(b)
	fmt.Println("nonce written")
}

func readBin(fn string) []byte {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	bs := hex.EncodeToString(b)
	fmt.Println(fn,":", bs)

	return b
}

func cgm(key,nonce []byte) cipher.AEAD {
	block,err:=aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	aescgm,err:=cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	return aescgm

}