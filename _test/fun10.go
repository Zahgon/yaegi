package main

import "fmt"

func f() func() { _ = "STUB: not implemented"; return nil }

func main() {
	g := f()
	fmt.Printf("%T %v\n", g, g)
	if g == nil {
		fmt.Println("nil func")
	}
}
