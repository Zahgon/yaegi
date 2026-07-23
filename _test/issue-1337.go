package main

import (
	"io"
	"os"
)

func f(i interface{}) { _ = "STUB: not implemented"; return }

func main() {
	var fd *os.File
	var r io.Reader = fd
	f(r)
}
