package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type readAutoCloser struct {
	r io.ReadCloser
}

func (a readAutoCloser) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (a readAutoCloser) Close() error { _ = "STUB: not implemented"; return nil }

type pipe struct {
	Reader readAutoCloser
}

func newReadAutoCloser(r io.Reader) readAutoCloser {
	_ = "STUB: not implemented"
	return *new(readAutoCloser)
}

func main() {
	p := &pipe{}
	p.Reader = newReadAutoCloser(strings.NewReader("test"))
	b, err := io.ReadAll(p.Reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
