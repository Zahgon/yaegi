package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func client(uri string) { _ = "STUB: not implemented"; return }

func server(ln net.Listener, ready chan bool) { _ = "STUB: not implemented"; return }

func main() {
	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	ready := make(chan bool)
	go server(ln, ready)
	<-ready

	client(fmt.Sprintf("http://%s/hello", ln.Addr().String()))
	http.DefaultServeMux = &http.ServeMux{}
}
