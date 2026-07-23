package main

import (
	"io"
	"os"
)

type sink interface {
	io.Writer
	io.Closer
}

func newSink() sink { _ = "STUB: not implemented"; return *new(sink) }

func main() {
	s := newSink()
	n, err := s.Write([]byte("Hello\n"))
	if err != nil {
		panic(err)
	}
	var writer io.Writer = s
	m, err := writer.Write([]byte("Hello\n"))
	if err != nil {
		panic(err)
	}
	var closer io.Closer = s
	err = closer.Close()
	if err != nil {
		panic(err)
	}
	err = os.Remove(s.(*os.File).Name())
	if err != nil {
		panic(err)
	}
	println(m, n)
}
