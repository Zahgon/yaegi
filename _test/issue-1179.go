package main

type I interface {
	F()
}

type T struct {
	Name string
}

func (t *T) F() { _ = "STUB: not implemented"; return }

func NewI(s string) I { _ = "STUB: not implemented"; return *new(I) }

func newT(s string) *T { _ = "STUB: not implemented"; return nil }

func main() {
	i := NewI("test")
	i.F()
}
