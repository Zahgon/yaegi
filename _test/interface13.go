package main

import (
	"fmt"
)

type X struct{}

func (X) Foo() int { _ = "STUB: not implemented"; return 0 }

func (X) Bar() int { _ = "STUB: not implemented"; return 0 }

type Foo interface {
	Foo() int
}
type Bar interface {
	Bar() int
}

func main() {
	var x X
	var i Foo = x
	j := i.(Bar)

	fmt.Println(j.Bar())
}
