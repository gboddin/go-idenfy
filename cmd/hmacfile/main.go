package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
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
	key, err := base64.StdEncoding.DecodeString(os.Args[2])
	if err != nil {
		panic(err)
	}
	mac := hmac.New(sha256.New, key)
	_, err = io.Copy(mac, file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Signature: %s", base64.StdEncoding.EncodeToString(mac.Sum(nil)))
}
