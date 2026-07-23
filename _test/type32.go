package main

import "fmt"

func (f *Foo) Bar() int { _ = "STUB: not implemented"; return 0 }

type Foo = int

func main() {
	x := Foo(1)
	fmt.Println(x.Bar())
}
