package main

type Foo struct {
}

func (Foo) Show() { _ = "STUB: not implemented"; return }

func (f Foo) Call() { _ = "STUB: not implemented"; return }

type Bar struct {
	Foo
}

type Baz struct {
	Foo
}

func (Baz) Call() { _ = "STUB: not implemented"; return }

func (Baz) Show() { _ = "STUB: not implemented"; return }

func main() {
	Foo{}.Call()
	Bar{}.Call()
	Baz{}.Call()
}
