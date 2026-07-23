package main

import "fmt"

type Foo = int

func (f Foo) Bar() int { _ = "STUB: not implemented"; return 0 }

func main() {
	x := Foo(1)
	fmt.Println(x.Bar())
}
