package main

func (f *Foo) Boo() { _ = "STUB: not implemented"; return }

type Foo struct {
	name string
	fun  func(f *Foo)
}

func main() {
	t := &Foo{name: "foo"}
	t.Boo()
}
