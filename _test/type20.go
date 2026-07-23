package main

import (
	"io"
	"strings"
)

func isCloser(r io.Reader) bool { _ = "STUB: not implemented"; return false }

func main() {
	println(isCloser(strings.NewReader("test")))
}
