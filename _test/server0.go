package main

import (
	"net/http"
)

var v string = "v1.0"

func myHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
