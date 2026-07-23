package main

import "fmt"

type Foo struct{}

func foo() *Foo { _ = "STUB: not implemented"; return nil }

func main() {
	f := foo()
	fmt.Println(f)
}
