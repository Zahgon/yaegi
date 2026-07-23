package main

import "net/http"

type T struct {
	header string
}

func (b *T) ServeHTTP(rw http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func main() {
	println("ok")
}
