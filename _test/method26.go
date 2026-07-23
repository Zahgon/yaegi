package main

func NewT(name string) *T { _ = "STUB: not implemented"; return nil }

var C = NewT("test")

func (t *T) f() { _ = "STUB: not implemented"; return }

type T struct {
	Name string
}

func main() {
	C.f()
}
