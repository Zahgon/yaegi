package main

func Bar() { _ = "STUB: not implemented"; return }

var Obj = &T{}

type T struct{}

func (t *T) Foo() bool { _ = "STUB: not implemented"; return false }

func main() {
	Bar()
}
