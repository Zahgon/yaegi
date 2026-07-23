package main

func Bar() { _ = "STUB: not implemented"; return }

var Obj = NewT()

func NewT() *T { _ = "STUB: not implemented"; return nil }

type T struct{}

func (t *T) Foo() bool { _ = "STUB: not implemented"; return false }

func main() {
	Bar()
}
