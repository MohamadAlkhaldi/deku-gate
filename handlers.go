package main

import (
	"net/http"
	"io"
	"fmt"
)

func Test(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("{\"status\":\"SUCCESS\"}"))
}