package main

import (
	"io"
	"log"
	"strings"
)

type WrappedReader struct {
	reader io.Reader
}

func (wr WrappedReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (wr WrappedReader) WriteTo(w io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func main() {
	f := strings.NewReader("hello world")
	wr := WrappedReader{reader: f}

	if _, err := io.Copy(io.Discard, wr); err != nil {
		log.Fatal(err)
	}
}
