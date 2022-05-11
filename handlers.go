package main

import (
	crypto "crypto/ed25519"
	"fmt"
	"io"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {

	// check crypto/ed25519 documentation here https://pkg.go.dev/crypto/ed25519
	signature := crypto.Sign([]byte("edsk4CLwrvzBrd2njFMr6EgGcrYrpptZu8h8hbYEyFVd2kQxHVC6t3hhhhhhhhhh"), []byte("message"))

	fmt.Println("signature: ", signature)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("{\"status\":\"SUCCESS\"}"))
}
