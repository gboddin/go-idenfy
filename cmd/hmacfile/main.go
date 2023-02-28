package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Need filename and key as arguments")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	key := []byte(os.Args[2])
	if err != nil {
		panic(err)
	}
	mac := hmac.New(sha256.New, key)
	_, err = io.Copy(mac, file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Signature: %s", hex.EncodeToString(mac.Sum(nil)))
}
