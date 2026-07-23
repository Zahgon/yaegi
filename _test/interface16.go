package main

import "fmt"

type Barer interface {
	fmt.Stringer
	Bar()
}

type T struct{}

func (*T) String() string { _ = "STUB: not implemented"; return "" }
func (*T) Bar()           { _ = "STUB: not implemented"; return }

var t = &T{}

func main() {
	var f Barer
	if f != t {
		println("ok")
	}
}
