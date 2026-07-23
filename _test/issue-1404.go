package main

type I interface {
	inI()
}

type T struct {
	name string
}

func (t *T) inI() { _ = "STUB: not implemented"; return }

func main() {
	var i I = &T{name: "foo"}

	if i, ok := i.(*T); ok {
		println(i.name)
	}
}
