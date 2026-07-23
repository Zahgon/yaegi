package main

type Foo struct {
}

func (Foo) Call() { _ = "STUB: not implemented"; return }

type Bar struct {
	Foo
}

type Baz struct {
	Foo
}

func (Baz) Call() { _ = "STUB: not implemented"; return }

func main() {
	Foo{}.Call()
	Bar{}.Call()
	Baz{}.Call()
}
