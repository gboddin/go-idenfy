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
