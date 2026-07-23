package main

type I interface {
	Foo() string
}

type Printer struct {
	i I
}

func New(i I) *Printer { _ = "STUB: not implemented"; return nil }

func (p *Printer) Print() { _ = "STUB: not implemented"; return }

type T struct{}

func (t *T) Foo() string { _ = "STUB: not implemented"; return "" }

func main() {
	g := New(&T{})
	g.Print()
}
