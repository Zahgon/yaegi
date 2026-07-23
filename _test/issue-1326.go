package main

type Option interface {
	apply(*T)
}

type T struct {
	s string
}

type opt struct {
	name string
}

func (o *opt) apply(t *T) { _ = "STUB: not implemented"; return }

func BuildOptions() []Option { _ = "STUB: not implemented"; return nil }

func NewT(name string, options ...Option) *T { _ = "STUB: not implemented"; return nil }

func main() {
	t := NewT("hello", BuildOptions()...)
	println(t.s)
}
