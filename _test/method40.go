package main

import (
	"bytes"
	"io"
)

type TMemoryBuffer struct {
	*bytes.Buffer
	size int
}

func newTMemoryBuffer() *TMemoryBuffer { _ = "STUB: not implemented"; return nil }

var globalMemoryBuffer = newTMemoryBuffer()

type TTransport interface {
	io.ReadWriter
}

func check(t TTransport) { _ = "STUB: not implemented"; return }

func main() {
	check(globalMemoryBuffer)
}
