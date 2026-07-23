package main

import "fmt"

type foo struct {
	bar string
}

func (f *foo) String() string { _ = "STUB: not implemented"; return "" }

func main() {
	var f fmt.Stringer = &foo{bar: "bar"}
	fmt.Println(f)
}
