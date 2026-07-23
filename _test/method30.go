package main

type T struct {
	Name string
}

func (t *T) foo(a string) string { _ = "STUB: not implemented"; return "" }

var g = &T{"global"}

var f = g.foo

func main() {
	println(f("-x"))
}
