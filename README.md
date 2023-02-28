# INCOMPLETE, WORK IN PROGRESS
# Idenfy Golang Library

[![GitHub Release](https://img.shields.io/github/v/release/gboddin/go-idenfy)](https://github.com/gboddin/go-idenfy/releases)
[![Follow on Twitter](https://img.shields.io/twitter/follow/gboddin.svg?logo=twitter)](https://twitter.com/gboddin)

Idenfy Golang is a Golang library for requesting and receiving verifications from the Idenfy API.

## Features

- Client side request validation
- Token request
- Webhook parsing
- Webhook signature verification (mandatory)

## Usage

See the [Go documentation](https://pkg.go.dev/github.com/gboddin/go-idenfy) for more details.

### Request verification token

```golang
package main

import (
	"context"
	"github.com/gboddin/go-idenfy"
	"log"
)

func main() {
	// Create an idenfy client
	client, err := idenfy.NewClient(
		idenfy.WithApiCredentials("access-key", "secret-key"),
		idenfy.WithCustomEndpoint("http://localhost:8080"),
	)
	if err != nil {
		panic(err)
	}
	// Create a token request
	tokenReq := idenfy.TokenRequest{
		ClientId:            "your-local-id",
		DateOfBirth:         &idenfy.Date{Year: 1981, Month: 3, Day: 4},
		GenerateDigitString: idenfy.Bool(false),
	}
	// create verification session
	tokenResp, err := client.CreateIdentityVerificationSession(context.Background(), tokenReq)
	if err != nil {
		panic(err)
	}
	// Get authToken and/or redirect URL
	log.Println(tokenResp.AuthToken, tokenResp.DateOfBirth, tokenResp.GetRedirectUrl())
}

```

### Receive webhook

```golang
package main

import (
	"github.com/gboddin/go-idenfy"
	"log"
	"net/http"
)

var idenfyClient *idenfy.Client

func main() {
	var err error
	idenfyClient, err = idenfy.NewClient(
		idenfy.WithCallbackSignatureKey("signature-key"),
		idenfy.WithCustomEndpoint("http://localhost:8080"),
	)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/webhook/idenfy", receiveHook)
	log.Println("Starting server at http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}

func receiveHook(w http.ResponseWriter, req *http.Request) {
	// Decode hook request and verify its signature
	reply, err := idenfyClient.DecodeHttpRequestIdentityCallback(req)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	// Do something with the reply
	log.Println(reply.ClientId)
}
```
