package main

type Fooer interface {
	Foo() string
}

type Barer interface {
	Fooer
	Bar()
}

type T struct{}

func (t *T) Foo() string { _ = "STUB: not implemented"; return "" }
func (*T) Bar()          { _ = "STUB: not implemented"; return }

var t = &T{}

func main() {
	var f Barer
	if f != t {
		println("ok")
	}
}
