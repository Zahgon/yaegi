package main

type Filter interface {
	Foo()
}

type GIFT struct {
	Filters []Filter
}

func New(filters ...Filter) *GIFT { _ = "STUB: not implemented"; return nil }

func (g *GIFT) List() { _ = "STUB: not implemented"; return }

type MyFilter struct{}

func (f *MyFilter) Foo() { _ = "STUB: not implemented"; return }

func main() {
	g := New(&MyFilter{})
	g.List()
}
