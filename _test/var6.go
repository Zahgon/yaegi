package main

import (
	"fmt"
)

type Foo struct {
	A string
}

var f = Foo{"world"}

func Hello() { _ = "STUB: not implemented"; return }

var name = "v1"

func main() {
	Hello()
	fmt.Println("Hello", f.A, name)
}
