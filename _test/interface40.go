package main

import "fmt"

type foo struct {
	bar string
}

func (f foo) String() string { _ = "STUB: not implemented"; return "" }

func Foo(s string) fmt.Stringer { _ = "STUB: not implemented"; return *new(fmt.Stringer) }

func main() {
	f := Foo("bar")
	fmt.Println(f)
}
