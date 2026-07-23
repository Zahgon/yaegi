package main

type I interface {
	Hello()
}

type T struct {
	h I
}

func (t *T) Hello() { _ = "STUB: not implemented"; return }

func main() {
	t := &T{}
	println(t.h != nil)
	println(t.h == nil)
	t.h = t
	println(t.h != nil)
	println(t.h == nil)
	t.h.Hello()
}
