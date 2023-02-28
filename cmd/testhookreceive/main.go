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
	http.HandleFunc("/webhook/idenfy/identity", receiveHook)
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
