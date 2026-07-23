package main

type I interface {
	A() string
	B() string
}

type s struct{}

func NewS() (I, error) { _ = "STUB: not implemented"; return *new(I), nil }

func (c *s) A() string { _ = "STUB: not implemented"; return "" }
func (c *s) B() string { _ = "STUB: not implemented"; return "" }

func main() {
	s, _ := NewS()
	println(s.A())
}
