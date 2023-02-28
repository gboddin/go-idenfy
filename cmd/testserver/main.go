package main

import (
	"encoding/json"
	"github.com/gboddin/go-idenfy"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/v2/token", genToken)
	log.Println("Starting server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func genToken(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		log.Println("error: expected POST")
		w.WriteHeader(400)
		return
	}
	tee := io.TeeReader(req.Body, os.Stdout)
	decoder := json.NewDecoder(tee)
	var response idenfy.TokenResponse
	err := decoder.Decode(&response)
	os.Stdout.WriteString("\n")
	if err != nil {
		log.Println("error: decoding JSON failed", err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	response.AuthToken = "TestingToken72365da647f5aea5d97f8d0bfa72be07c832ed5d"
	encoder := json.NewEncoder(w)
	w.WriteHeader(200)
	if err := encoder.Encode(&response); err != nil {
		log.Println(err)
	}
}
